package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	fmt.Println("Experimenting with break in a for loop")
	// write a program to break if n if divisble by 7
	// and only print odd numbers start from 1
	i = 1
	for i <= 10 {
		if i%2 != 0 {
			if i%7 == 0 {
				break
			}
			fmt.Println(i)
		}
		i += 1
	}
}
