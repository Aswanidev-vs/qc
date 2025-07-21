package qc

func Add(a, b int) int {
	return a + b
}
func Sub(b, c int) int {
	return b - c
}
func Multi(c, d int) int {

	return c * d
}

func Div(d, e float64) float64 {

	return d / e
}

func Prime(n int) int {
	if n <= 1 {

		return 0
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return 0
		}
	}

	return n
}

func Fact(n int) int {
	if n < 0 {

		return 0
	}
	f := 1
	for n > 1 {
		f *= n
		n--
	}
	return f
}
