package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portnumber = ":8081"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world! This is the home page.")
}

func about(w http.ResponseWriter, r *http.Request) {
	sum := addValue(2, 3)
	fmt.Fprintf(w, fmt.Sprintf("The sum at the about page is %d", sum))
}

func divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(10.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 10.0, 10.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func addValue(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/divide", divide)

	fmt.Printf("Starting the app on port %s\n", portnumber)
	err := http.ListenAndServe(portnumber, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
