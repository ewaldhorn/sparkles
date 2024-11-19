package main

import (
	"syscall/js"
)

var (
	outputCanvasWidth  = 0
	outputCanvasHeight = 0
	imageData          js.Value
)

// ----------------------------------------------------------------------------
func initCanvas() {
	outputCanvas = js.Global().Get("document").Call("getElementById", "mainCanvas")

	options := js.Global().Get("Object").New()
	options.Set("alpha", false)
	options.Set("willReadFrequently", true)

	outputContext2d = outputCanvas.Call("getContext", "2d", options)
}

// ----------------------------------------------------------------------------
func drawImageToContext(video js.Value) {
	if canProcess {
		outputContext2d.Call("drawImage", video, 0, 0, outputCanvasWidth, outputCanvasHeight)

		tmpData := processVideoData(getImageData())
		js.CopyBytesToJS(imageData.Get("data"), tmpData)

		outputContext2d.Call("putImageData", imageData, 0, 0)
	}
}

// ----------------------------------------------------------------------------
func getImageData() []uint8 {
	data := outputContext2d.Call("getImageData", 0, 0, outputCanvasWidth, outputCanvasHeight).Get("data")

	bufferSize := data.Get("length").Int()
	goData := make([]uint8, bufferSize)
	js.CopyBytesToGo(goData, data)

	return goData
}
