package main

import "fmt"

func main() {

	numbers := []int{8, 95, 21, 02, 45, 10, 52, 23, 9, 89, 80}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number, " is even")
		} else {
			fmt.Println(number, " is odd")
		}
	}

}
