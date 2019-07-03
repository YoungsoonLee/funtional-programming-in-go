package main

import (
	"fmt"
)

func main() {
	// 순수 함수형 프로그래밍은 원하는 바가 무엇인지만 함자에게 알려준다.
	ints := []int{1, 2, 3}
	res := []int{}

	// imperative
	for _, v := range ints {
		res = append(res, v+1)
	}

	fmt.Println("imperative: ", res)

	// function
	// add1 := func(i int) int { return i + 1 }
	// fpInts := Functor(ints).Map(add1)

}
