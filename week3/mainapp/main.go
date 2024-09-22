// main.go
package main

import (
	"fmt"

	"github.com/adilalimgozha/mymodule/mymath"
	"github.com/adilalimgozha/mymodule/utils"
)

func main() {
	sum := mymath.Add(10, 5)
	product := mymath.Multiply(10, 5)

	fmt.Println("Sum:", sum)
	fmt.Println("Product:", product)

	str := "Hello, Go!"
	upperStr := utils.ToUpper(str)
	reversedStr := utils.Reverse(str)

	fmt.Println("Uppercase:", upperStr)
	fmt.Println("Reversed:", reversedStr)
}
