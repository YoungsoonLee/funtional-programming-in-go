package hof

type Collection []string

type IndexedCar struct {
	Index int    `json:"index"`
	Car   string `json:"car"`
}

type FilterFunc func(string) bool

type Payload struct {
	IndexedCars []IndexedCar
}

type MapFunc func(string) string
