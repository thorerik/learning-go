package main

import "fmt"

func f(from string) {
	for i := 0; i < 50; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine1")
	go f("goroutine2")

	go func(msg string) {
		for i := 0; i < 50; i++ {
			fmt.Println(msg, ":", i)
		}
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
