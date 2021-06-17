package basics

func sqrt(x float64) float64 {
	var y float64
	z := 1.0

	for range [15]int{} {
		y = (z*z - x) / (2 * z)
		if y == 0 {
			break
		} else {
			z -= y
		}
	}

	return z
}

func MainSqrt() {
	sqrt(4)
}
