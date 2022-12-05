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
				bpt_json, err := json.Marshal("L")
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
				agentid := d.SetEmptyAgentinfo(csperf.AgentID)
				perf_table, dbtype := d.GetTableinfo("realtimeperf")
				cpu_table, _ := d.GetTableinfo("realtimecpu")

				if dbtype == "pg" {
					lrtp := data.Lastrealtimeperf{}
					perf := data.RealtimeperfPg{}
					cpu := data.RealtimecpuPg{}

					d.SetPerfPg(csperf, agentid, &lrtp, &perf, &cpu)

					if tcpRequestKeys.Code == app.LASTPERF_CODE {
						bpt_json, err := json.Marshal(lrtp.GetString())
						if err != nil {
							log.Println("JSON Data Conversion error")
						}
						tcpResponseData <- bpt_json
					}
					if tcpRequestKeys.Code == app.BASIC_CODE {
						perf_json, err := json.Marshal(perf.GetString())
						if err != nil {
							log.Println("JSON Data Conversion error")
						}
						tcpResponseData <- perf_json
					}
					if tcpRequestKeys.Code == app.CPU_CODE {
						cpu_json, err := json.Marshal(cpu.GetString())
						if err != nil {
							log.Println("JSON Data Conversion error")
						}
						tcpResponseData <- cpu_json
					}

					d.InsertPerfPg(&lrtp, &perf, &cpu, agentid, perf_table, cpu_table)
				} else {
					lrtp := data.Lastrealtimeperf{}
					perf := data.RealtimeperfTs{}
					cpu := data.RealtimecpuTs{}

					d.SetPerfTs(csperf, &lrtp, &perf, &cpu, agentid)

					if tcpRequestKeys.Code == app.LASTPERF_CODE {
						bpt_json, err := json.Marshal("L")
						if err != nil {
							log.Println("JSON Data Conversion error")
						}
						tcpResponseData <- bpt_json
					}
					if tcpRequestKeys.Code == app.BASIC_CODE {
						perf_json, err := json.Marshal("B")
						if err != nil {
							log.Println("JSON Data Conversion error")
						}
						tcpResponseData <- perf_json
					}
					if tcpRequestKeys.Code == app.CPU_CODE {
						cpu_json, err := json.Marshal("C")
						if err != nil {
							log.Println("JSON Data Conversion error")
						}
						tcpResponseData <- cpu_json
					}

					d.InsertPerfTs(&lrtp, &perf, &cpu, agentid, perf_table, cpu_table)
				}
			}
		case cspid := <-ch.ConsumerData.Realtimepid:
			for _, d := range db_handler {
				// Check Agent
				agentid := d.SetEmptyAgentinfo(cspid.AgentID)
				pid_table, dbtype := d.GetTableinfo("realtimepid")
				proc_table, _ := d.GetTableinfo("realtimeproc")

				if dbtype == "pg" {
					pid := data.RealtimepidPgArr{}
					proc := data.RealtimeprocPgArr{}

					d.SetPidPg(cspid, &pid, &proc, agentid)
					d.InsertPidPg(&pid, &proc, pid_table, proc_table)
				} else {
					pid := data.RealtimepidTsArr{}
					proc := data.RealtimeprocTsArr{}

					d.SetPidTs(cspid, &pid, &proc, agentid)
					d.InsertPidTs(&pid, &proc, pid_table, proc_table)
				}
			}
		case csdisk := <-ch.ConsumerData.Realtimedisk:
			for _, d := range db_handler {
				agentid := d.SetEmptyAgentinfo(csdisk.AgentID)
				tablename, dbtype := d.GetTableinfo("realtimedisk")

				if dbtype == "pg" {
					disk := data.RealtimediskPgArr{}

					d.SetDiskPg(csdisk, &disk, agentid)

					if tcpRequestKeys.Code == app.DISK_CODE {
						disk_json, err := json.Marshal("D")
						if err != nil {
							log.Println("JSON Data Conversion error")
						}
						tcpResponseData <- disk_json
					}

					d.InsertDiskPg(&disk, tablename)
				} else {
					disk := data.RealtimediskTsArr{}

					d.SetDiskTs(csdisk, &disk, agentid)

					if tcpRequestKeys.Code == app.DISK_CODE {
						disk_json, err := json.Marshal("D")
						if err != nil {
							log.Println("JSON Data Conversion error")
						}
						tcpResponseData <- disk_json
					}

					d.InsertDiskTs(&disk, tablename)
				}
			}
		case csnet := <-ch.ConsumerData.Realtimenet:
			for _, d := range db_handler {
				agentid := d.SetEmptyAgentinfo(csnet.AgentID)
				tablename, dbtype := d.GetTableinfo("realtimenet")

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

					d.InsertNetPg(&net, tablename)
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

					d.InsertNetTs(&net, tablename)
				}
			}
		case req_keys := <-tcpRequestChan:
			//fmt.Printf("main %v\n", req_keys)
			tcpRequestKeys = *req_keys
		}
	}
}
