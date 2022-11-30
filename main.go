package main

import (
	"encoding/json"
	"fmt"
	"log"
	"manager/app"
)

func main() {
	config := app.GetConfig("config.yml")
	ch := app.ChannelStruct{}
	ch.ChannelInit()

	db_handler := make([]app.DBHandler, 0)

	go app.ConsumerProcessing(ch.ConsumerData, config.Demo.Interval)
	go app.GetDemoAgentinfo(ch.Agentinfo, config.Demo.HostCount)
	go app.GetDemoChangeStateAgent(ch.ChangeStateAgents, ch.ChangeStateAgentStr, config.Demo)
	go app.GetDemoLastrealtimeperf(ch.Lastrealtimeperf, config.Demo)

	agentinfo := <-ch.Agentinfo
	for _, dbinfo := range config.Database {
		db_handler = append(db_handler, *app.DBInit(dbinfo, agentinfo))
	}

	for {
		select {
		case <-ch.ChangeStateAgents:
			state_agent_str := <-ch.ChangeStateAgentStr
			agent_json, err := json.Marshal(state_agent_str)
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
		case con := <-ch.ConsumerData:
			fmt.Println(con)
		}
	}
}
