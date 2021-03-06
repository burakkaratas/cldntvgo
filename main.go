package main

import (
	"os"
	"net/http"
	"github.com/codegangsta/negroni"

	"fmt"
	"math/rand"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	n := negroni.Classic()
	n.UseHandler(mux)
	hostString := fmt.Sprintf(":%s", port)
	n.Run(hostString)
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello Tutku!")
}

func dieRoll(size int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(size) + 1
}