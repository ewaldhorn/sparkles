package main

import "syscall/js"

var (
	canProcess      bool = false
	camera          js.Value
	window          js.Value
	outputCanvas    js.Value
	outputContext2d js.Value
)

// ----------------------------------------------------------------------------
func init() {
	initWebcam()
	initCanvas()

	window = js.Global().Get("window")
	camera = setupCameraFeed()
}

// ----------------------------------------------------------------------------
func mainLoop(this js.Value, args []js.Value) interface{} {
	window.Call("requestAnimationFrame", js.FuncOf(mainLoop))
	drawImageToContext(camera)
	return nil
}

// ----------------------------------------------------------------------------
func main() {
	js.Global().Get("console").Call("log", "Sparkles - Started")

	mainLoop(js.ValueOf(nil), make([]js.Value, 0))

	select {}
}
