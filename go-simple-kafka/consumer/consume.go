package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func Must(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9099", "go-simple-kafka", 0)
	Must(err)

	conn.SetReadDeadline(time.Now().Add(time.Second * 3))
	// below commented code just read only the first message
	/*
		message, err1 := conn.ReadMessage(1e6)
		Must(err1)
		fmt.Println(string(message.Value))
	*/
	batch := conn.ReadBatch(1e3, 1e9) //1e3 = 1000 KB
	bytes := make([]byte, 1e3)

	for {
		_, err := batch.Read(bytes)
		Must(err)
		fmt.Println(string(bytes))
	}
}
