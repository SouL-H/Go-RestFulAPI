package main

import (
	"net/http"
)
func aboutHandler(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("About page"))
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	http.HandleFunc("/about",aboutHandler)
	http.ListenAndServe(":8081", nil)
}
