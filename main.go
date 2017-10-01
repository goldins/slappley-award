package main

import (
	"os"
	"net/http"
	"log"
	"github.com/goldins/slappley-award/cmd"
)

var (
	port  string = "80"
	token string
)

type Config struct {
	Listen   string
	token string
}


func init() {
	token = os.Getenv("SCIENCE_TOKEN")
	if "" == token {
		panic("SCIENCE_TOKEN is not set!")
	}

	if "" != os.Getenv("PORT") {
		port = os.Getenv("PORT")
	}
}

func main() {
	http.HandleFunc("/", cmd.ScienceHandler())
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
