package main

import "fmt"

func main() {
	var numbers []int

	for i := 0; i < 10; i++ {
		fmt.Print("Enter a number: ")
		fmt.Scan(&numbers)
	}

	fmt.Print("Ordering with bubblesort: ")
	BubbleSort(numbers)

}

func BubbleSort(numbers []int) []int {
	//Start the loop in reverse order, so the loop will start with length
	//which is equal to the length of input array and then loop until reaches 1
	for i := len(numbers); i > 0; i-- {
		//The inner loop will first iterate through the full length
		//the next iteration will be through n-1
		// the next will be through n-2 and so on
		for j := 1; j < i; j++ {
			if numbers[j-1] > numbers[j] {
				Swap(numbers, j)
			}
		}
	}
	return numbers
}

func Swap(slice []int, index int) {
	slice[index] = slice[index] ^ slice[index-1]
	slice[index-1] = slice[index] ^ slice[index-1] // now index-1 is index
	slice[index] = slice[index] ^ slice[index-1]   // now index is index-1
}
