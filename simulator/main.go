package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	kafka2 "github.com/xXHachimanXx/Imersao-FullCycle-2.0/application/kafka"
	kafka "github.com/xXHachimanXx/Imersao-FullCycle-2.0/infra/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		go kafka2.Produce(msg)
		fmt.Println(string(msg.Value))
	}
	// route := route.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }

	// route.LoadPositions()
	// stringJson, _ := route.ExportJsonPositions()
	// fmt.Println(stringJson[1])
}
