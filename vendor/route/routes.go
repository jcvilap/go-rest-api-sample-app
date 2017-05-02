package route

import (
	"net/http"
	"controller"
)

func Route() {
	// App Routes
	http.HandleFunc("/", controller.Main)
}