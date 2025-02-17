package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

type order struct {
	id          string
	productName string
	amount      float32
	status      string
	orderedAt   time.Time
	customer
}

func (o order) getProductName() string {
	return o.productName
}

func (o *order) changeStatus(updatedStatus string) {
	o.status = updatedStatus
}

func main() {

	customerOne := customer{
		name:  "Anurag Band",
		phone: "1212121212121",
	}

	orderOne := order{
		id:          "1",
		productName: "Iphone 16",
		amount:      1499,
		status:      "UPI",
		orderedAt:   time.Now(),
		customer:    customerOne,
	}

	fmt.Println(orderOne)

	orderOne.changeStatus("cash_on_delivery")

	fmt.Println(orderOne)

}
