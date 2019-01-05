package captcha

func Generate(pattern int, firstOp int, operator int, secondOp int) string {
	switch pattern {
	case 1:
		return "1 + One"
	case 2:
		return "One + 1"
	default:
		return "default"
	}
}
