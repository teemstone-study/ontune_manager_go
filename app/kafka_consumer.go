package app

import (
	"encoding/json"
	"manager/data"
)

func ConsumerHost(cshost chan<- *data.AgentHostAgentInfo, config SettingKafka, plist []int32) {
	for {
		for cdata := range plist {
			consumer := KafkaConsumerControllerInit(&config, "host")
			cPartition := KafkaConsumerControllerGetPartitionConsumer(&config, consumer, "host", int32(cdata))
			msg := <-cPartition.Messages()
			agenthostinfo := data.AgentHostAgentInfo{}
			json.Unmarshal(msg.Value, &agenthostinfo)
			cshost <- &agenthostinfo
		}
	}
}

func ConsumerPerf(csperf chan<- *data.AgentRealTimePerf, config SettingKafka, plist []int32) {
	for {
		for cdata := range plist {
			perfconsumer := KafkaConsumerControllerInit(&config, "realtimeperf")
			perfcPartition := KafkaConsumerControllerGetPartitionConsumer(&config, perfconsumer, "realtimeperf", int32(cdata))
			msg := <-perfcPartition.Messages()
			realtimeperfData := data.AgentRealTimePerf{}
			json.Unmarshal(msg.Value, &realtimeperfData)
			csperf <- &realtimeperfData
		}
	}
}

func ConsumerPid(cspid chan<- *data.AgentRealTimePID, config SettingKafka, plist []int32) {
	for {
		for cdata := range plist {
			pidconsumer := KafkaConsumerControllerInit(&config, "realtimepid")
			pidcPartition := KafkaConsumerControllerGetPartitionConsumer(&config, pidconsumer, "realtimepid", int32(cdata))
			msg := <-pidcPartition.Messages()
			realTimePIDreceive := data.AgentRealTimePID{}
			json.Unmarshal(msg.Value, &realTimePIDreceive)
			cspid <- &realTimePIDreceive
		}
	}
}

func ConsumerDisk(csdisk chan<- *data.AgentRealTimeDisk, config SettingKafka, plist []int32) {
	for {
		for cdata := range plist {
			diskconsumer := KafkaConsumerControllerInit(&config, "realtimedisk")
			diskcPartition := KafkaConsumerControllerGetPartitionConsumer(&config, diskconsumer, "realtimedisk", int32(cdata))
			msg := <-diskcPartition.Messages()
			realTimeDiskreceive := data.AgentRealTimeDisk{}
			json.Unmarshal(msg.Value, &realTimeDiskreceive)
			csdisk <- &realTimeDiskreceive
		}
	}
}

func ConsumerNet(csnet chan<- *data.AgentRealTimeNet, config SettingKafka, plist []int32) {
	for {
		for cdata := range plist {
			netconsumer := KafkaConsumerControllerInit(&config, "realtimenet")
			netcPartition := KafkaConsumerControllerGetPartitionConsumer(&config, netconsumer, "realtimenet", int32(cdata))
			msg := <-netcPartition.Messages()
			realTimeNewreceive := data.AgentRealTimeNet{}
			json.Unmarshal(msg.Value, &realTimeNewreceive)
			csnet <- &realTimeNewreceive
		}
	}
}
