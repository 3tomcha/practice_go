package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		word := r.FormValue("searchword")
		log.Printf("searchword = %s\n", word)

		words, ok := r.Form["searchword"]
		log.Printf("search words = %v has values %v\n", words, ok)

		log.Print("all queries")
		for key, values := range r.Form {
			log.Printf("  %s: %v\n", key, values)
		}
	})

	http.ListenAndServe(":8888", nil)
}
