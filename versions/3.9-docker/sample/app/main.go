package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pottava/docker-gulp/versions/3.9-docker/sample/app/module"
)

func main() {
	http.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, module.Hello(r.URL.Path[len("/"):]))
	}))
	log.Fatal(http.ListenAndServe(":80", nil))
}
