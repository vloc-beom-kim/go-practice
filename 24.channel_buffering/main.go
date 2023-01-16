package main

import "fmt"

func main() {
	messages := make(chan string)
	messages <- "buffered"
	messages <- "buffered"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	// fmt.Println(<-messages)

}
