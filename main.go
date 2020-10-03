package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Serving...")

	hh := handlers.NewHello()
	http.ListenAndServe(":9000", nil)

}
