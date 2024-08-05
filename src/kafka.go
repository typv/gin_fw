package main

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	_ "github.com/IBM/sarama"
	"log"
	"os"
	"os/signal"
	"syscall"
)

/* Register consumer, consumer group */
func CreateKafkaConsumer() sarama.Consumer {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// SASL/PLAIN
	//config.Net.SASL.Enable = true
	//config.Net.SASL.User = helpers.GetEnv("KAFKA_USERNAME", "root")
	//config.Net.SASL.Password = helpers.GetEnv("KAFKA_PASSWORD", "secret")
	//config.Net.SASL.Handshake = true
	//config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	//
	//tlsConfig := tls.Config{}
	//config.Net.TLS.Enable = true
	//config.Net.TLS.Config = &tlsConfig

	host := os.Getenv("KAFKA_HOST")
	port := os.Getenv("KAFKA_PORT")
	brokers := []string{fmt.Sprintf("%s:%s", host, port)}
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Printf("Failed to create Kafka consumer: %v", err)
	}

	return consumer
}

func StartConsuming(consumer sarama.Consumer, topic string, partition int32) error {
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		return fmt.Errorf("error creating partition consumer: %w", err)
	}
	defer partitionConsumer.Close()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Received message: key = %s, value = %s, partition = %d, offset = %d\n",
				string(msg.Key), string(msg.Value), msg.Partition, msg.Offset)
			fmt.Println(msg)
		case err := <-partitionConsumer.Errors():
			fmt.Printf("Error: %v\n", err)
		case <-sigchan:
			fmt.Println("Shutting down consumer...")
			return nil
		}
	}
}

func CreateKafkaConsumerGroup(groupID string) sarama.ConsumerGroup {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// SASL/PLAIN
	//config.Net.SASL.Enable = true
	//config.Net.SASL.User = helpers.GetEnv("KAFKA_USERNAME", "root")
	//config.Net.SASL.Password = helpers.GetEnv("KAFKA_PASSWORD", "secret")
	//config.Net.SASL.Handshake = true
	//config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	//
	//tlsConfig := tls.Config{}
	//config.Net.TLS.Enable = true
	//config.Net.TLS.Config = &tlsConfig

	host := os.Getenv("KAFKA_HOST")
	port := os.Getenv("KAFKA_PORT")
	brokers := []string{fmt.Sprintf("%s:%s", host, port)}
	consumer, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Printf("Failed to create Kafka consumer: %v", err)
	}

	return consumer
}

type ConsumerGroupHandler struct{}

func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Printf("Message claimed: value = %s, timestamp = %v, topic = %s\n", string(message.Value), message.Timestamp, message.Topic)
		session.MarkMessage(message, "")
	}
	return nil
}

func StartGroupConsuming(consumerGroup sarama.ConsumerGroup, topics []string, ctx context.Context) error {
	handler := ConsumerGroupHandler{}

	for {
		if err := consumerGroup.Consume(ctx, topics, handler); err != nil {
			return fmt.Errorf("error from consumer: %w", err)
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}

/* Register producer */
var Producer sarama.SyncProducer

func CreateKafkaProducer() sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// SASL/PLAIN
	//config.Net.SASL.Enable = true
	//config.Net.SASL.User = helpers.GetEnv("KAFKA_USERNAME", "root")
	//config.Net.SASL.Password = helpers.GetEnv("KAFKA_PASSWORD", "secret")
	//config.Net.SASL.Handshake = true
	//config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	//
	//tlsConfig := tls.Config{}
	//config.Net.TLS.Enable = true
	//config.Net.TLS.Config = &tlsConfig

	host := os.Getenv("KAFKA_HOST")
	port := os.Getenv("KAFKA_PORT")
	brokers := []string{fmt.Sprintf("%s:%s", host, port)}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Printf("Failed to create Kafka producer: %v", err)
	}

	Producer = producer

	return Producer
}
