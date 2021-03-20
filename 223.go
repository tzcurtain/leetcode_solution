package main

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	x1, y1 := max(A, E), max(B, F)
	x2, y2 := min(C, G), min(D, H)
	s1, s2 := calcArea(A, C, B, D), calcArea(E, G, F, H)
	if x1 >= x2 || y1 >= y2 {
		return s1 + s2
	}
	return s1 + s2 - calcArea(x1, x2, y1, y2)
}

func calcArea(x1, x2, y1, y2 int) int {
	return (x2 - x1) * (y2 - y1)
}
