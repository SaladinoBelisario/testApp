package main

import (
	"fmt"
	"math"
)

// s = ½ a t2 + vot + so

func main() {
	var acc float64
	var iVel float64
	var iDis float64
	var time float64

	fmt.Print("Enter acceleration: ")
	fmt.Scan(&acc)
	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&iVel)
	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&iDis)

	fmt.Print("Enter time to calc: ")
	fmt.Scan(&time)

	fn := GenDisplaceFn(acc, iVel, iDis)
	fmt.Println(fn(time))

}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return (1/2)*(a*math.Pow(t, 2)) + (v0 * t) + s0
	}
}
