package main

import "fmt"

func main() {
	cal := &Calculate{}
	cal.SetCalc(17, 19, 0)
	x, y, result := cal.calc_add()
	fmt.Printf("\n%d + %d = %d", x, y, result)

}
