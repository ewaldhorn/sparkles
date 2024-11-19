package main

import (
	"math"
)

var (
	targetRed   uint8   = 235
	targetGreen uint8   = 235
	targetBlue  uint8   = 235
	sensitivity float64 = 100
)

// ----------------------------------------------------------------------------
func processVideoData(data []uint8) []uint8 {
	tmpData := make([]uint8, len(data))

	for y := 0; y < outputCanvasHeight; y++ {
		for x := outputCanvasWidth - 1; x >= 0; x-- {
			offset := (y*outputCanvasWidth + x) * 4

			red := data[offset]
			green := data[offset+1]
			blue := data[offset+2]
			alpha := data[offset+3]

			distance := math.Sqrt(math.Pow(float64(targetRed-red), 2) + math.Pow(float64(targetGreen-green), 2) + math.Pow(float64(targetBlue-blue), 2))

			if distance < sensitivity {
				red = 20
				green = 20
				blue = 20
				alpha = 20
			}

			tmpData[offset] = red
			tmpData[offset+1] = green
			tmpData[offset+2] = blue
			tmpData[offset+3] = alpha
		}
	}

	return tmpData
}
