package captcha

import (
	"fmt"
)

func Generate(pattern int, firstOp int, operator int, secondOp int) string {
	switch pattern {
	case 1:
		return fmt.Sprintf("%d %s %s", firstOp, operatorMap[operator], numberMap[secondOp])
	case 2:
		return fmt.Sprintf("%s %s %d", numberMap[firstOp], operatorMap[operator], secondOp)
	default:
		return "default"
	}
}

var operatorMap = map[int]string{
	1: "+",
	2: "-",
	3: "*",
}

var numberMap = map[int]string{
	0: "Zero",
	1: "One",
	2: "Two",
	3: "Three",
	4: "Four",
	5: "Five",
	6: "Six",
	7: "Seven",
	8: "Eight",
	9: "Nine",
}
