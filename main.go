package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"html/template"
	"net/http"
	"os"
)

var tpl = template.Must(template.ParseFiles("assets/index.html"))

func indexHandler(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, req.URL.String())
	if err != nil {
		res.WriteHeader(500)
		_, _ = res.Write([]byte("ISE:" + err.Error()))
	}
}

func searchHandler(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	q := query.Get("q")
	page := query.Get("page")

	if q == "" {
		http.Error(res, "Empty query won't find any results", http.StatusNotFound)
	}
	if page == "" {
		page = "1"
	}

	fmt.Printf("Search query reached. Params: q=%s page=%s", q, page)
}

const PortVar = "PORT"
const DefaultPort = "8080"

func main() {
	err := LoadEnv()
	if err != nil {
		fmt.Printf("Cannot load .env file and cannot set default PORT. %s\n", err)
	}
	port := GetPort()

	fs := http.FileServer(http.Dir("assets"))
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/search", searchHandler)
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fmt.Println("Server started at port " + port)
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
