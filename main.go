package main

import (
	"encoding/json"
	"log"
	"manager/app"
	"manager/data"
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
	//consumer := app.KafkaConsumerControllerInit(&kafkaconfig, "host")
	//paritionList := app.KafkaConsumerControllerPartition(&kafkaconfig, consumer, "host")

	go app.TcpProcessing(tcpRequestChan, tcpResponseData, config.ApiServerInfo)

	// go app.ConsumerHost(ch.ConsumerData.Host, kafkaconfig, paritionList)
	// go app.ConsumerPerf(ch.ConsumerData.Realtimeperf, kafkaconfig, paritionList)
	// go app.ConsumerPid(ch.ConsumerData.Realtimepid, kafkaconfig, paritionList)
	// go app.ConsumerDisk(ch.ConsumerData.Realtimedisk, kafkaconfig, paritionList)
	// go app.ConsumerNet(ch.ConsumerData.Realtimenet, kafkaconfig, paritionList)

	app.ConsumerHostGroup(ch.ConsumerData.Host, kafkaconfig)
	app.ConsumerPerfGroup(ch.ConsumerData.Realtimeperf, kafkaconfig)
	app.ConsumerPIDGroup(ch.ConsumerData.Realtimepid, kafkaconfig)
	app.ConsumerDiskGroup(ch.ConsumerData.Realtimedisk, kafkaconfig)
	app.ConsumerNetGroup(ch.ConsumerData.Realtimenet, kafkaconfig)

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
				app.ErrorJson(err)
				tcpResponseData <- agent_json
			}
			for _, d := range db_handler {
				d.DemoHostStateChange(state_agent_str)
			}
		case lrtp := <-ch.Lastrealtimeperf:
			if tcpRequestKeys.Code == app.LASTPERF_CODE {
				bpt_json, err := json.Marshal("L")
				app.ErrorJson(err)
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
			for idx, d := range db_handler {
				agentid := d.SetEmptyAgentinfo(csperf.AgentID)
				dbtype := d.GetTabletype("realtimeperf")

				if dbtype == "pg" {
					lrtp := data.Lastrealtimeperf{}
					perf := data.RealtimeperfPg{}
					cpu := data.RealtimecpuPg{}

					d.SetPerf(csperf, agentid, &lrtp, &perf, &cpu)

					// TCP 데이터는 1회만 넘기도록 해야 함
					if idx == 0 {
						if tcpRequestKeys.Code == app.LASTPERF_CODE {
							bpt_json, err := json.Marshal(lrtp.GetString())
							app.ErrorJson(err)
							tcpResponseData <- bpt_json
						}
						if tcpRequestKeys.Code == app.BASIC_CODE {
							perf_json, err := json.Marshal(perf.GetString())
							app.ErrorJson(err)
							tcpResponseData <- perf_json
						}
						if tcpRequestKeys.Code == app.CPU_CODE {
							cpu_json, err := json.Marshal(cpu.GetString())
							app.ErrorJson(err)
							tcpResponseData <- cpu_json
						}
					}

					d.InsertPerf(agentid, &lrtp, &perf, &cpu)
				} else {
					lrtp := data.Lastrealtimeperf{}
					perf := data.RealtimeperfTs{}
					cpu := data.RealtimecpuTs{}

					d.SetPerf(csperf, agentid, &lrtp, &perf, &cpu)
					d.InsertPerf(agentid, &lrtp, &perf, &cpu)
				}
			}
		case cspid := <-ch.ConsumerData.Realtimepid:
			for _, d := range db_handler {
				// Check Agent
				agentid := d.SetEmptyAgentinfo(cspid.AgentID)
				dbtype := d.GetTabletype("realtimepid")

				if dbtype == "pg" {
					pid := data.RealtimepidPgArr{}
					proc := data.RealtimeprocPgArr{}

					d.SetPidPg(cspid, &pid, &proc, agentid)
					d.InsertPidPg(&pid, &proc)
				} else {
					pid := data.RealtimepidTsArr{}
					proc := data.RealtimeprocTsArr{}

					d.SetPidTs(cspid, &pid, &proc, agentid)
					d.InsertPidTs(&pid, &proc)
				}
			}
		case csdisk := <-ch.ConsumerData.Realtimedisk:
			for _, d := range db_handler {
				agentid := d.SetEmptyAgentinfo(csdisk.AgentID)
				dbtype := d.GetTabletype("realtimedisk")

				if dbtype == "pg" {
					disk := data.RealtimediskPgArr{}

					d.SetDiskPg(csdisk, &disk, agentid)

					if tcpRequestKeys.Code == app.DISK_CODE {
						disk_json, err := json.Marshal("D")
						app.ErrorJson(err)
						tcpResponseData <- disk_json
					}

					d.InsertDiskPg(&disk)
				} else {
					disk := data.RealtimediskTsArr{}

					d.SetDiskTs(csdisk, &disk, agentid)

					if tcpRequestKeys.Code == app.DISK_CODE {
						disk_json, err := json.Marshal("D")
						app.ErrorJson(err)
						tcpResponseData <- disk_json
					}

					d.InsertDiskTs(&disk)
				}
			}
		case csnet := <-ch.ConsumerData.Realtimenet:
			for _, d := range db_handler {
				agentid := d.SetEmptyAgentinfo(csnet.AgentID)
				dbtype := d.GetTabletype("realtimenet")

				if dbtype == "pg" {
					net := data.RealtimenetPgArr{}

					d.SetNetPg(csnet, &net, agentid)

					if tcpRequestKeys.Code == app.NET_CODE {
						net_json, err := json.Marshal("N")
						if err != nil {
							log.Println("JSON Data Conversion error")
						}
						tcpResponseData <- net_json
					}

					d.InsertNetPg(&net)
				} else {
					net := data.RealtimenetTsArr{}

					d.SetNetTs(csnet, &net, agentid)

					if tcpRequestKeys.Code == app.NET_CODE {
						net_json, err := json.Marshal("N")
						if err != nil {
							log.Println("JSON Data Conversion error")
						}
						tcpResponseData <- net_json
					}

					d.InsertNetTs(&net)
				}
			}
		case req_keys := <-tcpRequestChan:
			//fmt.Printf("main %v\n", req_keys)
			tcpRequestKeys = *req_keys
		}
	}
}
