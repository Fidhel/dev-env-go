package main

import (
	"fmt"
	"os"
)

func main() {

	//printStacks()
	//sampleLoop()
	//echo()
	//concat("Vuvu", "Zella")
	start()
}

func printStacks() {
	var jobs [7]string
	jobs[0] = "front"
	jobs[1] = "back"
	jobs[2] = "qa"
	jobs[3] = "ux"
	jobs[4] = "pm"
	jobs[5] = "pd"
	jobs[6] = "tm"

	for i, s := range jobs {
		fmt.Println(i, s)
	}
}

func sampleLoop() {
	var x = 0
	for {
		if x > 10 {
			x--
			break
		}
		fmt.Println(x)
		x++
	}
	for x != 0 {
		x--
		fmt.Println(x)
	}
}

func echo() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func concat(first_name string, last_name string) {
	var full_name string
	full_name += first_name
	full_name += last_name
	fmt.Println(full_name)
}

func start() {
	fmt.Println("Vuvu Start")
}
