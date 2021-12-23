/* Explanation
------------------------------------------------------------------------------------------------------
Race Condition
A data race happens when two goroutines access the same variable concurrently, and at least one of the
access is a write instruction.

This is an example of race condition in where 2 goroutines tries to read & write sharedInt and there
is no access control.
*/

package main

import "time"

var sharedInt = 0
var unusedValue = 0

func main() {
	go runSimpleReader()
	go runSimpleWriter()
	time.Sleep(10 * time.Second)
}

func runSimpleReader() {
	for {
		var val = sharedInt
		if val%10 == 0 {
			unusedValue = unusedValue + 1
		}
	}
}

func runSimpleWriter() {
	for {
		sharedInt = sharedInt + 1
	}
}
