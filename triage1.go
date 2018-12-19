//======================================================================
// Go 1.11 
// Numbers classification
// date: 19.12.2018
//======================================================================
package main
 
import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	fmt.Println("// Test 1 : ")
	fmt.Println("// For numbers from 'min' to 'max', determine and print the attribute of integer divisibility by 3 and(or) 5:",
	"\n// - if number is divisible by 3 without remainder, then 'Three';",
	"\n// - if number is divisible by 5 without remainder, then 'Five';",
	"\n// - if it is divided without remainder both by 3 and 5, then 'ThreeFive'.")
	s := int(1)
	e := int(20)
	// Function time + empty time
	start := time.Now()
	tri(s, e)
	elapsed := time.Since(start)
	// Empty time
	start0 := time.Now()
	elapsed0 := time.Since(start0)
	fmt.Println("// Computing time : ", elapsed - elapsed0)
}

func tri(min int, max int) bool {
	i := min
	for i <= max {
		mess := strconv.Itoa(i)
		mess3 :=""
		mess5 :=""
		if (i%3 == 0) {
			mess = ""
			mess3 = "three"
		}
		if (i%5 == 0) {
			mess = ""
			mess5 = "five"
		}
		mess = mess + mess3 + mess5
		fmt.Println("// Number ",i," has the integer divisibility attribute : ", mess)
		i++
	}
		return true
}
//======================================================================#
// Test 1 : 
// For numbers from "nim" to "max", determine and print the attribute of integer divisibility by 3 and(or) 5: 
// - if number is divisible by 3 without remainder, then 'Three'; 
// - if number is divisible by 5 without remainder, then 'Five'; 
// - if it is divided without remainder both by 3 and 5, then 'ThreeFive'
// Number  1  has the integer divisibility attribute :  1
// Number  2  has the integer divisibility attribute :  2
// Number  3  has the integer divisibility attribute :  three
// Number  4  has the integer divisibility attribute :  4
// Number  5  has the integer divisibility attribute :  five
// Number  6  has the integer divisibility attribute :  three
// Number  7  has the integer divisibility attribute :  7
// Number  8  has the integer divisibility attribute :  8
// Number  9  has the integer divisibility attribute :  three
// Number  10  has the integer divisibility attribute :  five
// Number  11  has the integer divisibility attribute :  11
// Number  12  has the integer divisibility attribute :  three
// Number  13  has the integer divisibility attribute :  13
// Number  14  has the integer divisibility attribute :  14
// Number  15  has the integer divisibility attribute :  threefive
// Number  16  has the integer divisibility attribute :  16
// Number  17  has the integer divisibility attribute :  17
// Number  18  has the integer divisibility attribute :  three
// Number  19  has the integer divisibility attribute :  19
// Number  20  has the integer divisibility attribute :  five
// Computing time :  327.396µs
//======================================================================#
// To run this test execute in terminal:
// export GOPATH=/home/bbrs/Desktop/wildBeries/Golang/
// go install triage1
// $GOPATH/bin/triage1
//======================================================================#
