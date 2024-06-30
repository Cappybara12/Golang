package handlers

import (
	"net/http"

	"github.com/aksha/my_project/pkg/config"
	"github.com/aksha/my_project/pkg/models"
	"github.com/aksha/my_project/pkg/render"
)

// holds data sent to he handlers
// repo the rpeository used by hadnekrs
var Repo *Repository

// Reposiroty is the repository type
type Repository struct {
	App *config.AppConfig
}

// created the new repo
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// new handlers set the repository to hadnlers
func Newhandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.Templatedata{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//perforn some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again"

	render.RenderTemplate(w, "about.page.tmpl", &models.Templatedata{
		StringMap: stringMap,
	})
}
