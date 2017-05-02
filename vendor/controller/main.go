package controller

import (
	"net/http"
)

func Main(w http.ResponseWriter, req *http.Request) {
	// Controller logic here...
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}