package logrus_kafka

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

type KafkaHook struct {
	// Id of the hook
	id string

	// Log levels allowed
	levels []logrus.Level

	// Log entry formatter
	formatter logrus.Formatter

	// sarama.AsyncProducer
	producer sarama.SyncProducer

	topic string

	errChannel chan error
}

// Create a new KafkaHook.
func NewKafkaSyncHook(id string,
	levels []logrus.Level,
	formatter logrus.Formatter,
	brokers []string,
	topic string) (*KafkaHook, error) {

	// Kafka async configuration
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	kafkaConfig.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	kafkaConfig.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

	producer, err := sarama.NewSyncProducer(brokers, nil)

	if err != nil {
		return nil, err
	}

	// We will just log to STDOUT if we're not able to produce messages.
	// Note: messages will only be returned here after all retry attempts are exhausted.

	hook := new(KafkaHook)
	hook.id = id
	hook.levels = levels
	hook.formatter = formatter
	hook.producer = producer
	hook.topic = topic
	hook.errChannel = make(chan error, 1)

	go func() {
		for err := range hook.Errors() {
			log.Printf("Failed to send log entry to kafka: %v\n", err)
		}
	}()

	return hook, nil
}

func (hook *KafkaHook) Id() string {
	return hook.id
}

func (hook *KafkaHook) Levels() []logrus.Level {
	return hook.levels
}

func (hook *KafkaHook) Errors() chan error {
	return hook.errChannel
}

func (hook *KafkaHook) Fire(entry *logrus.Entry) error {
	// Check time for partition key
	var partitionKey sarama.ByteEncoder

	// Get field time
	t, _ := entry.Data["time"].(time.Time)

	// Convert it to bytes
	b, err := t.MarshalBinary()

	if err != nil {
		return err
	}

	partitionKey = sarama.ByteEncoder(b)

	// Format before writing
	b, err = hook.formatter.Format(entry)

	if err != nil {
		return err
	}

	value := sarama.ByteEncoder(b)

	msg := &sarama.ProducerMessage{
		Key:   partitionKey,
		Topic: hook.topic,
		Value: value,
	}

	partition, offset, err := hook.producer.SendMessage(msg)

	if err != nil {
		log.Printf("FAILED to send message: %s%d%d\n", err, partition, offset)
		hook.Errors() <- err
	}
	return nil
}
