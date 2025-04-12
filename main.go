package main

import (
	"fmt"
	"golangConcepts/models"
)

func main() {
	Ramana := models.Employee{
		Person:     models.Person{Name: "Ramana", Id: 1},
		Department: "IT",
	}

	Sandeep := models.Employee{
		Person:     models.Person{Name: "Sandeep", Id: 2},
		Department: "HR",
	}

	// Explicitly accessing the Draw method on the embedded Person field
	fmt.Println(Ramana.Draw())  // Output: Drawing, Ramana
	fmt.Println(Sandeep.Draw()) // Output: Drawing, Sandeep
}
