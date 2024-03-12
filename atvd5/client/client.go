package main

import (
	"atvd5/client/utils"
	"atvd5/common"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
)

func runClient(boardNumber int, iterations int) {

	rand.Seed(time.Now().UTC().UnixNano())

	conn, err := amqp.Dial(common.GetConnectionSting())
	common.HandleError(err, "Unable to connect to messaging server")
	defer conn.Close()

	ch, err := conn.Channel()
	common.HandleError(err, "Unable to establish a communication channel with the messaging server")
	defer ch.Close()

	replyQueue, err := ch.QueueDeclare(
		common.ResponseQueue,
		false,
		false,
		true,
		false,
		nil,
	)

	common.HandleError(err, "Unable to connect to replyQueue")

	msgs, err := ch.Consume(
		replyQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	common.HandleError(err, "Failed to register the server with the broker")

	times := make([]float64, 0)

	board, err := utils.GetBoard(boardNumber)
	common.HandleError(err, "Unable to get a board")

	for i := 0; i < iterations; i++ {
		startTime := time.Now()

		// prepara mensagem
		msgRequest := common.Request{Board: board, ShouldTurnOff: false}
		msgRequestBytes, err := json.Marshal(msgRequest)
		common.HandleError(err, "Failed to serialize message")

		correlationID := common.RandomId(32)

		err = ch.Publish(
			"",
			common.RequestQueue,
			false,
			false,
			amqp.Publishing{
				ContentType:   "text/plain",
				CorrelationId: correlationID,
				ReplyTo:       replyQueue.Name,
				Body:          msgRequestBytes,
			},
		)

		common.HandleError(err, "Failed to publish message")

		m := <-msgs

		msgResponse := common.Reply{}
		err = json.Unmarshal(m.Body, &msgResponse)
		common.HandleError(err, "Error deserializing the response")

		// common.PrintBoard(msgResponse.R, 9)

		duration := time.Since(startTime)

		times = append(times, float64(duration.Microseconds()))
	}

	utils.WriteDataInFile(times, "rabbitmq", boardNumber)

	msgRequest := common.Request{Board: []int{}, ShouldTurnOff: true}
	msgRequestBytes, err := json.Marshal(msgRequest)
	common.HandleError(err, "Failed to serialize message")

	err = ch.Publish(
		"",
		common.RequestQueue,
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			ReplyTo:       replyQueue.Name,
			CorrelationId: "finish",
			Body:          msgRequestBytes,
		},
	)

	common.HandleError(err, "Failed to publish message")
}
