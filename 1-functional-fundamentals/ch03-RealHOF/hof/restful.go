package hof

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func CarsIndexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response, err := getAllCarsJson()
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(response))
}

func CarHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	carIndex, err := strconv.Atoi(p[0].Value)
	if err != nil {
		log.Fatal("CarHandler unable to find car (%v) by index\n", p[0].Value)
	}
	response, err := getThisCarJson(carIndex)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(response))
}

func getAllCarsJson() ([]byte, error) {
	return json.MarshalIndent(Payload{CarsDB}, "", "  ")
}

func getThisCarJson(carIndex int) ([]byte, error) {
	return json.MarshalIndent(CarsDB[carIndex], "", "  ")
}
