package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dx)
	for i := 0; i < dx; i++ {
		ret[i] = make([]uint8, dy)
		for j := 0; j < dx; j++ {
			ret[i][j] = uint8(i*2 + j + (i+j)/2)
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
