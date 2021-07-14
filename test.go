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

	// topic pointer
	gg := "I am gg!"
	var pt *string = &gg // or pt := &gg
	fmt.Println(pt)      // => adress of gg eg. 0xc000012568
	fmt.Println(&pt)     // => adress of pt eg. 0xc000006d18
	fmt.Println(*&pt)    // => *&<something> = <something> in this ex. pt = 0xc000012568, &pt = 0xc000006d18 so *&pt = 0xc000012568
	fmt.Println(*pt)     // => value of address of gg
	*pt = "I am gg new!" // => when need to access variable to edit or something, just add * in first letter
	fmt.Println(gg)

	wp := "ssadasdadasdasd"

	for _, v := range wp {
		fmt.Printf("test %c\n", v)
	}

	// topic test char of string => ["Hi" =>  'H'+ 'i' + '\0']
	testchar := "thisisstring"
	fmt.Println(string(testchar[11])) // => g
	// fmt.Println(string(testchar[12])) // error out of index

	// topic array
	var array1 []int = []int{1, 2, 3, 4} // or ar array1 [4int = [4]int{1, 2, 3, 4} or array1 := []int{1,2,3,4} or array1 := [4]int{1,2,3,4}
	fmt.Println(array1)

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
