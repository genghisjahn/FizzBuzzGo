package fizzbuzz

import (
	"strconv"
)

type FizzBuzzIs42 struct{}

func (fb42 FizzBuzzIs42) GetResult(num int) string {
	result := strconv.Itoa(num)
	if num == 42 {
		result = "The ultimate answer!!"
	}
	return result
}
