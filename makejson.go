package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Make a map
	m := make(map[string]string)

	var name string
	var adrr string

	in := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your name: ")
	in.Scan()
	name = in.Text()

	fmt.Print("Enter your address: ")
	in.Scan()
	adrr = in.Text()

	// Assign value to map
	m["name"] = name
	m["adrr"] = adrr

	// Marshaling the JSON object
	json, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(json))
}
