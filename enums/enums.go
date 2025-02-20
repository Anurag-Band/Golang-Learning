package main

import "fmt"

// ? for int type
// type OrderStatus int

// const (
// 	Received OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )

// ? for string type
type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delevered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order Status: ", status)
}

func main() {
	changeOrderStatus(Prepared)
}
