package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/aksha/my_project/pkg/config"
	"github.com/aksha/my_project/pkg/models"
)

var functions = template.FuncMap{

}
var app *config.AppConfig
//new templates set the confif frro temaptle apcakge 
func NewTemplates(a *config.AppConfig){
	app=a
}
func AddDefaultData(td *models.Templatedata) *models.Templatedata{
	//add whatevr you wnat to add and then it will be parsed on the page 
	return td 
}
// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.Templatedata) {
	var tc map[string]*template.Template
	if app.UseCache{
		tc =app.TemplateCache
	}else{
		tc,_=CreateTemplateCache()
	}
	// get the template cache from the app config

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("cpuld not get tamptle from cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)
	td=AddDefaultData(td)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}

}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
