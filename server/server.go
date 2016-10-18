package main

import "net/http"
import "fmt"
import "html"
import "log"

func witHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/witticism", witHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
