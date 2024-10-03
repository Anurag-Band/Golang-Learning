package main

import "fmt"

// const outAge = 12

// var newName string = "Param"

//? := will not work outside main()

func main() {
	// fmt.Println("Hello World !")

	//? Variables

	// var name string = "Anurag"
	// fmt.Println(name)

	// age := 22
	// fmt.Println(age)

	// var location string
	// location = "Karanja"
	// fmt.Println(location)

	//? Constants

	const friend = "devang"

	//? for multiple constants grouping

	// const (
	// 	port = 8080
	// 	server = "go"
	// )

	// fmt.Println(port, server)

	//? Loops - for -> only construct in go for looping (while also implemented using for)

	//? while loop
	// i := 0
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//? infinite loop
	// for {
	// 	fmt.Println("fuck")
	// }

	//? classic for loop
	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	//? continue keyword
	// 	for i := 0; i <= 5; i++ {

	// 		if i == 4 {
	// 			continue
	// 		}
	// 		fmt.Println(i)
	// 	}

	//? break keyword
	// for i := 0; i <= 5; i++ {

	// 	if i == 4 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	//? range keyword
	// for i := range 10 {
	// 	fmt.Println(i)
	// }

	//? if else if conditions

	// age := 13

	// if age >= 18 {
	// 	fmt.Println("you are an adult")
	// } else if age >= 12 {
	// 	fmt.Println("you are a teenager")
	// } else {
	// 	fmt.Println("you are a kid")
	// }

	//? we can declear a variable inside if construct

	// if age := 13; age >= 18 {
	// 	fmt.Println("you are an adult")
	// } else if age >= 12 {
	// 	fmt.Println("you are a teenager")
	// } else {
	// 	fmt.Println("you are a kid")
	// }

	//? simple switch case

	// 	num := 4

	// 	switch {
	// 	case num == 1:
	// 		fmt.Println("One")
	// 	case num == 2:
	// 		fmt.Println("Two")
	// 	case num == 3:
	// 		fmt.Println("Three")
	// 	default:
	// 		fmt.Println("Other Number")
	// 	}

	//? multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("today is Holiday")
	// default:
	// 	fmt.Println("today is Workday")
	// }

	//? type switch

	// whoAmI := func(i interface{}) {
	// 	switch i.(type) {
	// 	case int:
	// 		fmt.Println("it's an Integer")
	// 	case bool:
	// 		fmt.Println("it's a Boolean")
	// 	case float32, float64:
	// 		fmt.Println("it's a floating num")
	// 	default:
	// 		fmt.Println("it's of other category")
	// 	}
	// }
	// whoAmI("ANurag")

	//? arrays 
	// - fixed length - use when data is predictable
	// - Memory Optimization
	// - Constant time access
	

	// var name [4]string
	// name[0] = "Anurag" //? others will be empty strings
	// fmt.Println(name)

	// number := [5]int{10} //? others will be Zero
	// fmt.Println(number)

	// var bool [5]bool
	// bool[2] = true //? others will be false
	// fmt.Println(bool)

	//? 3d arrays
	twoDarray := [3][3]int{{12, 34, 45}, {56, 76, 78}, {11, 22, 33}}

	fmt.Println(twoDarray)

}
