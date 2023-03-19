package main

import (
	"cb-ldp-backend/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	ginRouter := router.SetUpRouter()
	fmt.Println("Application loaded successfully ")
	log.Fatal(http.ListenAndServe(":4000", ginRouter))
}
