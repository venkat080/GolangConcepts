package models

import (
	"fmt"
)

type Person struct {
	Name string
	Id   int
}

func (p *Person) Draw() string {
	return fmt.Sprintf("Drawing, %s", p.Name)
}
func (p *Person) Play() string {
	return fmt.Sprintf("Playing, %s", p.Name)
}
func (p *Person) Fight() string {
	return fmt.Sprintf("Fighting, %s", p.Name)
}
