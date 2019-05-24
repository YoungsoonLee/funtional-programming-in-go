package main

import (
	"log"
	"net/http"
	"os"

	"github.com/YoungsoonLee/funtional-programming-in-go/1-functional-fundamentals/ch03-RealHOF/hof"
	"github.com/julienschmidt/httprouter"
)

func init() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
}

func main() {
	if os.Getenv("RUN_HTTP_SERVVER") == "TRUE" {
		router := httprouter.New()
		router.GET("/cars", hof.CarsIndexHandler)
		router.GET("/cars/:id", hof.CarHandler)
		log.Println("Listening on port 8000")
		log.Fatal(http.ListenAndServe(":8000", router))
	} else {
		cars := hof.LoadCars()
		log.Printf("cars: %+v", cars)

		hof.PrintCars("Numeric", cars.Filter(hof.ByHasNumber()))

		hof.PrintCars("Foreign, Numeric, Toyota", cars.Filter(hof.ByForegin()).Filter(hof.ByHasNumber()).Filter(hof.ByMake("Toyota")))

		moreCars := hof.LoadMoreCars()
		hof.PrintCars("More Cars, Domestic, Numeric, GM",
			cars.AddCars(moreCars).
				Filter(hof.ByDomestic()).
				Filter(hof.ByHasNumber()).
				Filter(hof.ByMake("GM")),
		)

		hof.PrintCars("Numeric, Foreign, Map Upgraded",
			cars.Filter(hof.ByHasNumber()).
				Filter(hof.ByForegin()).
				Map(hof.Upgrade()),
		)
	}

}
