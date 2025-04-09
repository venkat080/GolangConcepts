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

	// Sandeep := models.Employee{
	// 	models.Person{Name: "sandeep", Id: 2},
	// 	"IT",
	// }

	fmt.Println(Ramana.Draw())
	// fmt.Println(Sandeep.Draw())

}
