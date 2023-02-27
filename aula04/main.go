package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position int
	Person
}

func (e Employee) PrintEmployee() {
	fmt.Println(
		"Name", e.Name,
		"Data de Aniversário", e.DateOfBirth,
		"ID_Person", e.Person.ID,
		"ID_Employee", e.ID,
		"Position", e.Position,
	)
}

func main() {
	person := Person{
		ID:          1,
		Name:        "João",
		DateOfBirth: "01/01/2020",
	}

	employee := Employee{
		ID:       9,
		Position: 1,
		Person:   person,
	}

	employee.PrintEmployee()
}
