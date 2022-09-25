package main

import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
