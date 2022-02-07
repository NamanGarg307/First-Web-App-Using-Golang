package main

import (
	"fmt"
	"net/http"
	"github.com/namangarg307/First-Web-App-In-Go/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println("Web server started at port number", portNumber)
	http.ListenAndServe(portNumber, nil)
}
