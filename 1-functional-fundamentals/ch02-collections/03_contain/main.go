package main

type Car struct {
	Make  string
	Model string
}

type Cars []*Car

func (cars Cars) Contains(modelName string) bool {
	for _, a := range cars {
		if a.Model == modelName {
			return true
		}
	}
	return false
}

type Object interface{}
type Collection []Object

func (list Collection) Contains(e string) bool {
	for _, t := range list {
		if t == e {
			return true
		}
	}
	return false
}

func main() {
	crv := &Car{"Honda", "CRV"}
	is250 := &Car{"Lexus", "IS250"}
	highlander := &Car{"Toyota", "Highlander"}
	cars := Cars{crv, is250, highlander}

	if cars.Contains("Highlander") {
		println("Found Highlander")
	}
	if !cars.Contains("Hummer") {
		println("Dia not find a Hummer")
	}
}
