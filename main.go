package main

import (
	"fmt"
	"net/http"
)
func Home(w http.ResponseWriter , r *http.Request){



}
func main() {
	//we jsut took the rpeosne writer and request writer and request writer takes in the pointer
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// it will print but the writer also privde swith the err and bytes
		n, err := fmt.Fprintf(w, "hello world!")

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("number of bytes : %d", n))
	})
	_ = http.ListenAndServe(":8081", nil)
}
