package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// Hello World API
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Colgate de esta pancho en Go!")
	}

	http.HandleFunc("/Pancho", helloHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
