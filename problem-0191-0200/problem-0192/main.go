package main

func main() {

}

func advanceToEnd(steps []int) bool {
	length := len(steps)
	current := steps[0]
	if length == 1 {
		return true
	}

	if current >= length {
		return true
	}

	for i := 1; i <= current && i < length; i++ {
		if advanceToEnd(steps[i:]) {
			return true
		}
	}

	return false
}
