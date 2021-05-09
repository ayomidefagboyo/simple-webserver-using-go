package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("hello welcome to my http local webserver built using golang")

		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "sorry man, it didnt work", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "hello %s\n", d)
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye lil nigga")
	})
	http.ListenAndServe(":9090", nil)
}
