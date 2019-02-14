// Copyright 2018 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gkafka_test

import (
    "fmt"
    "github.com/gogf/gf/g/os/gtime"
    "github.com/gogf/gkafka"
    "testing"
    "time"
)

// 创建kafka生产客户端
func newKafkaClientProducer(topic string) *gkafka.Client {
    kafkaConfig               := gkafka.NewConfig()
    kafkaConfig.Servers        = "localhost:9092"
    kafkaConfig.AutoMarkOffset = false
    kafkaConfig.Topics         = topic
    return gkafka.NewClient(kafkaConfig)
}

// 创建kafka消费客户端
func newKafkaClientConsumer(topic, group string) *gkafka.Client {
    kafkaConfig               := gkafka.NewConfig()
    kafkaConfig.Servers        = "localhost:9092"
    kafkaConfig.AutoMarkOffset = false
    kafkaConfig.Topics         = topic
    kafkaConfig.GroupId        = group
    return gkafka.NewClient(kafkaConfig)
}

// 生产者
func Test_Producer(t *testing.T) {
    client := newKafkaClientProducer("test")
    defer client.Close()
    for {
        s := gtime.Now().String()
        fmt.Println("produce:", s)
        if err := client.SyncSend(&gkafka.Message{Value: []byte(s)}); err != nil {
            fmt.Println(err)
        }
        time.Sleep(time.Second)
    }
}

// 消费者
func Test_Consumer(t *testing.T) {
    group  := "test-group"
    topic  := "test"
    client := newKafkaClientConsumer(topic, group)
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

// 获取所有topics
func Test_Topics(t *testing.T) {
    config        := gkafka.NewConfig()
    config.Servers = "localhost:9092"

    client := gkafka.NewClient(config)
    defer client.Close()

    fmt.Println(client.Topics())
}
