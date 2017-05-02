package route

import (
	"net/http"
	"controller"
)

func NewRouter() {
	// App Routes
	http.HandleFunc("/", controller.Main)
}