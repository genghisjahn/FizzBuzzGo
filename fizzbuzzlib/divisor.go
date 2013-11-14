package fizzbuzz

import(
	"strconv"
)
type FizzBuzzDivisor struct{
	Value int //Value to be tested.
	Message string //Message if a condition is met.
}	
func (fbd FizzBuzzDivisor) GetResult(num int) string {
	result:=strconv.Itoa(num)
	if num % fbd.Value==0{
		result = fbd.Message
	} 
	return result
}

//another test

