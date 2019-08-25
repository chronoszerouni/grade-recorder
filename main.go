package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "text/html")
		html := `
<Doctype html>
<html>
<body>
<a href="http://baidu.com" target="_blank">Hello world!</a>
</body>
</html>
`
		fmt.Fprintln(writer, html)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
