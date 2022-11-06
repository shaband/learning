package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/shaband/learning/controllers"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", controllers.StaticPage("home.gohtml"))
	r.Get("/about", controllers.StaticPage("about.gohtml"))
	r.Get("/contract", controllers.StaticPage("contract.gohtml"))
	r.Get("/faq", controllers.FAQPage("faq.gohtml"))
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}

}
