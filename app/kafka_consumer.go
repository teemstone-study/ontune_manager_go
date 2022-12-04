package app

import (
	"encoding/json"
	"fmt"
	"manager/data"
	"time"
)

const goroutine int = 1

func ConsumerHost(cshost chan<- *data.AgentHostAgentInfo, config SettingKafka, plist []int32) {
	for cdata := range plist {
		consumer := KafkaConsumerControllerInit(&config, "host")
		cPartition := KafkaConsumerControllerGetPartitionConsumer(&config, consumer, "host", int32(cdata))
		for i := 0; i < goroutine; i++ {
			go func() {
				for {
					fmt.Printf("host before %d", time.Now().UnixMicro())
					msg := <-cPartition.Messages()
					agenthostinfo := data.AgentHostAgentInfo{}
					err := json.Unmarshal(msg.Value, &agenthostinfo)
					if err == nil {
						cshost <- &agenthostinfo
					}
					fmt.Printf("host after %d", time.Now().UnixMicro())
				}
			}()
		}
	}
}

func ConsumerPerf(csperf chan<- *data.AgentRealTimePerf, config SettingKafka, plist []int32) {
	startTime := time.Now()
	for cdata := range plist {
		perfconsumer := KafkaConsumerControllerInit(&config, "realtimeperf")
		perfcPartition := KafkaConsumerControllerGetPartitionConsumer(&config, perfconsumer, "realtimeperf", int32(cdata))
		for i := 0; i < goroutine; i++ {
			go func() {
				for {
					msg := <-perfcPartition.Messages()
					realtimeperfData := data.AgentRealTimePerf{}
					err := json.Unmarshal(msg.Value, &realtimeperfData)
					if err == nil {
						csperf <- &realtimeperfData
						fmt.Println(time.Now().Sub(startTime))
					}
				}
			}()
		}
	}
}

func ConsumerPid(cspid chan<- *data.AgentRealTimePID, config SettingKafka, plist []int32) {
	for cdata := range plist {
		pidconsumer := KafkaConsumerControllerInit(&config, "realtimepid")
		pidcPartition := KafkaConsumerControllerGetPartitionConsumer(&config, pidconsumer, "realtimepid", int32(cdata))
		for i := 0; i < goroutine; i++ {
			go func() {
				for {
					msg := <-pidcPartition.Messages()
					realTimePIDreceive := data.AgentRealTimePID{}
					err := json.Unmarshal(msg.Value, &realTimePIDreceive)
					if err == nil {
						cspid <- &realTimePIDreceive
					}
				}
			}()
		}
	}
}

func ConsumerDisk(csdisk chan<- *data.AgentRealTimeDisk, config SettingKafka, plist []int32) {
	for cdata := range plist {
		diskconsumer := KafkaConsumerControllerInit(&config, "realtimedisk")
		diskcPartition := KafkaConsumerControllerGetPartitionConsumer(&config, diskconsumer, "realtimedisk", int32(cdata))
		for i := 0; i < goroutine; i++ {
			go func() {
				for {
					msg := <-diskcPartition.Messages()
					realTimeDiskreceive := data.AgentRealTimeDisk{}
					err := json.Unmarshal(msg.Value, &realTimeDiskreceive)
					if err == nil {
						csdisk <- &realTimeDiskreceive
					}
				}
			}()
		}
	}
}

func ConsumerNet(csnet chan<- *data.AgentRealTimeNet, config SettingKafka, plist []int32) {
	for cdata := range plist {
		netconsumer := KafkaConsumerControllerInit(&config, "realtimenet")
		netcPartition := KafkaConsumerControllerGetPartitionConsumer(&config, netconsumer, "realtimenet", int32(cdata))
		for i := 0; i < goroutine; i++ {
			go func() {
				for {
					msg := <-netcPartition.Messages()
					realTimeNewreceive := data.AgentRealTimeNet{}
					err := json.Unmarshal(msg.Value, &realTimeNewreceive)
					if err == nil {
						csnet <- &realTimeNewreceive
					}
				}
			}()
		}
	}
}
