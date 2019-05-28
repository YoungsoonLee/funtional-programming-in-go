package car

import "fmt"

type Car struct {
	Make  string
	Model string
}

func (c Car) Tries() int { return 4 }
func (c Car) PrintInfo() {
	fmt.Printf("%v has %d tries.\n", c, c.Tries())
}

type CarWithSpare struct {
	Car
}

func (o CarWithSpare) Tries() int { return 5 }
