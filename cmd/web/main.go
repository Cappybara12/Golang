package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aksha/my_project/pkg/config"
	"github.com/aksha/my_project/pkg/handlers"
	"github.com/aksha/my_project/pkg/render"
)

const portNumber = ":8081"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.Newhandlers(repo)
	render.NewTemplates(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
