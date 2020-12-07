package main

import "fmt"

func main() {
	var maxGridX uint = 20
	var maxGridY uint = 20

	Grid := make([][]uint, maxGridY+1)
	for i := range Grid {
		Grid[i] = make([]uint, maxGridX+1)
	}

	Grid[0][0] = 1
	for y := uint(0); y < maxGridY+1; y++ {
		for x := uint(0); x < maxGridX+1; x++ {
			if x < maxGridX {
				Grid[y][x+1] += Grid[y][x]
			}
			if y < maxGridY {
				Grid[y+1][x] += Grid[y][x]
			}
		}
	}
	fmt.Println(Grid[maxGridY][maxGridX])
}
