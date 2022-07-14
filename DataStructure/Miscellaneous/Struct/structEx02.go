package main

import (
	"fmt"
	"strings"
)

type Part struct {
	Id   int
	Name string
}

func (part *Part) LowerCase() {
	part.Name = strings.ToLower(part.Name)
}

func (part *Part) UpperCase() {
	part.Name = strings.ToUpper(part.Name)
}

func (part Part) String() string {
	return fmt.Sprintf("«%d %q»", part.Id, part.Name)
}

func (part Part) HasPrefix(prefix string) bool {
	return strings.HasPrefix(part.Name, prefix)
}

func main() {
	part := Part{10, "wrench"}
	part.UpperCase()
	fmt.Println(part.String(), part.HasPrefix("W"))

	part.Id += 100
	fmt.Println(part.String(), part.HasPrefix("W"))

}
