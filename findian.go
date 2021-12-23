package main

import (
	"fmt"
	"strings"
)

func findian() {
	var s string
	var t string
	println("Enter a string:")
	fmt.Scan(&s)

	lenS := len(s)
	if lenS < 3 {
		fmt.Println("Not Found!")
		return
	}

	t = strings.ToLower(s)

	isValidFirstCharacter := t[0] == 105
	isValidEndCharacter := t[lenS-1] == 110
	isValid := isValidFirstCharacter && isValidEndCharacter

	if isValid {
		for i := 1; i < lenS-1; i++ {
			if t[i] == 97 {
				fmt.Println("Found!")
				return
			}
		}
	}

	fmt.Println("Not Found!")
	return
}
