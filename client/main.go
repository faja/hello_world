package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	r, e := http.Get("http://localhost:8080/")
	if e != nil {
		log.Fatalf("error: %v\n", e)
	}

	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(b))

}
