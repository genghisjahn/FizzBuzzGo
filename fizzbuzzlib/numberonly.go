package fizzbuzz

import(
	"strconv"
)
type FizzBuzzNumberOnly struct{}	

func (fbno FizzBuzzNumberOnly) GetResult(num int) string {
	return strconv.Itoa(num) + "."
}