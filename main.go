package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HomePage Finish")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	test := add(5, 6)
	b := square(5)
	fmt.Println(test, b)
	handleRequest()
}

func add(a int, b int) int {
	return a + b + 2
}

func square(a int) int {
	return a * a
}
func squareV2(a int, b int) int {
	return a*a + b*b
}
