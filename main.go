package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(100.0, 20.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result of division is", result)
}

func divide(s, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	result := s / y
	return result, nil
}
