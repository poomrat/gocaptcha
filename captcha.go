package captcha

import (
	"fmt"
)

func Generate(pattern int, firstOp int, operator int, secondOp int) string {
	switch pattern {
	case 1:
		return fmt.Sprintf("%d %s %s", firstOp, "+", toWord(secondOp))
	case 2:
		return fmt.Sprintf("%s %s %d", toWord(firstOp), "+", secondOp)
	default:
		return "default"
	}
}

func toWord(num int) string {
	switch num {
	case 0:
		return "Zero"
	case 1:
		return "One"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	default:
		return "Unknown"
	}
}
