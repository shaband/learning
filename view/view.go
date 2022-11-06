package view

import (
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
)

type Template struct {
	htmlTmp *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {

	err := t.htmlTmp.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func Parse(filename string) (*Template, error) {

	t, err := template.ParseFiles(filepath.Join("templates", filename))
	if err != nil {
		return nil, err
	}
	return &Template{
		htmlTmp: t,
	}, nil
}
func ParseFS(Fs fs.FS, patterns ...string) (*Template, error) {

	layout, err := fs.Glob(Fs, "layout/*.gohtml")
	if err != nil {
		panic(err)
		return nil, err
	}

	patterns = append(patterns, layout...)
	t, err := template.ParseFS(Fs, patterns...)
	if err != nil {
		return nil, err
	}
	return &Template{
		htmlTmp: t,
	}, nil
}
func Must(temp *Template, err error) *Template {

	if err != nil {
		panic(err)
	}
	return temp
}
