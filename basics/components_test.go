package basics

import "testing"

func TestSum(t *testing.T) {
	cases := []struct {
		in   []float64
		want float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 15},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8}, 36},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 55},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, 91},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 120},
	}

	for _, c := range cases {
		got := Sum(c.in...)
		if got != c.want {
			t.Errorf("Sum(%v) got %v instead of %v", c.in, got, c.want)
		}
	}
}

func TestMultiply(t *testing.T) {
	cases := []struct {
		in   []float64
		want float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 120},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8}, 40320},
		{[]float64{1.45, 2.230, -3, -10}, 97.005},
		{[]float64{42.72, 12, 48}, 24606.72},
		{[]float64{8, 9, 14, 15}, 15120},
	}

	for _, c := range cases {
		got := Multiply(c.in...)
		if got != c.want {
			t.Errorf("Multiply(%v) got %v instead of %v", c.in, got, c.want)
		}
	}
}

func TestSubtract(t *testing.T) {
	cases := []struct {
		in   []float64
		want float64
	}{
		{[]float64{1, 2, 3, 4, 5}, -13},
		{[]float64{12, 8, 2}, 2},
		{[]float64{5, 1, 2, 2}, 0},
		{[]float64{9, 4, 7}, -2},
		{[]float64{45861, 546, 888, 10101}, 34326},
	}

	for _, c := range cases {
		got := Subtract(c.in...)
		if got != c.want {
			t.Errorf("Subtract(%v) got %v instead of %v", c.in, got, c.want)
		}
	}
}

func TestSqrt(t *testing.T) {
	cases := []struct {
		in   []float64
		want []float64
	}{
		{[]float64{1, 4, 9, 16, 25}, []float64{1, 2, 3, 4, 5}},
		{[]float64{27.5625, 11.9716, 144}, []float64{5.25, 3.46, 12}},
		{[]float64{81, 64, 225}, []float64{9, 8, 15}},
		{[]float64{121, 36}, []float64{11, 6}},
		{[]float64{100, 5625, 1089, 1702.3876, 256, 8317.44}, []float64{10, 75, 33, 41.26, 16, 91.2}},
	}

	for _, c := range cases {
		got := Sqrt(c.in...)
		for i := 0; i < len(c.want); i++ {
			if got[i] != c.want[i] {
				t.Errorf("Sqrt(%v) got %v instead of %v", c.in, got, c.want)
			}
		}
	}
}

func TestConvertBase(t *testing.T) {
	cases := []struct {
		base int
		num  []int
		want []string
	}{
		{2, []int{23, 2, 6, 8, 16, 11, 13, 10}, []string{"10111", "10", "110", "1000", "10000", "1011", "1101", "1010"}},
		{8, []int{23, 2, 6, 8, 16, 11, 13, 10}, []string{"27", "2", "6", "10", "20", "13", "15", "12"}},
		{10, []int{23, 2, 6, 8, 16, 11, 13, 10}, []string{"23", "2", "6", "8", "16", "11", "13", "10"}},
		{16, []int{23, 2, 6, 8, 16, 11, 13, 10}, []string{"17", "2", "6", "8", "10", "B", "D", "A"}},
		{2, []int{}, []string{}},
	}

	for _, c := range cases {
		got := ConvertBase(c.base, c.num...)
		for i := 0; i < len(c.want); i++ {
			if got[i] != c.want[i] {
				t.Errorf("ConvertBase(%v, %v) got %v instead of %v", c.base, c.num, got, c.want)
			}
		}
	}
}

func TestDivide(t *testing.T) {
	cases := []struct {
		in   []float64
		want float64
	}{
		{[]float64{25, 5, 4}, 1.25},
		{[]float64{68, 8, 4}, 2.125},
		{[]float64{46, 4, 5}, 2.3},
		{[]float64{144, 6, 8}, 3},
		{[]float64{153, 5, 4}, 7.65},
	}

	for _, c := range cases {
		got := Divide(c.in...)
		if got != c.want {
			t.Errorf("Divide(%v) got %v instead of %v", c.in, got, c.want)
		}
	}
}
