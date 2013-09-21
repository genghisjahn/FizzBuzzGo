package fizzbuzz

import(
	"strconv"
)

type FizzBuzzPalindrome struct{
	
}

func(fbpalin FizzBuzzPalindrome) GetResult(num int) string{
	result:=""
	normal:=strconv.Itoa(num)
	result=normal
	revstr := Reverse(normal)
	if normal == revstr{
		result="Palindrome"
	}
	return result
}


func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}