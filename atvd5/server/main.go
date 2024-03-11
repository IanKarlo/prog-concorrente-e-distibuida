package main

import (
	"atvd5/common"
	"atvd5/server/impl"
	"encoding/json"
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial(common.GetConnectionSting())
	common.HandleError(err, "Unable to connect to broker")

	defer conn.Close()

	ch, err := conn.Channel()
	common.HandleError(err, "Unable to establish a communication channel with the broker")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		common.RequestQueue,
		false,
		false,
		false,
		false,
		nil)
	common.HandleError(err, "Unable to create queue in broker")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	common.HandleError(err, "Failed to register the consumer with the broker")

	for data := range msgs {
		msg := common.Request{}
		err := json.Unmarshal(data.Body, &msg)
		common.HandleError(err, "Failed to deserialize message")

		solver := impl.SudokuSolver{}

		if msg.ShouldTurnOff {
			fmt.Println("Shutting Down Server")
			os.Exit(0)
		}

		r := solver.Run(msg)

		replyMsg := common.Reply{R: r}
		replyMsgBytes, err := json.Marshal(replyMsg)
		common.HandleError(err, "Failed to serialize message")

		err = ch.Publish(
			"",
			data.ReplyTo,
			false,
			false,
			amqp.Publishing{
				ContentType:   "text/plain",
				CorrelationId: data.CorrelationId,
				Body:          replyMsgBytes,
			},
		)
		common.HandleError(err, "Failed to send the message to the broker")
	}
}
