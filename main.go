package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello %s", "chrons")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
