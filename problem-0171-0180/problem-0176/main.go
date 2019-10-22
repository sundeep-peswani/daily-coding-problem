package main

func main() {

}

func isMappable(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1CharCount, s2CharCount := generateCharCount(s1), generateCharCount(s2)
	if len(s1CharCount) != len(s2CharCount) {
		return false
	}

	for i := 0; i < len(s1CharCount); i++ {
		if s1CharCount[i] != s2CharCount[i] {
			return false
		}
	}

	return true
}

func generateCharCount(str string) []int {
	m := make(map[byte]int)
	max := 0
	for i := 0; i < len(str); i++ {
		c := str[i] - 'a'
		m[c]++

		if m[c] > max {
			max = m[c]
		}
	}

	count := make([]int, max + 1, max + 1)
	for _, v := range m {
		count[v]++
	}

	return count
}
