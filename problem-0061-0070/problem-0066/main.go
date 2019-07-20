package main

func main() {

}

func tossUnbiased(tossBiased func() int) int {
	a, b := tossBiased(), tossBiased()

	if a == b {
		return 0
	}

	return 1
}
