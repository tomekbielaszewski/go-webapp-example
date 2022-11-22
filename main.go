package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func indexHandler(res http.ResponseWriter, req *http.Request) {
	resString := fmt.Sprintf("<html><body><h1>You've visited: %s</h1></body></html>", req.URL.String())
	res.Write([]byte(resString))
}

const PortVar = "PORT"
const DefaultPort = "8080"

func main() {
	err := LoadEnv()
	if err != nil {
		fmt.Printf("Cannot load .env file and cannot set default PORT. %s\n", err)
	}
	port := GetPort()

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		fmt.Println(fmt.Errorf("problem when starting server: %s", err))
		os.Exit(-1)
	}
}

func GetPort() string {
	port := os.Getenv(PortVar)
	if port == "" {
		port = DefaultPort
	}
	return port
}

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		err = os.Setenv(PortVar, DefaultPort)
	}
	return err
}
