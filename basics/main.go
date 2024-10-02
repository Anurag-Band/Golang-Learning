package main

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

}
