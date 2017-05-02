package controller

import (
	"net/http"
	"encoding/json"
)

func MainController(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("Hello From the Main Controler!")
}
