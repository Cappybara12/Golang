package main

import (
	"net/http"

	"github.com/aksha/my_project/pkg/config"
	"github.com/aksha/my_project/pkg/handlers"
	"github.com/bmizerany/pat"
)

func Routes(app *config.AppConfig) http.Handler{
	mux :=pat.New()
	mux.Get("/",http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about",http.HandlerFunc(handlers.Repo.About))

	return mux
}