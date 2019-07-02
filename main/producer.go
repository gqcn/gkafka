package main

import (
    "fmt"
    "github.com/gogf/gkafka"
    "time"
)

// newKafkaClientProducer creates and returns a new kafka producer client.
func newKafkaClientProducer(topic string) *gkafka.Client {
    kafkaConfig               := gkafka.NewConfig()
    kafkaConfig.Servers        = "localhost:9092"
    kafkaConfig.AutoMarkOffset = false
    kafkaConfig.Topics         = topic
    return gkafka.NewClient(kafkaConfig)
}

func main()  {
    client := newKafkaClientProducer("test")
    defer client.Close()
    for {
        s := time.Now().String()
        fmt.Println("produce:", s)
        if err := client.SyncSend(&gkafka.Message{Value: []byte(s)}); err != nil {
            fmt.Println(err)
        }
        time.Sleep(time.Second)
    }
}
