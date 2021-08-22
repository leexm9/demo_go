package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println(http.ListenAndServe("localhost:8082", nil))
}
