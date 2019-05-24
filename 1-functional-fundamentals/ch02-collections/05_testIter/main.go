package main

import (
	"reflect"
	"testing"

	. "github.com/yanatan16/itertools"
)

func TestIterEq(t *testing.T, it1, it2 Iter) {
	t.Log("start")
	for el1 := range it1 {
		if el2, ok := <-it2; ok {
			t.Error("it2 shorter than it1!", el1)
			return
		} else if !reflect.DeepEqual(el1, el2) {
			t.Error("Elements are not equal", el1, el2)
			return
		} else {
			t.Log(el1, el2)
		}
	}
	if el2, ok := <-it2; ok {
		t.Error("it1 shorter than it2!", el2)
		return
	}
	t.Log("stop")
}

func main() {

}
