package app

import (
	"encoding/json"
	"manager/data"
	"time"
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

type RealData struct {
	Code uint32 `json:"code"`
	Data string `json:"data"`
}

type ConsumerTime struct {
	Host time.Time
	Pid  time.Time
	Disk time.Time
	Net  time.Time
}

const (
	DATAKEY_CODE  = 0x00000001
	HOST_CODE     = 0x00000002
	LASTPERF_CODE = 0x00000004
	BASIC_CODE    = 0x00000008
	CPU_CODE      = 0x00000010
	MEM_CODE      = 0x00000020
	NET_CODE      = 0x00000040
	DISK_CODE     = 0x00000080
)

type DataKey struct {
	Code uint32  `json:"code"`
	Key  Bitmask `json:"key"`
}

type DataCode struct {
	Code Bitmask `json:"code"`
}

type Bitmask uint32

func (value Bitmask) IsSet(key Bitmask) bool {
	return value&key != 0
}

func (d DataKey) IsDataMapping(key Bitmask) bool {
	return d.Code == DATAKEY_CODE && d.Key.IsSet(key)
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

func ConvertJson(code uint32, rdata string) []byte {
	realData := RealData{}
	realData.Code = code
	realData.Data = rdata

	json_data, err := json.Marshal(realData)
	ErrorJson(err, code)

	return json_data
}
