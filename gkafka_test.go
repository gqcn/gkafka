// Copyright 2018 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gkafka_test

import (
    "fmt"
    "github.com/gogf/gkafka"
    "time"
)

// Create producer.
//
// 生产者
func Example_producer() {
    topic                     := "test"
    kafkaConfig               := gkafka.NewConfig()
    kafkaConfig.Servers        = "localhost:9092"
    kafkaConfig.AutoMarkOffset = false
    kafkaConfig.Topics         = topic

    client := gkafka.NewClient(kafkaConfig)
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

// Create consumer.
//
// 消费者
func Example_consumer() {
    group                     := "test-group"
    topic                     := "test"
    kafkaConfig               := gkafka.NewConfig()
    kafkaConfig.Servers        = "localhost:9092"
    kafkaConfig.AutoMarkOffset = false
    kafkaConfig.Topics         = topic
    kafkaConfig.GroupId        = group

    client := gkafka.NewClient(kafkaConfig)
    defer client.Close()

    // 标记开始读取的offset位置
    client.MarkOffset(topic, 0, 6)
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

// Fetch all topics from server.
//
// 获取所有topics
func Example_topics() {
    config        := gkafka.NewConfig()
    config.Servers = "localhost:9092"

    client := gkafka.NewClient(config)
    defer client.Close()

    fmt.Println(client.Topics())
}
