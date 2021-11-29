package main

import "fmt"

func main() {
	oddChan := make(chan int)
	evenChan := make(chan int)

	fmt.Println("Enter upto the number you want to print....")

	var limit int
	fmt.Scanln(&limit)
	go printNumbers(oddChan, evenChan)
	for i := 1; i <= limit; i++ {
		if i%2 == 0 {
			evenChan <- i
		} else {
			oddChan <- i
		}
	}
}

func printNumbers(odd, even <-chan int) {
	for {
		select {
		case val := <-odd:
			fmt.Println(val)
		case val := <-even:
			fmt.Println(val)
		}
	}
}
