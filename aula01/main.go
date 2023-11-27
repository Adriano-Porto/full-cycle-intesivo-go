package main

import (
	"database/sql"
	"fmt"

	"gihub.com/Adriano-Porto/go/internal/infra/database"
	"gihub.com/Adriano-Porto/go/internal/usecase"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	orderRepository := &database.OrderRepository{
		Db: db,
	}

	order := usecase.CalculateFinalPrice{
		OrderRepository: orderRepository,
	}

	input := usecase.OrderInput{
		ID:    "1",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := order.Execute(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)

}
