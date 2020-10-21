package main

import (
	f "fmt"
	_ "strings"
	"time"

	s "github.com/inancgumus/screen"
)

type field = []bool

func main() {

	// set dimension
	rowLength, colHeight := 50, 10
	posX := 0
	posY := 0
	vx, vy := 1, 1

	var ball string = "⚾"
	s.Clear()
	for {

		for y := 0; y < colHeight; y++ {

			// generate empty grid
			var grid = []field{}
			for i := 0; i < colHeight; i++ {
				var row field
				for j := 0; j < rowLength; j++ {
					row = append(row, false)
				}
				grid = append(grid, row)
			}
			//change position
			posX += vx
			posY += vy

			//reverse direction if position goes beyond dimension
			if posX <= 0 || posX >= rowLength-1 {
				vx *= -1
			}

			if posY <= 0 || posY >= colHeight-1 {
				vy *= -1
			}

			grid[posY][posX] = true
			s.MoveTopLeft()
			for _, col := range grid {
				for _, row := range col {
					if row {
						f.Printf("%v", ball)

					} else {
						f.Printf(" ")
					}
				}
				f.Printf("\n")
			}
			time.Sleep(time.Millisecond * 100)
		}
	}

	time.Sleep(time.Second)
}
