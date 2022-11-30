package app

import (
	"log"
	"manager/data"
	"net"
)

type ChannelStruct struct {
	ConsumerData        chan string
	Agentinfo           chan *data.AgentinfoArr
	ChangeStateAgents   chan []int
	ChangeStateAgentStr chan string
	Lastrealtimeperf    chan *data.LastrealtimeperfArr
}

func (c *ChannelStruct) ChannelInit() {
	c.ConsumerData = make(chan string)
	c.Agentinfo = make(chan *data.AgentinfoArr)
	c.ChangeStateAgents = make(chan []int)
	c.ChangeStateAgentStr = make(chan string)
	c.Lastrealtimeperf = make(chan *data.LastrealtimeperfArr)
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
