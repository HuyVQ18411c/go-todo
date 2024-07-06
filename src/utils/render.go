package utils

import (
	"html/template"
	"net/http"
)

func Render(response http.ResponseWriter, templatePath string, data any) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}

	t.Execute(response, data)
}
