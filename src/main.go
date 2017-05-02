package main

import (
	"net/http"
	"route"
)

func main() {
	router := Load()
	http.Handle(router, nil)
}