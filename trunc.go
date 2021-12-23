package main

import "fmt"

func trunc() {
	var x float32
	fmt.Println("Enter a float number:")
	fmt.Scan(&x)
	fmt.Printf("The truncated number is %.2f\n", x)
}
