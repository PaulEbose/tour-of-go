package basics

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	ss := strings.Fields(s)

	for _, v := range ss {
		m[v] += 1
	}

	return m
}

func WcTest() {
	wc.Test(WordCount)
}
