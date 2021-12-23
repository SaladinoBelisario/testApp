package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	firstName string
	lastName  string
}

func main() {
	var fileName string
	nameSli := make([]Name, 0)
	var nameObj Name

	fmt.Print("Enter file name: ")
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		if len(s[0]) > 20 {
			s[0] = twentyChar(s[0])
		}
		if len(s[1]) > 20 {
			s[1] = twentyChar(s[1])
		}

		nameObj.firstName, nameObj.lastName = s[0], s[1]
		nameSli = append(nameSli, nameObj)
	}

	for _, v := range nameSli {
		fmt.Println(v.firstName, v.lastName)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func twentyChar(s string) string {
	runes := []rune(s)
	return string(runes[0:20])
}
