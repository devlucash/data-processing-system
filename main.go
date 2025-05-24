package main

import (
	"net/http"

	"github.com/devlucash/data-processing-system/controller"
)

// replace "your-app" with your actual module name

func main() {
	http.HandleFunc("/", controller.Home)
	http.ListenAndServe(":8080", nil)
}
