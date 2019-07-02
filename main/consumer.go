package main

import (
    "fmt"
    "github.com/gogf/gkafka"
)

// newKafkaClientConsumer creates and returns a new kafka producer client.
func newKafkaClientConsumer(topic, group string) *gkafka.Client {
    kafkaConfig               := gkafka.NewConfig()
    kafkaConfig.Servers        = "localhost:9092"
    kafkaConfig.AutoMarkOffset = false
    kafkaConfig.Topics         = topic
    kafkaConfig.GroupId        = group
    return gkafka.NewClient(kafkaConfig)
}

func main()  {
    group  := "test-group"
    topic  := "test"
    client := newKafkaClientConsumer(topic, group)
    defer client.Close()

	// Mark the offset from reading.
    //client.MarkOffset(topic, 0, 6)
    for {
        if msg, err := client.Receive(); err != nil {
            fmt.Println(err)
            break
        } else {
            fmt.Println("consume:", msg.Partition, msg.Offset, string(msg.Value))
            msg.MarkOffset()
        }
    }
}
