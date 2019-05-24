package hof

import (
	"fmt"
	"regexp"
	s "strings"
)

var CarsDB = initCarsDB()

func initCarsDB() []IndexedCar {
	var indexedCars []IndexedCar

	for i, car := range LoadCars() {
		indexedCars = append(indexedCars, IndexedCar{i, car})
	}

	lenCars := len(indexedCars)
	for i, car := range LoadMoreCars() {
		indexedCars = append(indexedCars, IndexedCar{lenCars + i, car})
	}

	return indexedCars
}

func LoadCars() Collection {
	return CsvToStruct("cars.csv")
}

func LoadMoreCars() Collection {
	return CsvToStruct("more_cars.csv")
}

func (cars Collection) AddCars(carsToAdd Collection) Collection {
	return append(cars, carsToAdd...)
}

func (cars Collection) Filter(fn FilterFunc) Collection {
	filteredCars := make(Collection, 0)
	for _, car := range cars {
		if fn(car) {
			filteredCars = append(filteredCars, car)
		}
	}
	return filteredCars
}

func ByMake(make string) FilterFunc {
	return func(car string) bool {
		return s.Contains(car, make)
	}
}

func ByHasNumber() FilterFunc {
	return func(car string) bool {
		match, _ := regexp.MatchString(".+[0-9].*", car)
		return match
	}
}

func ByForegin() FilterFunc {
	return func(car string) bool {
		return !isDomestic(car)
	}
}

func ByDomestic() FilterFunc {
	return func(car string) bool {
		return isDomestic(car)
	}
}

func isDomestic(car string) bool {
	return s.Contains(car, "Ford") || s.Contains(car, "GM") || s.Contains(car, "Chrysler")
}

func (cars Collection) Map(fn MapFunc) Collection {
	mappedCars := make(Collection, 0, len(cars))
	for _, car := range cars {
		mappedCars = append(mappedCars, fn(car))
	}
	return mappedCars
}

func Upgrade() MapFunc {
	return func(car string) string {
		fmt.Println("tt: ", UpgradeLabel(car))
		return fmt.Sprintf("%s %s", car, UpgradeLabel(car))
	}
}

func UpgradeLabel(car string) string {
	return map[string]string{
		"Honda":  "LX",
		"Lexus":  "LS",
		"Toyota": "EV",
		"Ford":   "XL",
		"GM":     "X",
	}[GetMake(car)]
}
