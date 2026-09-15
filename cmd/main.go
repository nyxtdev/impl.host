package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		file := "web/index.html"

		switch request.URL.Path {
		case "/", "/index.html":
			file = "web/index.html"
		case "/other", "/other.html":
			file = "web/other.html"
		default:
			http.NotFound(writer, request)
			return
		}

		http.ServeFile(writer, request, file)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
