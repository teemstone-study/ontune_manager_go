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

	tcpRequestChan := make(chan *app.DataKey)
	tcpRequestKeys := app.DataKey{}
	tcpResponseData := make(chan []byte)

	ch := app.ChannelStruct{}
	ch.ChannelInit()

	go app.TcpProcessing(tcpRequestChan, tcpResponseData, config.ApiServerInfo)

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
			// 이 부분 TCP 데이터는 일단 넘기도록 하나, 아래의 Host 정보와 넘기는 형식은 다름
			// 여기에서는 변경될 Agent ID만 넘기는 형태가 됨
			for _, d := range db_handler {
				//fmt.Printf("change state before %d %d\n", idx, time.Now().UnixMicro())
				if tcpRequestKeys.IsDataMapping(app.HOST_CODE) {
					tcpResponseData <- app.ConvertJson(app.HOST_CODE, state_agent_str)
				}

				d.DemoHostStateChange(state_agent_str)
				//fmt.Printf("change state after %d %d\n", idx, time.Now().UnixMicro())
			}
		case lrtp := <-ch.Lastrealtimeperf:
			if tcpRequestKeys.IsDataMapping(app.LASTPERF_CODE) {
				for _, l := range lrtp.GetArrString() {
					tcpResponseData <- app.ConvertJson(app.LASTPERF_CODE, l)
				}
			}
			for _, d := range db_handler {
				//fmt.Printf("lastrealtimeperf before %v %d\n", idx, time.Now().UnixMicro())
				d.DemoBptUpdate(lrtp)
				//fmt.Printf("lastrealtimeperf after %v %d\n", idx, time.Now().UnixMicro())
			}
		case cshost := <-ch.ConsumerData.Host:
			for idx, d := range db_handler {
				agentinfo_arr := data.AgentinfoArr{}
				//fmt.Printf("agent host before %v %d\n", idx, time.Now().UnixMicro())
				d.SetHost(cshost, &agentinfo_arr)
				//fmt.Printf("agent host after %v %d\n", idx, time.Now().UnixMicro())

				// TCP 데이터는 1회만 넘기도록 해야 함
				if idx == 0 && tcpRequestKeys.IsDataMapping(app.HOST_CODE) {
					for _, a := range agentinfo_arr.GetArrString() {
						tcpResponseData <- app.ConvertJson(app.HOST_CODE, a)
					}
				}
			}
		case csperf := <-ch.ConsumerData.Realtimeperf:
			for idx, d := range db_handler {
				agentid := d.GetAgentId(csperf.AgentID)
				dbtype := d.GetTabletype("realtimeperf")

				if dbtype == "pg" {
					lrtp := data.Lastrealtimeperf{}
					perf := data.RealtimeperfPg{}
					cpu := data.RealtimecpuPg{}

					//fmt.Printf("realtimeperf set %v %d\n", idx, time.Now().UnixMicro())
					d.SetPerf(csperf, agentid, &lrtp, &perf, &cpu)

					// TCP 데이터는 1회만 넘기도록 해야 함
					if idx == 0 {
						if tcpRequestKeys.IsDataMapping(app.LASTPERF_CODE) {
							tcpResponseData <- app.ConvertJson(app.LASTPERF_CODE, lrtp.GetString())
						}
						if tcpRequestKeys.IsDataMapping(app.BASIC_CODE) {
							tcpResponseData <- app.ConvertJson(app.BASIC_CODE, perf.GetString())
						}
						if tcpRequestKeys.IsDataMapping(app.CPU_CODE) {
							tcpResponseData <- app.ConvertJson(app.CPU_CODE, cpu.GetString())
						}
					}

					//fmt.Printf("realtimeperf before %v %d\n", idx, time.Now().UnixMicro())
					d.InsertPerf(agentid, &lrtp, &perf, &cpu)
					//fmt.Printf("realtimeperf after %v %d\n", idx, time.Now().UnixMicro())
				} else {
					lrtp := data.Lastrealtimeperf{}
					perf := data.RealtimeperfTs{}
					cpu := data.RealtimecpuTs{}

					//fmt.Printf("realtimeperf set %v %d\n", idx, time.Now().UnixMicro())
					d.SetPerf(csperf, agentid, &lrtp, &perf, &cpu)
					//fmt.Printf("realtimeperf before %v %d\n", idx, time.Now().UnixMicro())
					d.InsertPerf(agentid, &lrtp, &perf, &cpu)
					//fmt.Printf("realtimeperf after %v %d\n", idx, time.Now().UnixMicro())
				}
			}
		case cspid := <-ch.ConsumerData.Realtimepid:
			for _, d := range db_handler {
				// Check Agent
				agentid := d.GetAgentId(cspid.AgentID)
				dbtype := d.GetTabletype("realtimepid")

				if dbtype == "pg" {
					pid := data.RealtimepidPgArr{}
					proc := data.RealtimeprocPgArr{}

					//fmt.Printf("realtimepid set %v %d\n", idx, time.Now().UnixMicro())
					d.SetPid(cspid, agentid, &pid, &proc)
					//fmt.Printf("realtimepid before %v %d\n", idx, time.Now().UnixMicro())
					d.InsertTableArr(&pid, &proc)
					//fmt.Printf("realtimepid after %v %d\n", idx, time.Now().UnixMicro())
				} else {
					pid := data.RealtimepidTsArr{}
					proc := data.RealtimeprocTsArr{}

					//fmt.Printf("realtimepid set %v %d\n", idx, time.Now().UnixMicro())
					d.SetPid(cspid, agentid, &pid, &proc)
					//fmt.Printf("realtimepid before %v %d\n", idx, time.Now().UnixMicro())
					d.InsertTableArr(&pid, &proc)
					//fmt.Printf("realtimepid after %v %d\n", idx, time.Now().UnixMicro())
				}
			}
		case csdisk := <-ch.ConsumerData.Realtimedisk:
			for idx, d := range db_handler {
				agentid := d.GetAgentId(csdisk.AgentID)
				dbtype := d.GetTabletype("realtimedisk")

				if dbtype == "pg" {
					disk := data.RealtimediskPgArr{}

					//fmt.Printf("realtimedisk set %v %d\n", idx, time.Now().UnixMicro())
					d.SetDisk(csdisk, agentid, &disk)

					if idx == 0 && tcpRequestKeys.IsDataMapping(app.DISK_CODE) {
						for _, a := range disk.GetArrString() {
							tcpResponseData <- app.ConvertJson(app.DISK_CODE, a)
						}
					}

					//fmt.Printf("realtimedisk before %v %d\n", idx, time.Now().UnixMicro())
					d.InsertTableArr(&disk)
					//fmt.Printf("realtimedisk after %v %d\n", idx, time.Now().UnixMicro())
				} else {
					disk := data.RealtimediskTsArr{}

					//fmt.Printf("realtimedisk set %v %d\n", idx, time.Now().UnixMicro())
					d.SetDisk(csdisk, agentid, &disk)
					//fmt.Printf("realtimedisk before %v %d\n", idx, time.Now().UnixMicro())
					d.InsertTableArr(&disk)
					//fmt.Printf("realtimedisk after %v %d\n", idx, time.Now().UnixMicro())
				}
			}
		case csnet := <-ch.ConsumerData.Realtimenet:
			for idx, d := range db_handler {
				agentid := d.GetAgentId(csnet.AgentID)
				dbtype := d.GetTabletype("realtimenet")

				if dbtype == "pg" {
					net := data.RealtimenetPgArr{}

					//fmt.Printf("realtimenet set %v %d\n", idx, time.Now().UnixMicro())
					d.SetNet(csnet, agentid, &net)

					if idx == 0 && tcpRequestKeys.IsDataMapping(app.NET_CODE) {
						for _, a := range net.GetArrString() {
							tcpResponseData <- app.ConvertJson(app.NET_CODE, a)
						}
					}

					//fmt.Printf("realtimenet before %v %d\n", idx, time.Now().UnixMicro())
					d.InsertTableArr(&net)
					//fmt.Printf("realtimenet after %v %d\n", idx, time.Now().UnixMicro())
				} else {
					net := data.RealtimenetTsArr{}

					//fmt.Printf("realtimenet set %v %d\n", idx, time.Now().UnixMicro())
					d.SetNet(csnet, agentid, &net)
					//fmt.Printf("realtimenet before %v %d\n", idx, time.Now().UnixMicro())
					d.InsertTableArr(&net)
					//fmt.Printf("realtimenet after %v %d\n", idx, time.Now().UnixMicro())
				}
			}
		case req_keys := <-tcpRequestChan:
			////fmt.Printf("main %v\n", req_keys)
			tcpRequestKeys = *req_keys
		}
	}
}
