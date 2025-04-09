package main

import (
	"GOPROJECT/testing/models"
	"fmt"
)

func main() {

	Ramana := models.Person{Name: "Ramana", Id: 1}
	fmt.Println(Ramana.Draw())
	sandeep := models.Person{Name: "sandeep", Id: 2}
	fmt.Println(sandeep.Play())
	sandeep1 := models.Person{Name: "sandeep1", Id: 3}
	fmt.Println(sandeep1.Fight())
	fmt.Println(sandeep1.Draw())
	fmt.Println(sandeep1.Play())
	fmt.Println(sandeep1.Play())

}
