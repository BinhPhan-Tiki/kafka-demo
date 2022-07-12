package producer

import (
	"context"
	"fmt"
	"strconv"

	"github.com/segmentio/kafka-go"
)

const (
	topic          = "my-topic-1"
	broker1Address = "localhost:9092"
)

func Produce(ctx context.Context) {
	i := 0
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
	})

	// each kafka message has a key and value. The key is used
	// to decide which partition (and consequently, which broker)
	// the message gets published on
	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(strconv.Itoa(i)),
		// create an arbitrary message payload for the value
		Value: []byte("this is message" + strconv.Itoa(i)),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}

	// log a confirmation once the message is written
	fmt.Println("writes:", i)
	i++
	// sleep for a second

	// conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost", "topic_test", 1)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

	// 	conn.WriteMessages(kafka.Message{Value: []byte("hello kafka again")})
	// }
}
