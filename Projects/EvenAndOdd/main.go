package main

func main() {
	ints := []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, nu := range ints {
		if nu%2 == 0 {
			println(nu, "is Even")
		} else {
			println(nu, "is Odd")
		}
	}
}
