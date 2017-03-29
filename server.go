package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", "158.37.63.137", "Rest Service")

func addHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}

func main() {

	flag.Parse()
	log.Printf("Start Server...")

	http.HandleFunc("/", addHandler)
	http.ListenAndServe(*addr, nil)
}
