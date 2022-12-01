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
	consumer := app.KafkaConsumerControllerInit(&kafkaconfig, "host")
	paritionList := app.KafkaConsumerControllerPartition(&kafkaconfig, consumer, "host")

	ch := app.ChannelStruct{}
	ch.ChannelInit()

	db_handler := make([]app.DBHandler, 0)

	go app.ConsumerHost(ch.ConsumerData.Host, kafkaconfig, paritionList)
	go app.ConsumerPerf(ch.ConsumerData.Realtimeperf, kafkaconfig, paritionList)
	go app.ConsumerPid(ch.ConsumerData.Realtimepid, kafkaconfig, paritionList)
	go app.ConsumerDisk(ch.ConsumerData.Realtimedisk, kafkaconfig, paritionList)
	go app.ConsumerNet(ch.ConsumerData.Realtimenet, kafkaconfig, paritionList)

	go app.GetDemoAgentinfo(ch.Agentinfo, config.Demo.HostCount)
	go app.GetDemoChangeStateAgent(ch.ChangeStateAgents, ch.ChangeStateAgentStr, config.Demo)
	go app.GetDemoLastrealtimeperf(ch.Lastrealtimeperf, config.Demo)

	agentinfo := <-ch.Agentinfo
	for _, dbinfo := range config.Database {
		db_handler = append(db_handler, *app.DBInit(dbinfo, agentinfo))
	}

	for {
		select {
		case state_agents := <-ch.ChangeStateAgents:
			state_agent_str := <-ch.ChangeStateAgentStr
			agent_json, err := json.Marshal(state_agents)
			if err != nil {
				log.Println("JSON Data Conversion Error")
			}
			go app.TcpSend(agent_json)
			for _, d := range db_handler {
				d.DemoHostStateChange(state_agent_str)
			}
		case lrtp := <-ch.Lastrealtimeperf:
			bpt_json, err := json.Marshal(lrtp)
			if err != nil {
				log.Println("JSON Data Conversion Error")
			}
			go app.TcpSend(bpt_json)
			for _, d := range db_handler {
				d.DemoBptUpdate(lrtp)
			}
		case cshost := <-ch.ConsumerData.Host:
			fmt.Println(cshost)
		case csperf := <-ch.ConsumerData.Realtimeperf:
			fmt.Println(csperf)
		case cspid := <-ch.ConsumerData.Realtimepid:
			fmt.Println(cspid)
		case csdisk := <-ch.ConsumerData.Realtimedisk:
			fmt.Println(csdisk)
		case csnet := <-ch.ConsumerData.Realtimenet:
			fmt.Println(csnet)
		}
	}
}
