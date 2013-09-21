package fizzbuzz

import(
	"strconv"
	"math"
)

type FizzBuzzPerfectSquare struct {}
func(fbps FizzBuzzPerfectSquare) GetResult(num int) string{
	result:=strconv.Itoa(num)
	root:=math.Sqrt(float64(num))
	introot:=int(root)
	if root==float64(introot){
		result = "Perfect Square"
	}
	return result
}