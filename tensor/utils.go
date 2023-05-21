package tensor

func docMul(x, y []int) int {
	s := 0
	for j := 0; j < len(x); j++ {
		s += x[j] * y[j]
	}
	return s
}

func sumInt(x []int) int {
	s := 0
	for i := 0; i < len(x); i++ {
		s += x[i]
	}
	return s
}

func prodInt(x []int) int {
	s := 1
	for i := 0; i < len(x); i++ {
		s *= x[i]
	}
	return s
}
