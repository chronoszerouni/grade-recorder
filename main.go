package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	port := 8080
	log.Printf("Server is listening at port %d", port)
	log.Printf( "Please visit http://localhost:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
