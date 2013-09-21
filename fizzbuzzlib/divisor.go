package fizzbuzz

import(
	"strconv"
)
type FizzBuzzDivisor struct{
	Value int //Value to be tested.
	Message string //Message if a condition is met.
}	
func (fbd FizzBuzzDivisor) GetResult(num int) string {
	result:=""
	if num % fbd.Value==0{
		result = fbd.Message
	} else {
		result = strconv.Itoa(num)
	}
	return result
}



