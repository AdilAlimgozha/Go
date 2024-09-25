package main

import (
	"fmt"

	"github.com/adilalimgozha/lolModule/mymath"
	"github.com/adilalimgozha/lolModule/utils"
)

func main() {
	// Используем функцию из пакета mymath
	result := mymath.Add(5, 3)
	fmt.Println("Result:", result)

	// Используем функцию из пакета utils
	utils.PrintMessage("Hello from utils package!")
}
