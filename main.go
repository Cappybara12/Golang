package main

import (
	"fmt"

	"github.com/aksha/my_project/helpers"
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	// Example usage of CalculateValue
	go CalculateValue(intChan)

	// Receive the calculated value from the channel
	value := <-intChan
	fmt.Println("Random number generated:", value)
}
