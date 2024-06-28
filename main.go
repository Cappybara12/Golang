package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myjson := `
[
		{
		"first_name":"Clark",
		"last_name":"Kent",
		"hair_color":"Black",
		"has_dog":true
		},
		{
		"first_name":"Bruce",
		"last_name":"Wayne",
		"hair_color":"Black",
		"has_dog":false
		}
	]
	`

	var unmarshalled []User
	//we jsut basically unmarsheeled the json we got using aour strcut that had similarly dieclared values
	err := json.Unmarshal([]byte(myjson), &unmarshalled)
	if err != nil {
		fmt.Println("error unmarshalling json:", err)
		return
	}
	fmt.Println(err)
	fmt.Printf("unmarshalled: %+v\n", unmarshalled)

	// write json from struct
	var mySlice []User
	var m1 User
	m1.FirstName = "Wallly"
	m1.LastName = "Wallly"
	m1.HairColor = "brown"
	m1.HasDog = false

	mySlice = append(mySlice, m1)
	var m2 User
	m2.FirstName = "Waly2"
	m2.LastName = "Waly23"
	m2.HairColor = "brwn"
	m2.HasDog = false
	mySlice = append(mySlice, m2)
// during marshalinhg that is struct to json we get an err and a json so we handle it here 
	newJson, err := json.MarshalIndent(mySlice, "", "		")
	if err != nil {
		log.Println("error handling", err)
	}
	fmt.Println(string(newJson))

}
