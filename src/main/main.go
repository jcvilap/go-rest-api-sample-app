package src

import "net/http"

func main() {
	http.HandleFunc("/", someFunc)
	http.ListenAndServe(":7070", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello universe"))
}