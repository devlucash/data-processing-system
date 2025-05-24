// controller/home.go
package controller

import (
    "net/http"
    "html/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("view/index.html"))
    tmpl.Execute(w, nil)
}
