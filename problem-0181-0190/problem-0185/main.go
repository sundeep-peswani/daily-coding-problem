package main

import "fmt"

func main() {

}

func intersectingAreas(a, b rectangle) int {
	wSpan := min(a.right(), b.right()) - max(a.left, b.left)
	hSpan := min(a.bottom(), b.bottom()) - max(a.top, b.top)

	return wSpan * hSpan
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type rectangle struct {
	top, left int
	width, height int
}

func newRect(top, left, width, height int) rectangle {
	return rectangle{top, left, width, height}
}

func (r rectangle) bottom() int {
	return r.top + r.height
}

func (r rectangle) right() int {
	return r.left + r.width
}

func (r rectangle) String() string {
	return fmt.Sprintf(
		"(%d, %d) (%d, %d)\n(%d, %d) (%d, %d)", 
		r.top, r.left,
		r.top, r.right(),
		r.bottom(), r.left,
		r.bottom(), r.right(),
	)
}


