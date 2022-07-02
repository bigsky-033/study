package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/Shopify/sarama"
)

var (
	brokers = flag.String("brokers", os.Getenv("KAFKA_PEERS"), "The Kafka brokers to connect to, as a comma separated list")
)

func main() {
	flag.Parse()

	if *brokers == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

	brokerList := strings.Split(*brokers, ",")
	log.Printf("Kafka brokers: %s\n", strings.Join(brokerList, ","))

	producer, err := newProducer(brokerList)
	if err != nil {
		log.Fatalln("Failed to create producer:", err)
	}

	topic := "purchases"
	users := [...]string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
	items := [...]string{"book", "alarm clock", "t-shirts", "gift card", "batteries"}

	for n := 0; n < 10; n++ {
		key := users[rand.Intn(len(users))]
		data := items[rand.Intn(len(items))]

		if _, _, err := producer.SendMessage(&sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(key),
			Value: sarama.StringEncoder(data),
		}); err != nil {
			fmt.Printf("Failed to deliver message: %v\n", err)
		} else {
			fmt.Printf("Produced event to topic %s: key = %-10s value = %-10s\n", topic, key, data)
		}
	}

	producer.Close()
}

func newProducer(brokerList []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true

	return sarama.NewSyncProducer(brokerList, config)
}
