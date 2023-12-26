package ds

import (
	"fmt"
)

type transformFn func(int) int

func ds() {
	a := []int {1,2,3}
	b := transformNumber(a, double)
	b = transformNumber(b, triple)
	fmt.Println(b)

	transformFn1 := getTransformerFunction(&a)
	c := transformNumber(a, transformFn1)
	fmt.Println(c)

	// anonymous function function type should match as the one used in other place
	// function is declared at the point of declaration itself
	d := transformNumber(a, func (number int) int {
		return number*5
	})
	fmt.Println(d)


	e:=transformNumber(a, createTransformer(10))
	fmt.Println(e)
}

func transformNumber(a []int, transform func(int) int) []int {
	// pass function as parameter value
	// we are saying a function which accepts one param int and returns int
	dNumber := []int{}
	for _, val:= range a {
		dNumber = append(dNumber, transform(val))
	}
	return dNumber
}

func double(number int) int {
	return number*2
}
func triple(number int) int {
	return number*3
}

func getTransformerFunction(numbers *[]int) transformFn {
	if(*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}


func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number*factor
	}

}