package main

import (
	"net/http"

	"datasystem/controller"
)

func main() {
	http.HandleFunc("/", controller.Home)
	http.ListenAndServe(":8080", nil)
}
