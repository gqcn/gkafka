# gkafka 

[![Go Doc](https://godoc.org/github.com/gogf/gkafka?status.svg)](https://godoc.org/github.com/gogf/gkafka)
[![Documents](https://img.shields.io/badge/docs-100%25-green.svg)](https://goframe.org) 
[![License](https://img.shields.io/github/license/gogf/gkafka.svg?style=flat)](https://github.com/gogf/gkafka)
[![Language](https://img.shields.io/badge/language-go-blue.svg)](https://github.com/gogf/gkafka)
[![Release](https://img.shields.io/github/release/gogf/gkafka.svg?style=flat)](https://github.com/gogf/gkafka/releases)

Package gkafka provides producer and consumer client for kafka server.

# Installation
```
go get -u github.com/gogf/gkafka
```
or use `go.mod`
```
require github.com/gogf/gkafka latest
```

# Documentation

* https://godoc.org/github.com/gogf/gkafka

# Quick Start

## Producer
```go
package main

import (
    "fmt"
    "github.com/gogf/gkafka"
    "time"
)

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
```
## Consumer
```go
package main

import (
    "fmt"
    "github.com/gogf/gkafka"
)

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
```

# License

`gkafka` is licensed under the [MIT License](LICENSE), 100% free and open-source, forever.

