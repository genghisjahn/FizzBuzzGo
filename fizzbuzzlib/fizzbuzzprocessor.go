package fizzbuzz

import(
	"strconv"
)

type FizzBuzzProcessor struct{
	Items
}

type IFizzBuzzItem interface{
	GetResult(num int) string
}

type Items []IFizzBuzzItem

func(fbp *FizzBuzzProcessor) GetResult(num int) string{
	result:=""
	numstr:=strconv.Itoa(num)
	for i := 0; i < len(fbp.Items); i++ {			
			fbd:=fbp.Items[i]
			result=fbd.GetResult(num)
			if(numstr!=result){break}
	}
	return result
}

