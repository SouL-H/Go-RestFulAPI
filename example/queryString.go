package example

import "net/http"

func mainQuery() {
	http.HandleFunc("/search", search)
	http.ListenAndServe(":8082", nil)
}

func search(w http.ResponseWriter, r *http.Request) {
	h1 := r.FormValue("h1")
	source := r.FormValue("source")

	data := "/search?h1=" + h1 + "&source=" + source
	w.Write([]byte(data))

}
