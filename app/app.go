package app

import (
	"log"
	"manager/data"
	"net"
)

type ChannelStruct struct {
	ConsumerData        ConsumerStruct
	Agentinfo           chan *data.AgentinfoArr
	ChangeStateAgents   chan []int
	ChangeStateAgentStr chan string
	Lastrealtimeperf    chan *data.LastrealtimeperfArr
}

type ConsumerStruct struct {
	Host         chan *data.AgentHostAgentInfo
	Realtimeperf chan *data.AgentRealTimePerf
	Realtimepid  chan *data.AgentRealTimePID
	Realtimedisk chan *data.AgentRealTimeDisk
	Realtimenet  chan *data.AgentRealTimeNet
}

func (c *ChannelStruct) ChannelInit() {
	c.ConsumerData = ConsumerStruct{}
	c.ConsumerData.ConsumerInit()

	c.Agentinfo = make(chan *data.AgentinfoArr)
	c.ChangeStateAgents = make(chan []int)
	c.ChangeStateAgentStr = make(chan string)
	c.Lastrealtimeperf = make(chan *data.LastrealtimeperfArr)
}

func (c *ConsumerStruct) ConsumerInit() {
	c.Host = make(chan *data.AgentHostAgentInfo)
	c.Realtimeperf = make(chan *data.AgentRealTimePerf)
	c.Realtimepid = make(chan *data.AgentRealTimePID)
	c.Realtimedisk = make(chan *data.AgentRealTimeDisk)
	c.Realtimenet = make(chan *data.AgentRealTimeNet)
}

func TcpSend(data []byte) {
	client, err := net.Dial("tcp", "localhost:8088")
	if err != nil {
		log.Println("TCP Connection Error")
		return
	}
	defer client.Close()

	func(c net.Conn) {
		_, err := c.Write(data)
		if err != nil {
			log.Println(err)
			return
		}
	}(client)

}
