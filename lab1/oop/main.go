package main

import (
	"encoding/json"
	"fmt"
	"math"
)

// 1
type Person struct {
	Name string
	Age  int
}

func (p Person) Greeting() string {
	return "Hello " + p.Name
}

// 2
type Employee struct {
	Name string
	ID   int
}

type Manager struct {
	Employee
	Department string
}

func (e Employee) Work() {
	fmt.Println(e.Name, e.ID)
}

// 3
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

//4

type Product struct {
	Name     string  `json:"product_name"`
	Price    float64 `json:"product_price"`
	Quantity int     `json:"product_quantity"`
}

func convertToJson(prod Product) string {
	b, err := json.Marshal(prod)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func convertToStruct(jso string) {
	byt := []byte(jso)
	var dat Product

	err := json.Unmarshal(byt, &dat)
	if err != nil {
		panic(err)
	}
	fmt.Println(dat)
}

func main() {
	//1

	var person Person

	person.Name = "Adil"
	person.Age = 21

	fmt.Println(person.Greeting())

	//2
	var manager Manager

	manager.Name = "Adil"
	manager.ID = 123456789
	manager.Department = "KBTU"

	manager.Work()

	//3
	var circle Circle
	var rectangle Rectangle

	circle.Radius = 2
	rectangle.Length = 3
	rectangle.Width = 4

	PrintArea(circle)
	PrintArea(rectangle)

	//4

	var product Product

	product.Name = "Potato"
	product.Price = 12
	product.Quantity = 20

	json_ := convertToJson(product)
	fmt.Println(json_)

	convertToStruct(json_)
}
