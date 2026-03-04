package main

import "fmt"

func main() {
	/*
		// calculate.go
		cal := &Calculate{}
		cal.SetCalc(17, 19, 0)
		x, y, result := cal.calc_add()
		fmt.Printf("\n%d + %d = %d", x, y, result)
		// greeting.go
		gree := &Greeting{}
		gree.SetGreeting("", 0)
		greetig, current_time := gree.showGreeting()
		fmt.Printf("\nits %d now and %s", current_time, greetig)
	*/

	// inputname.go
	input := &Input_name{}
	input.SetInput_name("")
	fmt.Print("please input your name: ")
	input.do_inputName()
	inputname := input.GetInput_name()
	fmt.Printf("\nso you are %s", inputname)
	/*
		var i int

		fmt.Print("Type a number: ")
		fmt.Scan(&i)
		fmt.Println("Your number is:", i)
	*/

}
