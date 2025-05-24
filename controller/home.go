package controller

import (
	"datasystem/model"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := model.GenerateFakeData()
	tmpl := template.Must(template.ParseFiles("view/index.html"))
	tmpl.Execute(w, data)
}
