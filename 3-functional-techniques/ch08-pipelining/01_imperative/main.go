package main

import (
	"fmt"
	"log"

	gc "github.com/go-goodies/go_currency"
)

type Order struct {
	Ordernumber     int
	IsAuthenticated bool
	IsDecrypted     bool
	Credentials     string
	CCardNumber     string
	CCardExpDate    string
	LineItems       []LineItem
}

type LineItem struct {
	Description string
	Count       int
	PriceUSD    gc.USD
}

type Filterer interface {
	Filter(input chan Order) chan Order
}

func GetOrders() []*Order {
	order1 := &Order{
		1001,
		false,
		false,
		"alice,secret",
		"7b/HWvtIB9a16AYk+Yv6WWwer3GFbxpjoR+GO9iHIYY=",
		"0922",
		[]LineItem{
			LineItem{"Apples", 1, gc.USD{4, 50}},
			LineItem{"Oranges", 4, gc.USD{12, 00}},
		},
	}

	order2 := &Order{
		10002,
		true,
		false,
		"bob,secret",
		"EOc3kF/OmxY+dRCaYRrey8h24QoGzVU0/T2QKVCHb1Q=",
		"0123",
		[]LineItem{
			{"Milk", 2, gc.USD{8, 00}},
			{"Sugar", 1, gc.USD{2, 25}},
			{"Salt", 3, gc.USD{3, 75}},
		},
	}

	orders := []*Order{order1, order2}
	return orders
}

func Authenticate(o Order) Order {
	fmt.Printf("Order %d is Authenticate\n", o.Ordernumber)
	return o
}

func Decrypt(o Order) Order {
	fmt.Printf("Order %d is Decrypted\n", o.Ordernumber)
	return o
}

func Charge(o Order) Order {
	fmt.Printf("Order %d is Charged\n", o.Ordernumber)
	return o
}

func Pipeline(o Order) Order {
	o = Authenticate(o)
	o = Decrypt(o)
	o = Charge(o)
	return o
}

func main() {
	/*
		orders := GetOrders()
		for _, order := range orders {
			fmt.Printf("Processed order: %v\n", Pipeline(*order))
		}
	*/

	/* use channel
	input := make(chan Order)
	output := make(chan Order)

	go func() {
		for order := range input {
			output <- Pipeline(order)
		}
	}()

	orders := GetOrders()
	for _, order := range orders {
		fmt.Printf("Processed order: %v\n", Pipeline(*order))
	}
	close(input)
	*/

	/* use buffer */
	/*
		orders := GetOrders()
		numberOfOrders := len(orders)

		input := make(chan Order, numberOfOrders)
		output := make(chan Order, numberOfOrders)

		for i := 0; i < numberOfOrders; i++ {
			go func() {
				for order := range input {
					output <- Pipeline(order)
				}
			}()
		}

		for _, order := range orders {
			input <- *order
		}
		close(input)

		for i := 0; i < numberOfOrders; i++ {
			fmt.Println("The result is: ", <-output)
		}
	*/

	pipeline := BuildPipeline(Authenticate{}, Decrypt{}, Charge{})

	go func() {
		orders := GetOrders()
		for _, order := range orders {
			fmt.Printf("order: %v\n", order)
			pipeline.Send(*order)
		}
		log.Println("Close Pipeline")
		pipeline.Close()
	}()

	pipeline.Receive(func(o Order) {
		log.Printf("Received: %v", o)
	})

}

func BuildPipeline(filters ...Filterer) Filter {
	source := make(chan Order)
	var nextFilter chan Order
	for _, filter := range filters {
		if nextFilter == nil {
			nextFilter = filter.Filter(source)
		} else {
			nextFilter = filter.Filter(nextFilter)
		}
	}

	return Filter{input: source, output: nextFilter}
}

type Filter struct {
	input  chan Order
	output chan Order
}

func (f *Filter) Send(order Order) {
	f.input <- order
}

func (f *Filter) Receive(callback func(Order)) {
	for o := range f.output {
		callback(o)
	}
}

func (f *Filter) Close() {
	close(f.input)
}
