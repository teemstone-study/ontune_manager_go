package app

import (
	"fmt"

	"gopkg.in/Shopify/sarama.v1"
)

type SettingKafka struct {
	KafkaServerAddr string
	KafkaServerPort string
}

var m_kafkaConfig *sarama.Config = nil
var m_kafkaClient *sarama.Client = nil
var m_kafkaProducer sarama.SyncProducer = nil
var m_moduleProcedureInit bool = false

func kafkaControllerInit(configValue *SettingKafka) {
	if m_kafkaConfig == nil || m_kafkaClient == nil {
		m_kafkaConfig = sarama.NewConfig()
		m_kafkaConfig.Version = sarama.V0_10_2_0
		m_kafkaConfig.Producer.Return.Successes = true
		m_kafkaConfig.Producer.Compression = sarama.CompressionSnappy
		m_kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll //무손실 옵션. 속도가 느릴수도?
		m_kafkaConfig.Producer.MaxMessageBytes = 4000000        //기본값 1000000. 4배로 변경해봄
		m_kafkaConfig.ChannelBufferSize = 4096
		m_kafkaConfig.Consumer.Fetch.Default = 100000
		//m_kafkaConfig.Consumer.Offsets.AutoCommit.Enable = true

		connectionString := []string{configValue.KafkaServerAddr + ":" + configValue.KafkaServerPort}
		kafkaClient, err := sarama.NewClient(connectionString, m_kafkaConfig)
		if err != nil {
			panic(err)
		}
		m_kafkaClient = &kafkaClient
	}
}

func KafkaProducerControllerInit(configValue *SettingKafka) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("kafka Open Error!")
			m_moduleProcedureInit = false
		}
	}()
	kafkaControllerInit(configValue)
	if m_kafkaProducer == nil {
		producer, err := sarama.NewSyncProducerFromClient(*m_kafkaClient)
		if err != nil {
			panic(err)
		}
		m_kafkaProducer = producer
	}
	m_moduleProcedureInit = true
}

func SendKafkaData(topicName string, keyName string, SendData []byte) {
	if !m_moduleProcedureInit {
		panic("kafka Producer Not Init")
	}

	msg := &sarama.ProducerMessage{
		Topic: topicName,
		Key:   sarama.StringEncoder(keyName),
		Value: sarama.ByteEncoder(SendData),
	}
	_, _, _ = m_kafkaProducer.SendMessage(msg)
}

func KafkaConsumerControllerPartition(configValue *SettingKafka, setconsumer sarama.Consumer, topicName string) []int32 {
	if setconsumer != nil && configValue != nil {
		partitions, err := setconsumer.Partitions(topicName)
		if err != nil {
			panic(err)
		} else {
			return partitions
		}
	} else {
		panic("parameter is null")
	}
}

func KafkaConsumerControllerInit(configValue *SettingKafka, topicName string) sarama.Consumer {
	kafkaControllerInit(configValue)
	tconsumer, err := sarama.NewConsumerFromClient(*m_kafkaClient)
	if err != nil {
		panic(err)
	}
	return tconsumer
}

func KafkaConsumerControllerGetPartitionConsumer(configValue *SettingKafka, setconsumer sarama.Consumer, topicName string, partitionNum int32) sarama.PartitionConsumer {
	partConsumer, err := setconsumer.ConsumePartition(topicName, partitionNum, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	return partConsumer
}

func KafkaGroupConsumerControllerInit(configValue *SettingKafka, groupid string) sarama.ConsumerGroup {
	kafkaControllerInit(configValue)
	groupConsumer, err := sarama.NewConsumerGroupFromClient(groupid, *m_kafkaClient)
	if err != nil {
		panic(err)
	}
	return groupConsumer
}
