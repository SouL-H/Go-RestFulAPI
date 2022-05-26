package main

import (
	"net/http"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About page"))
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello user"))
	x := r.URL.Path[1:]

	w.Write([]byte(x))

}
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8081", nil)
}
