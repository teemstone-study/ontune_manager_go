package app

import (
	"context"
	"encoding/json"
	"manager/data"

	"gopkg.in/Shopify/sarama.v1"
)

func ConsumerHostGroup(cshost chan<- *data.AgentHostAgentInfo, config SettingKafka) {
	grouptype := kafkaHostGroup{Csperf: &cshost}
	groupconsumer := KafkaGroupConsumerControllerInit(&config, "hostgroup")
	ctx, _ := context.WithCancel(context.Background())
	go func() {
		for {
			groupconsumer.Consume(ctx, []string{"host"}, &grouptype)
		}
	}()
}

func ConsumerPerfGroup(cshost chan<- *data.AgentRealTimePerf, config SettingKafka) {
	grouptype := kafkarealtimeperfGroup{Csperf: &cshost}
	groupconsumer := KafkaGroupConsumerControllerInit(&config, "realtimeperfgroup")
	ctx, _ := context.WithCancel(context.Background())
	go func() {
		for {
			groupconsumer.Consume(ctx, []string{"realtimeperf"}, &grouptype)
		}
	}()
}

func ConsumerPIDGroup(cshost chan<- *data.AgentRealTimePID, config SettingKafka) {
	grouptype := kafkarealtimepidGroup{Csperf: &cshost}
	groupconsumer := KafkaGroupConsumerControllerInit(&config, "realtimepidgroup")
	ctx, _ := context.WithCancel(context.Background())
	go func() {
		for {
			groupconsumer.Consume(ctx, []string{"realtimepid"}, &grouptype)
		}
	}()
}

func ConsumerDiskGroup(cshost chan<- *data.AgentRealTimeDisk, config SettingKafka) {
	grouptype := kafkarealtimediskGroup{Csperf: &cshost}
	groupconsumer := KafkaGroupConsumerControllerInit(&config, "realtimediskgroup")
	ctx, _ := context.WithCancel(context.Background())
	go func() {
		for {
			groupconsumer.Consume(ctx, []string{"realtimedisk"}, &grouptype)
		}
	}()
}

func ConsumerNetGroup(cshost chan<- *data.AgentRealTimeNet, config SettingKafka) {
	grouptype := kafkarealtimenetGroup{Csperf: &cshost}
	groupconsumer := KafkaGroupConsumerControllerInit(&config, "realtimenetgroup")
	ctx, _ := context.WithCancel(context.Background())
	go func() {
		for {
			groupconsumer.Consume(ctx, []string{"realtimenet"}, &grouptype)
		}
	}()
}

type kafkaHostGroup struct {
	Csperf *chan<- *data.AgentHostAgentInfo
}

func (consumer *kafkaHostGroup) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *kafkaHostGroup) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *kafkaHostGroup) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			hostData := data.AgentHostAgentInfo{}
			err := json.Unmarshal(message.Value, &hostData)
			if err == nil {
				*consumer.Csperf <- &hostData
				session.MarkMessage(message, "")
			}

		case <-session.Context().Done():
			//fmt.Println("kafkaHostError :", session.Context().Err().Error())
			return nil
		}
	}
}

type kafkarealtimeperfGroup struct {
	Csperf *chan<- *data.AgentRealTimePerf
}

func (consumer *kafkarealtimeperfGroup) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *kafkarealtimeperfGroup) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *kafkarealtimeperfGroup) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			perfData := data.AgentRealTimePerf{}
			err := json.Unmarshal(message.Value, &perfData)
			if err == nil {
				*consumer.Csperf <- &perfData
				session.MarkMessage(message, "")
			}

		case <-session.Context().Done():
			//fmt.Println("kafkaPerfError :", session.Context().Err().Error())
			return nil
		}
	}
}

type kafkarealtimepidGroup struct {
	Csperf *chan<- *data.AgentRealTimePID
}

func (consumer *kafkarealtimepidGroup) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *kafkarealtimepidGroup) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *kafkarealtimepidGroup) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			hostData := data.AgentRealTimePID{}
			err := json.Unmarshal(message.Value, &hostData)
			if err == nil {
				*consumer.Csperf <- &hostData
				session.MarkMessage(message, "")
			}

		case <-session.Context().Done():
			//fmt.Println("kafkaPidError :", session.Context().Err().Error())
			return nil
		}
	}
}

type kafkarealtimediskGroup struct {
	Csperf *chan<- *data.AgentRealTimeDisk
}

func (consumer *kafkarealtimediskGroup) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *kafkarealtimediskGroup) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *kafkarealtimediskGroup) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			hostData := data.AgentRealTimeDisk{}
			err := json.Unmarshal(message.Value, &hostData)
			if err == nil {
				*consumer.Csperf <- &hostData
				session.MarkMessage(message, "")
			}

		case <-session.Context().Done():
			//fmt.Println("kafkaDiskError :", session.Context().Err().Error())
			return nil
		}
	}
}

type kafkarealtimenetGroup struct {
	Csperf *chan<- *data.AgentRealTimeNet
}

func (consumer *kafkarealtimenetGroup) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *kafkarealtimenetGroup) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *kafkarealtimenetGroup) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			hostData := data.AgentRealTimeNet{}
			err := json.Unmarshal(message.Value, &hostData)
			if err == nil {
				*consumer.Csperf <- &hostData
				session.MarkMessage(message, "")
			}

		case <-session.Context().Done():
			//fmt.Println("kafkaNetError :", session.Context().Err().Error())
			return nil
		}
	}
}
