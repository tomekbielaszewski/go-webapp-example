package main

import (
	"fmt"
	"net/http"
	"os"
)

func indexHandler(res http.ResponseWriter, req *http.Request) {
	resString := fmt.Sprintf("<html><body><h1>You've visited: %s</h1></body></html>", req.URL.String())
	res.Write([]byte(resString))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(fmt.Errorf("problem when starting server: %s", err))
		os.Exit(-1)
	}
}
