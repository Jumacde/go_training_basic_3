package main

import "fmt"

func main() {
	// calculate.go
	cal := &Calculate{}
	cal.SetCalc(17, 19, 0)
	x, y, result := cal.calc_add()
	fmt.Printf("\n%d + %d = %d", x, y, result)
	// greeting.go
	gree := &Greeting{}
	gree.SetGreeting("", 0)
	greetig, current_time := gree.showGreeting()
	fmt.Printf("\n its %d now and %s", current_time, greetig)

}
