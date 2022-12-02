package main

import (
	"encoding/json"
	"fmt"
	"log"
	"manager/app"
)

func main() {
	config := app.GetConfig("config.yml")
	kafkaconfig := app.SettingKafka{
		KafkaServerAddr: config.KafkaInfo.Host,
		KafkaServerPort: config.KafkaInfo.Port,
	}

	tcpRequestChan := make(chan *app.DataCode)
	tcpRequestKeys := app.DataCode{}
	tcpResponseData := make(chan []byte)

	ch := app.ChannelStruct{}
	ch.ChannelInit()
	consumer := app.KafkaConsumerControllerInit(&kafkaconfig, "host")
	paritionList := app.KafkaConsumerControllerPartition(&kafkaconfig, consumer, "host")

	go app.TcpProcessing(tcpRequestChan, tcpResponseData, config.ApiServerInfo)

	go app.ConsumerHost(ch.ConsumerData.Host, kafkaconfig, paritionList)
	go app.ConsumerPerf(ch.ConsumerData.Realtimeperf, kafkaconfig, paritionList)
	go app.ConsumerPid(ch.ConsumerData.Realtimepid, kafkaconfig, paritionList)
	go app.ConsumerDisk(ch.ConsumerData.Realtimedisk, kafkaconfig, paritionList)
	go app.ConsumerNet(ch.ConsumerData.Realtimenet, kafkaconfig, paritionList)

	go app.GetDemoAgentinfo(ch.Agentinfo, config.Demo.HostCount)
	go app.GetDemoChangeStateAgent(ch.ChangeStateAgentStr, config.Demo)
	go app.GetDemoLastrealtimeperf(ch.Lastrealtimeperf, config.Demo)

	db_handler := make([]app.DBHandler, 0)
	agentinfo := <-ch.Agentinfo
	for _, dbinfo := range config.Database {
		db_handler = append(db_handler, *app.DBInit(dbinfo, config.Demo, agentinfo))
	}

	for {
		select {
		case state_agent_str := <-ch.ChangeStateAgentStr:
			if tcpRequestKeys.Code == app.HOST_CODE {
				agent_json, err := json.Marshal("A")
				if err != nil {
					log.Println("JSON Data Conversion error")
				}
				tcpResponseData <- agent_json
			}
			for _, d := range db_handler {
				d.DemoHostStateChange(state_agent_str)
			}
		case lrtp := <-ch.Lastrealtimeperf:
			if tcpRequestKeys.Code == app.LASTPERF_CODE {
				bpt_json, err := json.Marshal("B")
				if err != nil {
					log.Println("JSON Data Conversion error")
				}
				tcpResponseData <- bpt_json
			}
			for _, d := range db_handler {
				d.DemoBptUpdate(lrtp)
			}
		case cshost := <-ch.ConsumerData.Host:
			for _, d := range db_handler {
				d.SetHost(cshost)
			}
		case csperf := <-ch.ConsumerData.Realtimeperf:
			for _, d := range db_handler {
				d.SetPerf(csperf)
			}
		case cspid := <-ch.ConsumerData.Realtimepid:
			for _, d := range db_handler {
				d.SetPid(cspid)
			}
		case csdisk := <-ch.ConsumerData.Realtimedisk:
			for _, d := range db_handler {
				d.SetDisk(csdisk)
			}
		case csnet := <-ch.ConsumerData.Realtimenet:
			for _, d := range db_handler {
				d.SetNet(csnet)
			}
		case req_keys := <-tcpRequestChan:
			fmt.Printf("main %v\n", req_keys)
			tcpRequestKeys = *req_keys
		}
	}
}
