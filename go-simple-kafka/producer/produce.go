package main

import (
	"context"
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

	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	conn.WriteMessages(kafka.Message{Value: []byte("I`m connected to Kafka through an external protocol:9099")})
}
