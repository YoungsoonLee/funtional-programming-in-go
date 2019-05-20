package main

type IntIterator interface {
	Next() (value string, ok bool)
}

const INVALID_INT_VAL = -1
const INVALID_STRING_VAL = ""

type Collection struct {
	index int
	List  []string
}

func (c *Collection) Next() (value string, ok bool) {
	c.index++
	if c.index >= len(c.List) {
		return INVALID_STRING_VAL, false
	}

	return c.List[c.index], true
}

func newSlice(s []string) *Collection {
	return &Collection{INVALID_INT_VAL, s}
}

func main() {
	var intCollection IntIterator
	intCollection = newSlice([]string{"CRV", "IS250", "Blzaer"})
	value, ok := intCollection.Next()
	for ok {
		println(value)
		value, ok = intCollection.Next()
	}
}
