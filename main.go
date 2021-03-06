package main

import "fmt"

type currency func(int64) int64

func (cu currency) convert(money int64) int64 {
	return cu(money)
}

var use = "USA"

func main() {
	store := make(map[string]currency)

	//register currency converters
	store["NGN"] = toNaria
	store["USA"] = toUSA

	//declear currency converter
	var cc currency

	//find the converter in the store and return the address to the function
	v := store[use]

	//assign currency converter to the selected converter
	cc = v

	//print out result
	fmt.Println(cc.convert(222))

}

func toNaria(n int64) int64 {
	return int64(n * 360)
}

func toUSA(n int64) int64 {
	return int64(n * 666)
}
