package serve

import (
	"html/template"
	"net/http"
)

func Template(writer http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("template/index.html")
	if err != nil {
		panic("parse template error")
	}
	temp.Execute(writer, nil)
}
