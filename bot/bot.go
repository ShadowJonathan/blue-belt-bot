package main

import (
	"io/ioutil"
	"strings"

	"fmt"

	"../belter"
)

func main() {
	token, err := ioutil.ReadFile("../token")
	if err != nil {
		fmt.Println("Error reading token file: " + err.Error())
		err := ioutil.WriteFile("token", token, 9000)
		if err != nil {
			fmt.Println("Error writing sample token file: " + err.Error())
		}
		return
	}
	Belt.Initialize(strings.TrimSpace(string(token)))
}
