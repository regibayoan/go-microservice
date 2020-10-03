package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Serving...")
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Inside / ")
		d, _ := ioutil.ReadAll((r.Body))
		fmt.Fprintf(rw, "Hello %s", d)
	})
	http.HandleFunc("/abc", func(http.ResponseWriter, *http.Request) {
		fmt.Println("Inside /abc ")
	})
	http.ListenAndServe(":9000", nil)

}
