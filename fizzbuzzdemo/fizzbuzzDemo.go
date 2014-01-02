package main

import (
	sfb "FizzBuzzGo/fizzbuzzlib"
	"fmt"
)

func main() {
	fbp := sfb.FizzBuzzProcessor{GetItems()}
	for i := 1; i <= 100; i++ {
		result := fbp.GetResult(i)
		fmt.Printf("\n%v", result)
	}
}
func GetItems() sfb.Items {
	fbpno := sfb.FizzBuzzNumberOnly{}
	var _ = fbpno
	fbps := sfb.FizzBuzzPerfectSquare{}
	fbpalin := sfb.FizzBuzzPalindrome{}
	fb42 := sfb.FizzBuzzIs42{}
	fbdiv15 := sfb.FizzBuzzDivisor{15, "FizzBuzz"}
	fbdiv5 := sfb.FizzBuzzDivisor{5, "Buzz"}
	fbdiv3 := sfb.FizzBuzzDivisor{3, "Fizz"}
	return sfb.Items{fbps, fb42, fbpalin, fbdiv15, fbdiv5, fbdiv3}
}

//I'm adding a comment.
//Add another comment.
//Changes
