package main

import (
	"github.com/goldins/slappley-award/cmd"
	"os"
	"net/http"
	"log"
	"fmt"
)

var (
	port  string = "80"
	token string
)

type Config struct {
	Listen   string
	token string
}

const SLACK_TOKEN = "SCIENCE_TOKEN"

func init() {
	token = os.Getenv(SLACK_TOKEN)
	if "" == token {
		panic(fmt.Sprintf("%s is not set!", SLACK_TOKEN))
	}

	if "" != os.Getenv("PORT") {
		port = os.Getenv("PORT")
	}
}

func main() {
	http.HandleFunc("/science", cmd.ScienceHandler())
	http.HandleFunc("/action", cmd.SendHandler())
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
