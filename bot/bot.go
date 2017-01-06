package main

import (
	"io/ioutil"
	"strings"

	"../belter"
)

func main() {
	token, _ := ioutil.ReadFile("token")
	Belt.Initialize(strings.TrimSpace(string(token)))
}
