package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()
	router.Host(BaseUrl)
	log.Fatal(http.ListenAndServe(":2020", router))
}
