package main

import "fmt"

func main(){
	//1
	fmt.Println("Hello, World")

	//2
	var a int = 50
	var b string = "Age"
	var c float64 = 1.5
	var d = true
	e := "Name"
	fmt.Println(a, b, c, d, e)

	//3
	//if
	var num int
	fmt.Scan(&num)
	if num > 0 {
		fmt.Println("Positive")
	}else if num < 0{
		fmt.Println("Negative")
	}else{
		fmt.Println("Zero")
	}

	//for
	var sum int = 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//switch
	var x int
	fmt.Scan(&x)
	switch x {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thirsday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("There is no such day")
	}

	//functions
	fmt.Println(add(4, 5))

	swap := func(a string, b string) (string, string){
		return b, a
	}

	fmt.Println(swap("good", "job"))

	fmt.Println(quot_remain(11, 2))

}

func add(a int, b int) int{
	return a + b
}

func quot_remain(a int, b int) (int, int){
	return a/b, a%b
}