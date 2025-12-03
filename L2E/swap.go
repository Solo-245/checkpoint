package student

func Swap(a *int, b *int) {
	x := *a
	y := *b
	*a = y
	*b = x
}
