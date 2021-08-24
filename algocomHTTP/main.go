package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Oláaa</h1><br><marquee>oiii</marquee>")
	})

	fmt.Print("Starting server at port 1818")
	http.ListenAndServe(":1818", nil)

}