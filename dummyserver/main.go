package main

import (
	"fmt"
	"net/http"

	"github.com/skkim-01/vue-practice/dummyserver/api"
)

func main() {
	// handle functions
	http.HandleFunc("/api/login", api.Login)
	http.HandleFunc("/api/open/count", api.OpenCount)
	// serve static
	http.Handle("/", http.FileServer(http.Dir("./static")))
	// start server with 9999
	fmt.Println("dummyserver is started with :9999")
	http.ListenAndServe(":9999", nil)

}
