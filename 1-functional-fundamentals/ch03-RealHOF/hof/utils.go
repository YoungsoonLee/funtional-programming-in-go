package hof

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	s "strings"
)

const DASHES = "-----------------------"

func PrintCars(title string, cars Collection) {
	log.Printf("\n%s\n%s\n", title, DASHES)
	for _, car := range cars {
		log.Printf("car: %v\n", car)
	}
}

func CsvToStruct(fileName string) Collection {
	/*
		pwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	*/
	pwd := "./"
	csvfile, err := os.Open(fmt.Sprintf("%s/%s", pwd, fileName))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	var cars Collection
	for _, car := range rawCSVdata {
		cars = append(cars, car[0])
	}

	return cars
}

func GetMake(sentence string) string {
	ret := sentence
	posSpace := s.Index(sentence, " ")

	//fmt.Println(ret, posSpace)

	if posSpace >= 0 {
		ret = sentence[:(posSpace)]
	}
	//fmt.Println(ret)
	return ret
}
