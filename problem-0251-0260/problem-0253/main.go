package main

func main() {

}

func zigzag(input string, k int) []string {
	rows := make([][]rune, k)

	for j := range rows {
		rows[j] = make([]rune, len(input) - j)
		for i := 0; i < len(rows[j]); i++ {
			rows[j][i] = ' '
		}
	}

	a, inc := 0, 1
	for i, b := range input {
		rows[a][i] = b
		a += inc
		if a == k {
			a = k - 2
			inc = -1
		} else if a == 0 && inc == -1 {
			a = 0
			inc = 1
		}
	}

	strs := make([]string, k)
	for i := 0; i < k; i++ {
		strs[i] = string(rows[i])
	}

	return strs
}
