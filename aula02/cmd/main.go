package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"gihub.com/Adriano-Porto/go/internal/infra/database"
	"gihub.com/Adriano-Porto/go/internal/usecase"
	"gihub.com/Adriano-Porto/go/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	println("Running")
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	orderRepository := &database.OrderRepository{
		Db: db,
	}

	uc := usecase.CalculateFinalPrice{
		OrderRepository: orderRepository,
	}

	defer db.Close()

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel)

	rabbitmqWorker(msgRabbitmqChannel, uc)
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc usecase.CalculateFinalPrice) {
	fmt.Println("Starting rabbitmq")
	for msg := range msgChan {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			panic(err)
		}

		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		println("Mensagem output: ", output)
	}
}
