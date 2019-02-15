package main

import (
    "fmt"
    "github.com/gogf/gkafka"
)

func main()  {
    config        := gkafka.NewConfig()
    config.Servers = "localhost:9092"

    client := gkafka.NewClient(config)
    defer client.Close()

    fmt.Println(client.Topics())
}
