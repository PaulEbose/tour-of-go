package basics

import "testing"

func TestSum(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 36},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 55},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, 91},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 120},
	}
	for _, c := range cases {
		got := Sum(c.in...)
		if got != c.want {
			t.Errorf("Sum(%v) got %v instead of %v", c.in, got, c.want)
		}
	}
}

func TestConvertBase(t *testing.T) {
	cases := []struct {
		base int
		num  int
		want int
	}{
		{2, 6, 110},
		{2, 23, 10111},
		{8, 4, 4},
		{8, 15, 17},
		{10, 188, 188},
	}
	for _, c := range cases {
		got := ConvertBase(c.base, c.num)
		if got != c.want {
			t.Errorf("ConvertBase(%v, %v) got %v instead of %v", c.base, c.num, got, c.want)
		}
	}
}

func TestDivide(t *testing.T) {
	cases := []struct {
		in   []int
		want float64
	}{
		{[]int{25, 5, 4}, 1.25},
		{[]int{68, 8, 4}, 2.125},
		{[]int{46, 4, 5}, 2.3},
		{[]int{144, 6, 8}, 3},
		{[]int{153, 5, 4}, 7.65},
	}
	for _, c := range cases {
		got := Divide(c.in...)
		if got != c.want {
			t.Errorf("Divide(%v) got %v instead of %v", c.in, got, c.want)
		}
	}
}
