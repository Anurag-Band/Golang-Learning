package main

import (
	"fmt"
	"time"
)

// ? 1st way
// type customer struct {
// 	name  string
// 	phone string
// }

// type order struct {
// 	id          string
// 	productName string
// 	amount      float32
// 	status      string
// 	orderedAt   time.Time
// 	customer
// }

// ? receiver type
// func (o order) getProductName() string {
// 	return o.productName
// }

// func (o *order) changeStatus(updatedStatus string) {
// 	o.status = updatedStatus
// }

// ? constructor with go (a custom wrapper to use classes like other programming languages)

type order struct {
	id          string
	productName string
	amount      float32
	status      string
	orderedAt   time.Time
}

func newOrder(id string, productName string, amount float32, status string) *order {
	// intitial setup goes here - like packages, etc...
	orderConstructor := order{
		id:          id,
		productName: productName,
		amount:      amount,
		status:      status,
	}
	return &orderConstructor
}

func main() {

	// ? 1st way
	// customerOne := customer{
	// 	name:  "Anurag Band",
	// 	phone: "1212121212121",
	// }

	// orderOne := order{
	// 	id:          "1",
	// 	productName: "Iphone 16",
	// 	amount:      1499,
	// 	status:      "UPI",
	// 	orderedAt:   time.Now(),
	// 	customer:    customerOne,
	// }

	// fmt.Println(orderOne)

	// orderOne.changeStatus("cash_on_delivery")

	// fmt.Println(orderOne)

	// ? by using constructor wrapper

	orderOne := newOrder("2", "Samsung S25 ultra", 1899.99, "dispatched")

	fmt.Println(orderOne)

}
