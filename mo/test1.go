package mo

import "fmt"

func Robo(x, y int) int {
	tt()
	return x + y
}

func tt() {
	fmt.Println("call package mo => test1")
}
