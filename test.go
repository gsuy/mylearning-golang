package main

import (
	"fmt"
	"test/mo" //syntax `[root-path]/<package-name>`it's Package alias ilke `as` in python

	"rsc.io/quote"
)

func main() {

	// topic : variable
	// static: var var_name data_type = value ex. var a int = 1
	// dynamic: var_name = value ex. a := 1
	var a int = 1
	b := 2
	a, b = b, a // swap
	fmt.Println(a, b)

	// topic : printf
	var s int
	var ss float64
	var sss bool
	var ssss string
	fmt.Printf("%v %v %v %q\n", s, ss, sss, ssss)

	// topic : if
	var va int = 1
	if va == 1 {
		fmt.Println(va)
	}

	// topic : for
	sum := 0
	for c := 1; c <= 10; c++ {
		sum += 1 // or sum++
	}
	fmt.Println(sum)

	// topic : while
	o := 1
	for o < 10 { // while condition
		fmt.Println(o)
		o += 1
	}

	count := 0
	for { // while true
		count++
		fmt.Println(count)
		if count == 3 {
			break
		}
	}

	// topic : function
	fmt.Println(sumtwo(34, 55))
	fmt.Println(identity(34, 55))

	// topic : local variable in if
	if i1, i2 := identity(2, 4); i1 < 10 && i2 < 10 {
		fmt.Println(i1, i2)
	}

	// topic : package
	// make sure the first letter of the identifier name for a variable or function is capitalized to public variable or function (first letter name isn’t capitalized to private)
	fmt.Println(quote.Go())
	// fmt.Println(mo.Robo(1, 2))
	// mo.Ttt()
	mo.Ttt()
}

func sumtwo(x int, y int) int { // or func sumtwo( x,y int ) int {
	return x + y
}

func identity(x, y float64) (float64, float64) {
	return x, y
}

// or can implement like this ***
// func identity(x,y float64) ( xx float64, yy float64 ) {
// 	return
// } it will return xx and yy without declare
