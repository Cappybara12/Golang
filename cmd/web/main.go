package main

import (
	"fmt"
	"net/http"

	"github.com/aksha/my_project/pkg/handlers"
)

const portnumber = ":8081"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf("Starting the app on port %s\n", portnumber)
	err := http.ListenAndServe(portnumber, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
