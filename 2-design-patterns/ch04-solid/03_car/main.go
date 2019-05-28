package main

import (
	"fmt"

	. "github.com/YoungsoonLee/funtional-programming-in-go/2-design-patterns/ch04-solid/03_car/car"
)

func main() {
	accord := Car{"Honda", "Accord"}
	accord.PrintInfo()

	highlander := CarWithSpare{Car{"Toyota", "Highlander"}}
	highlander.PrintInfo()
	fmt.Printf("%v has %d tires", highlander.Car, highlander.Tries())

}
