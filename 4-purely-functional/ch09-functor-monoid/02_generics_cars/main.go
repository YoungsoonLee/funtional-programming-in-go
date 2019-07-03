package main

import (
	"fmt"

	"github.com/YoungsoonLee/funtional-programming-in-go/4-purely-functional/ch09-functor-monoid/02_generics_cars/car"
)

func main() {
	var cars = car.CarSlice{
		car.Car{"Honda", "Accord", 3000},
		car.Car{"Lexus", "IS250", 40000},
		car.Car{"Toyota", "Highlander", 3500},
		car.Car{"Honda", "Accord ES", 3500},
	}

	fmt.Println("cars: ", cars)

	honda := func(c car.Car) bool {
		return c.Make == "Honda"
	}
	fmt.Println("filter cars by honda: ", cars.Where(honda))

	price := func(c car.Car) car.Dollars {
		return c.Price
	}

	fmt.Println("Honda prices: ", cars.Where(honda).SelectDollars(price))

	fmt.Println("Honda sum(prices): ", cars.Where(honda).SumDollars(price))

}
