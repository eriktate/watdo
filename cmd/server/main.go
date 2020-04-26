package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eriktate/watdo/env"
)

func main() {
	host := env.GetString("WATDO_HOST", "localhost")
	port := env.GetUint("WATDO_PORT", 8080)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), http.HandlerFunc(helloHandler)); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, world!"))
}
