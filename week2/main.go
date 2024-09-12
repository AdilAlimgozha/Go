package main

import "fmt"

type Student struct {
	id     int
	Name   string
	Course int
}

type Worker struct {
	id   int
	Name string
}

type Teacher struct {
	Base   Worker
	Degree string
}

type Dean struct {
	Worker
	Age int
}

type Job interface {
	Teach()
	Reasearch()
}

func (t Teacher) Teach() string {
	return "I'm teaching"
}

func (d Dean) Teach() string {
	return "I'm Dean, I'm teaching"
}

func (d Dean) Reasearch() string {
	return "I do researches"
}

func (s *Student) GetID() int {
	return s.id
}

func main() {
	student := &Student{
		id:     123,
		Name:   "Adil",
		Course: 4}

	var teacher Teacher
	teacher.Base.id = 12345
	teacher.Base.Name = "Adil"
	teacher.Degree = "phd"

	var dean Dean
	dean.id = 1234
	dean.Name = "qwerty"
	dean.Age = 40

	fmt.Println(student.id, student.Name, student.Course)
	fmt.Println(teacher.Base.id, teacher.Base.Name, teacher.Degree)
	fmt.Println(dean.id, dean.Name, dean.Age)

	fmt.Println(teacher.Teach())
	fmt.Println(dean.Teach())
	fmt.Println(dean.Reasearch())

	fmt.Println(student.GetID())
}
