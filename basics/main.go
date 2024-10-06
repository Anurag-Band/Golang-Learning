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

	// fmt.Println(len(name))

	// number := [5]int{10} //? others will be Zero
	// fmt.Println(number)

	// var bool [5]bool
	// bool[2] = true //? others will be false
	// fmt.Println(bool)

	//? 3d arrays
	// twoDarray := [3][3]int{{12, 34, 45}, {56, 76, 78}, {11, 22, 33}}

	// fmt.Println(twoDarray)

	//? slice -> dynamic
	// - most used construct in go

	//  uninitialized slice is nil by default

	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// to add values while initializing slices use make()

	// var numbers = make([]int, 3, 5)
	//? zero is added initially
	//? - it will double size after given size is passed
	//? - third parameter is to set maximum size initially
	//? - but this will also increase dynamically as we will as elements

	// numbers = append(numbers, 12)
	// numbers = append(numbers, 12)
	// numbers = append(numbers, 12)
	// numbers = append(numbers, 12)
	// numbers = append(numbers, 12)

	// fmt.Println(numbers)

	//? capacity -> maximum numbers of elements can fit (this is minimum, but will change dynamically)
	// fmt.Println(cap(numbers))

	// //? Array vs Slices ->>
	// Initialize an array of 5 integers
	// arr := [5]int{1, 2, 3, 4, 5}
	// arr[6] = 12 //? this will give error cause its range bound

	// Initialize a slice of integers
	// slice := []int{1, 2, 3, 4, 5}
	// slice[6] = 12 //? this will give error cause its range bound

	// fmt.Println("Array:", arr)
	// fmt.Println("Slice:", slice)

	// Append a new value to the slice
	// slice = append(slice, 12) //? This adds '12' to the end of the slice
	//? - append() method will be dynamic slice

	//? copy function

	// var num1 = make([]int, 2, 5)
	// num1 = append(num1, 2, 4)
	// var num2 = make([]int, len(num1))

	// copy(num2, num1)
	// fmt.Println(num1, num2)

	//? slice operator

	// var nums = []int{1,2,3,4,5}

	// fmt.Println(nums[0:2])
	// fmt.Println(nums[:5]) //? start with 0 till end number
	// fmt.Println(nums[2:]) //? start with 2 till last element

	//? slices methods

	//? Equal() - for comparing two slices
	// var num1 = []int{1,2,3}
	// var num2 = []int{3,2,1}

	// fmt.Println(slices.Equal(num1,num2))

	//? 2d slices

	// twoDslice := [][]int{{12, 34}, {56, 76}}

	// fmt.Println(twoDslice)

	//? maps
	// - key value pair - //? like object in javascript

	// m := make(map[string]string)

	// m["name"] = "Anurag"
	// m["position"] = "Software Developer"
	// m["delete"] = "Delete ME !"
	// fmt.Println(m)

	// deleting key:value pair
	// delete(m, "delete")
	// fmt.Println(m)

	// getting length
	// fmt.Println(len(m))

	// clearing map
	// clear(m)
	// fmt.Println(m)

	//? using object construct with interface{} for accepting any type of key:value pairs
	// myMap := map[string]interface{}{
	// 	"product": "Iphone 16 pro",
	// 	"price" : "$1299",
	// }
	// fmt.Println(myMap)

	// //? for compairng

	// myMap1 := map[string]interface{}{
	// 	"product": "Iphone 16 pro",
	// 	"price" : 99000,
	// }

	// myMap2 := map[string]interface{}{
	// 	"product": "Iphone 16 pro",
	// 	"price" : 79000,
	// }

	// fmt.Println(maps.Equal(myMap1,myMap2))

	//? for conditions

	// myMap := map[string]interface{}{
	// 	"product": "Iphone 16 pro",
	// 	"price":   99000,
	// }

	// _, isOk := myMap["price"]

	// if isOk {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	// Use type assertion to get the price
	// if price, ok := myMap["price"].(int); ok {
	// 	if price < 100000 {
	// 		fmt.Println("it's in our Budget!")
	// 	} else {
	// 		fmt.Println("it's out of our Budget!")
	// 	}
	// } else {
	// 	fmt.Println("Price not found or not an integer.")
	// }

	//? range

	// mySlice := []int{5,6,7,87,98,}
	// for i, v := range mySlice {
	// 	fmt.Println(i, v)
	// }

	// myMaps := map[string]string{
	// 	"fname" : "Anurag",
	// 	"lname" : "Band",
	// }

	// for k, v := range myMaps {
	// 	fmt.Println(k,"is",v)
	// }

	// unicode code point rune
	// starting byte of rune
	// 300 -> 1 byte, 2 byte
	// for i, c := range "Anurag" {
	// 	fmt.Println(i, c)
	// }

}
