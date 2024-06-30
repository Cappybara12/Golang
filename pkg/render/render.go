package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
	err = parsedTemplate.Execute(w, nil)
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//check to see if he have already temaptle in our cache
	_, inMap := tc[t]
	if !inMap {

		//need to create template
		err = createTemplateCache(t)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		//we have one in cahce
		fmt.Println("using cache template")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	//add teamptle to cache(map)
	tc[t] = tmpl
	return nil
}
