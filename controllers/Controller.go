package controllers

import (
	"html/template"
	"net/http"

	"github.com/shaband/learning/templates"
	"github.com/shaband/learning/view"
)

func StaticPage(tmpl string) http.HandlerFunc {

	temp := view.Must(view.ParseFS(templates.FS, tmpl))
	return func(w http.ResponseWriter, r *http.Request) {
		temp.Execute(w, nil)

	}
}
func FAQPage(tmpl string) http.HandlerFunc {

	faq := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is There is Free Verison",
			Answer:   "Yes ,We Offer a Free Trial For 30 Days",
		},
		{
			Question: "What is your support hours",
			Answer:   "we have support stuff answering emails 24/7, though response time may be bit slower on weekends",
		},
		{
			Question: "How Do I Contact Support?",
			Answer: `Email us - <a href="mailto=support@gamil.com">
			mailto=support@gamil.com </a>`,
		},
	}

	temp := view.Must(view.ParseFS(templates.FS, tmpl))
	return func(w http.ResponseWriter, r *http.Request) {
		temp.Execute(w, faq)

	}
}
