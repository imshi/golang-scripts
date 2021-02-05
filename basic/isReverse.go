package main

import "fmt"

func isReverse(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	// 记录每个字符串出现的字数
	m := make(map[rune]int)
	n := make(map[rune]int)

	for _, r := range a {
		m[r]++
	}

	for _, r := range b {
		n[r]++
	}

	for i, r := range m {
		if n[i] != r {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("The resoult is :%t", isReverse("asd", "sad"))
}
