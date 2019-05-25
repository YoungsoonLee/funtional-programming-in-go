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

type ReducerFunc func(string, Collection) Collection
type ReducerFunc2 func(string, CarCollection) CarCollection

type CarCollection []CarType
type CarType struct {
	Make  string `json:"make"`
	Model string `json:"model"`
}
