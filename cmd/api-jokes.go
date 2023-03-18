package main

import (
	"fmt"
	"log"
	"net/http"
	"test_app/controller"
)

func main() {
	fmt.Println("Starting server in port 8000")

	http.HandleFunc("/api/jokes", controller.GetJokes)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("Cannot init the server \n\n", err)
	}
}
