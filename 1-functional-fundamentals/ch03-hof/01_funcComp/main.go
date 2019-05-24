package main

import (
	"fmt"
	"strings"
)

type StrFunc func(string) string

func Compose(f StrFunc, g StrFunc) StrFunc {
	return func(s string) string {
		return g(f(s))
	}
}

func main() {
	var rec = func(name string) string {
		return fmt.Sprintf("Hey %s", name)
	}

	var emp = func(statement string) string {
		return fmt.Sprintf(strings.ToUpper(statement) + "!")
	}

	var greetFoG = Compose(rec, emp)
	fmt.Println(greetFoG("Gopher"))
}
