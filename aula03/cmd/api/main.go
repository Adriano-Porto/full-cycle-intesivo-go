package main

import (
	"gihub.com/Adriano-Porto/go/internal/entity"
	"github.com/labstack/echo/v4"
)

func main() {
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/order", OrderHandler)
	// http.ListenAndServe(":8888", r)
	e := echo.New()
	e.GET("/order", OrderHandler)
	e.Logger.Fatal(e.Start(":8888"))
}

func OrderHandler(c echo.Context) error {
	order, _ := entity.NewOrder("1", 10, 1)

	err := order.CalculateFinalPrice()

	if err != nil {
		return c.String(500, err.Error())
	}
	return c.JSON(200, order)
}

// func OrderHandler(w http.ResponseWriter, r *http.Request) {

// 	order, err := entity.NewOrder("1", 10, 1)

// 	if err != nil {
// 		w.WriteHeader(402)
// 	}
// 	err = order.CalculateFinalPrice()

// 	if err != nil {
// 		w.WriteHeader(500)
// 	}

// 	json.NewEncoder(w).Encode(order)
// }
