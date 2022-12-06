package main

import (
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
			// 이 부분 TCP 데이터는 일단 넘기지 않고 추후 검토
			for _, d := range db_handler {
				d.DemoHostStateChange(state_agent_str)
			}
		case lrtp := <-ch.Lastrealtimeperf:
			if tcpRequestKeys.Code == app.LASTPERF_CODE {
				for _, l := range lrtp.GetArrString() {
					tcpResponseData <- app.ConvertJson(tcpRequestKeys.Code, l)
				}
			}
			for _, d := range db_handler {
				d.DemoBptUpdate(lrtp)
			}
		case cshost := <-ch.ConsumerData.Host:
			for idx, d := range db_handler {
				agentinfo_arr := data.AgentinfoArr{}
				d.SetHost(cshost, &agentinfo_arr)

				// TCP 데이터는 1회만 넘기도록 해야 함
				if idx == 0 && tcpRequestKeys.Code == app.HOST_CODE {
					for _, a := range agentinfo_arr.GetArrString() {
						tcpResponseData <- app.ConvertJson(tcpRequestKeys.Code, a)
					}
				}
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
							tcpResponseData <- app.ConvertJson(tcpRequestKeys.Code, lrtp.GetString())
						}
						if tcpRequestKeys.Code == app.BASIC_CODE {
							tcpResponseData <- app.ConvertJson(tcpRequestKeys.Code, perf.GetString())
						}
						if tcpRequestKeys.Code == app.CPU_CODE {
							tcpResponseData <- app.ConvertJson(tcpRequestKeys.Code, cpu.GetString())
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

					d.SetPid(cspid, agentid, &pid, &proc)
					d.InsertTableArr(&pid, &proc)
				} else {
					pid := data.RealtimepidTsArr{}
					proc := data.RealtimeprocTsArr{}

					d.SetPid(cspid, agentid, &pid, &proc)
					d.InsertTableArr(&pid, &proc)
				}
			}
		case csdisk := <-ch.ConsumerData.Realtimedisk:
			for idx, d := range db_handler {
				agentid := d.SetEmptyAgentinfo(csdisk.AgentID)
				dbtype := d.GetTabletype("realtimedisk")

				if dbtype == "pg" {
					disk := data.RealtimediskPgArr{}

					d.SetDisk(csdisk, agentid, &disk)

					if idx == 0 && tcpRequestKeys.Code == app.DISK_CODE {
						for _, a := range disk.GetArrString() {
							tcpResponseData <- app.ConvertJson(tcpRequestKeys.Code, a)
						}
					}

					d.InsertTableArr(&disk)
				} else {
					disk := data.RealtimediskTsArr{}

					d.SetDisk(csdisk, agentid, &disk)
					d.InsertTableArr(&disk)
				}
			}
		case csnet := <-ch.ConsumerData.Realtimenet:
			for idx, d := range db_handler {
				agentid := d.SetEmptyAgentinfo(csnet.AgentID)
				dbtype := d.GetTabletype("realtimenet")

				if dbtype == "pg" {
					net := data.RealtimenetPgArr{}

					d.SetNet(csnet, agentid, &net)

					if idx == 0 && tcpRequestKeys.Code == app.NET_CODE {
						for _, a := range net.GetArrString() {
							tcpResponseData <- app.ConvertJson(tcpRequestKeys.Code, a)
						}
					}

					d.InsertTableArr(&net)
				} else {
					net := data.RealtimenetTsArr{}

					d.SetNet(csnet, agentid, &net)
					d.InsertTableArr(&net)
				}
			}
		case req_keys := <-tcpRequestChan:
			//fmt.Printf("main %v\n", req_keys)
			tcpRequestKeys = *req_keys
		}
	}
}
