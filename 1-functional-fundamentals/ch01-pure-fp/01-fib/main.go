package main

import "github.com/YoungsoonLee/funtional-programming-in-go/1-functional-fundamentals/ch01-pure-fp/01-fib/fibonacci"

func main() {
	println(fibonacci.FibSimple(4))
	println(fibonacci.FibMemoized(5))
	println(fibonacci.FibChanneled(6))
}
