package main

import (
	"fmt"
	"manager/app"
	"manager/data"
	"time"
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

	current_time := app.ConsumerTime{}
	previous_time := app.ConsumerTime{}
	perf_arr := make([]data.AgentRealTimePerf, 0)
	pid_arr := make([]data.AgentRealTimePID, 0)
	disk_arr := make([]data.AgentRealTimeDisk, 0)
	net_arr := make([]data.AgentRealTimeNet, 0)

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
				//fmt.Printf("change state before %d %d %d\n", idx, len(net_arr), time.Now().UnixMicro())
				if tcpRequestKeys.IsDataMapping(app.HOST_CODE) {
					tcpResponseData <- app.ConvertJson(app.HOST_CODE, state_agent_str)
				}

				d.DemoHostStateChange(state_agent_str)
				//fmt.Printf("change state after %d %d %d\n", idx, len(net_arr), time.Now().UnixMicro())
			}
		case lrtp := <-ch.Lastrealtimeperf:
			if tcpRequestKeys.IsDataMapping(app.LASTPERF_CODE) {
				for _, l := range lrtp.GetArrString() {
					tcpResponseData <- app.ConvertJson(app.LASTPERF_CODE, l)
				}
			}
			for _, d := range db_handler {
				//fmt.Printf("lastrealtimeperf before %v %d %d\n", idx, len(net_arr), time.Now().UnixMicro())
				d.DemoBptUpdate(lrtp)
				//fmt.Printf("lastrealtimeperf after %v %d %d\n", idx, len(net_arr), time.Now().UnixMicro())
			}
		case cshost := <-ch.ConsumerData.Host:
			for idx, d := range db_handler {
				agentinfo_arr := data.AgentinfoArr{}
				//fmt.Printf("agent host before %v %d %d\n", idx, len(net_arr), time.Now().UnixMicro())

				d.SetHost(cshost, &agentinfo_arr)
				//fmt.Printf("agent host after %v %d %d\n", idx, len(net_arr), time.Now().UnixMicro())

				// TCP 데이터는 1회만 넘기도록 해야 함
				if idx == 0 && tcpRequestKeys.IsDataMapping(app.HOST_CODE) {
					for _, a := range agentinfo_arr.GetArrString() {
						tcpResponseData <- app.ConvertJson(app.HOST_CODE, a)
					}
				}
			}
		case csperf := <-ch.ConsumerData.Realtimeperf:
			current_time.Host = int64(time.Unix(time.Now().Unix(), 0).Unix() / 2)

			if len(perf_arr) > 0 && current_time.Host > previous_time.Host {
				//fmt.Printf("perf %d\n", len(perf_arr))
				for idx, d := range db_handler {
					//agentid := d.GetAgentId(csperf.AgentID)
					perf_arr = app.RemoveDuplicate(perf_arr).([]data.AgentRealTimePerf)

					dbtype := d.GetTabletype("realtimeperf")

					if dbtype == "pg" {
						lrtp := data.LastrealtimeperfArray{}
						perf := data.RealtimeperfPgArray{}
						cpu := data.RealtimecpuPgArray{}

						fmt.Printf("realtimeperf set %v %d %d\n", idx, len(perf_arr), time.Now().UnixMicro())
						d.SetPerfArray(&perf_arr, &lrtp, &perf, &cpu)

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
						fmt.Printf("realtimeperf before %v %d %d\n", idx, len(perf_arr), time.Now().UnixMicro())
						d.InsertTableArray(dbtype, &lrtp, &perf, &cpu)
						fmt.Printf("realtimeperf after %v %d %d\n", idx, len(perf_arr), time.Now().UnixMicro())
					} else {
						lrtp := data.LastrealtimeperfArray{}
						perf := data.RealtimeperfTsArray{}
						cpu := data.RealtimecpuTsArray{}

						fmt.Printf("realtimeperf set %v %d %d\n", idx, len(perf_arr), time.Now().UnixMicro())
						d.SetPerfArray(&perf_arr, &lrtp, &perf, &cpu)
						fmt.Printf("realtimeperf before %v %d %d\n", idx, len(perf_arr), time.Now().UnixMicro())
						d.InsertTableArray(dbtype, &lrtp, &perf, &cpu)
						fmt.Printf("realtimeperf after %v %d %d\n", idx, len(perf_arr), time.Now().UnixMicro())
					}
				}

				perf_arr = nil
				previous_time.Host = current_time.Host

			} else {
				perf_arr = append(perf_arr, *csperf)
			}
		case cspid := <-ch.ConsumerData.Realtimepid:
			current_time.Pid = int64(time.Unix(time.Now().Unix(), 0).Unix() / 2)

			if len(pid_arr) > 0 && current_time.Pid > previous_time.Pid {
				//fmt.Printf("pid %d\n", len(pid_arr))
				pid_arr = app.RemoveDuplicate(pid_arr).([]data.AgentRealTimePID)

				for idx, d := range db_handler {
					// Check Agent
					dbtype := d.GetTabletype("realtimepid")

					if dbtype == "pg" {
						pid := data.RealtimepidPgArray{}
						proc := data.RealtimeprocPgArray{}

						fmt.Printf("realtimepid set %v %d %d\n", idx, len(pid_arr), time.Now().UnixMicro())
						d.SetPidArray(&pid_arr, &pid, &proc)
						fmt.Printf("realtimepid before %v %d %d\n", idx, len(pid_arr), time.Now().UnixMicro())
						d.InsertTableArray(dbtype, &pid, &proc)
						fmt.Printf("realtimepid after %v %d %d\n", idx, len(pid_arr), time.Now().UnixMicro())
					} else {
						pid := data.RealtimepidTsArray{}
						proc := data.RealtimeprocTsArray{}

						fmt.Printf("realtimepid set %v %d %d\n", idx, len(pid_arr), time.Now().UnixMicro())
						d.SetPidArray(&pid_arr, &pid, &proc)
						fmt.Printf("realtimepid before %v %d %d\n", idx, len(pid_arr), time.Now().UnixMicro())
						d.InsertTableArray(dbtype, &pid, &proc)
						fmt.Printf("realtimepid after %v %d %d\n", idx, len(pid_arr), time.Now().UnixMicro())
					}
				}

				pid_arr = nil
				previous_time.Pid = current_time.Pid
			} else {
				pid_arr = append(pid_arr, *cspid)
			}
		case csdisk := <-ch.ConsumerData.Realtimedisk:
			current_time.Disk = int64(time.Unix(time.Now().Unix(), 0).Unix() / 2)

			if len(disk_arr) > 0 && current_time.Disk > previous_time.Disk {
				//fmt.Printf("disk %d\n", len(disk_arr))
				disk_arr = app.RemoveDuplicate(disk_arr).([]data.AgentRealTimeDisk)

				for idx, d := range db_handler {
					dbtype := d.GetTabletype("realtimedisk")

					if dbtype == "pg" {
						disk := data.RealtimediskPgArray{}

						fmt.Printf("realtimedisk set %v %d %d\n", idx, len(disk_arr), time.Now().UnixMicro())
						d.SetDiskArray(&disk_arr, &disk)

						if idx == 0 && tcpRequestKeys.IsDataMapping(app.DISK_CODE) {
							tcpResponseData <- app.ConvertJson(app.DISK_CODE, disk.GetString())
						}

						fmt.Printf("realtimedisk before %v %d %d\n", idx, len(disk_arr), time.Now().UnixMicro())
						d.InsertTableArray(dbtype, &disk)
						fmt.Printf("realtimedisk after %v %d %d\n", idx, len(disk_arr), time.Now().UnixMicro())
					} else {
						disk := data.RealtimediskTsArray{}

						fmt.Printf("realtimedisk set %v %d %d\n", idx, len(disk_arr), time.Now().UnixMicro())
						d.SetDiskArray(&disk_arr, &disk)
						fmt.Printf("realtimedisk before %v %d %d\n", idx, len(disk_arr), time.Now().UnixMicro())
						d.InsertTableArray(dbtype, &disk)
						fmt.Printf("realtimedisk after %v %d %d\n", idx, len(disk_arr), time.Now().UnixMicro())
					}
				}

				disk_arr = nil
				previous_time.Disk = current_time.Disk
			} else {
				disk_arr = append(disk_arr, *csdisk)
			}
		case csnet := <-ch.ConsumerData.Realtimenet:
			current_time.Net = int64(time.Unix(time.Now().Unix(), 0).Unix() / 2)

			if len(net_arr) > 0 && current_time.Net > previous_time.Pid {
				//fmt.Printf("net %d\n", len(net_arr))
				net_arr = app.RemoveDuplicate(net_arr).([]data.AgentRealTimeNet)

				for idx, d := range db_handler {
					dbtype := d.GetTabletype("realtimenet")

					if dbtype == "pg" {
						net := data.RealtimenetPgArray{}

						fmt.Printf("realtimenet set %v %d %d\n", idx, len(net_arr), time.Now().UnixMicro())
						d.SetNetArray(&net_arr, &net)

						if idx == 0 && tcpRequestKeys.IsDataMapping(app.NET_CODE) {
							tcpResponseData <- app.ConvertJson(app.NET_CODE, net.GetString())
						}

						fmt.Printf("realtimenet before %v %d %d\n", idx, len(net_arr), time.Now().UnixMicro())
						d.InsertTableArray(dbtype, &net)
						fmt.Printf("realtimenet after %v %d %d\n", idx, len(net_arr), time.Now().UnixMicro())
					} else {
						net := data.RealtimenetTsArray{}

						fmt.Printf("realtimenet set %v %d %d\n", idx, len(net_arr), time.Now().UnixMicro())
						d.SetNetArray(&net_arr, &net)
						fmt.Printf("realtimenet before %v %d %d\n", idx, len(net_arr), time.Now().UnixMicro())
						d.InsertTableArray(dbtype, &net)
						fmt.Printf("realtimenet after %v %d %d\n", idx, len(net_arr), time.Now().UnixMicro())
					}
				}

				net_arr = nil
				previous_time.Net = current_time.Net
			} else {
				net_arr = append(net_arr, *csnet)
			}
		case req_keys := <-tcpRequestChan:
			//fmt.Printf("main %v\n", req_keys)
			tcpRequestKeys = *req_keys
		}
	}
}
