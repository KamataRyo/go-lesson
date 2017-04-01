package main

import (
	"fmt"
	"strconv"
)

func main() {
	helloWorld()
	fizzBuzz(100)
	array()
}

func helloWorld() {
	fmt.Printf("Hello, World\n")
}

func fizzBuzz(count int) {
	if count < 1 {
		return
	}
	for i := 1; i < count+1; i++ {
		var line string

		if i%3 != 0 && i%5 != 0 {
			line = strconv.Itoa(i)
		} else {
			if i%3 == 0 {
				line += "Fizz"
			}
			if i%5 == 0 {
				line += "Buzz"
			}
		}

		fmt.Printf(line + "\n")
	}
}

func array() {
	var arr [100]int
	arr[0] = 1
	arr[1] = 1
	for i := 2; i < 100; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	for i := 0; i < 100; i++ {
		fmt.Printf(strconv.Itoa(arr[i]) + "\n")
	}
}
