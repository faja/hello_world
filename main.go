package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world\n")
}

func handlerLoad(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile("/proc/loadavg")
	fmt.Fprint(w, string(data))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/loadavg", handlerLoad)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
