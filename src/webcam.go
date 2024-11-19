package main

import (
	"syscall/js"
)

// ----------------------------------------------------------------------------
var (
	navigator         js.Value
	videoElement      js.Value
	awaitVideoChannel chan interface{}
)

// ----------------------------------------------------------------------------
func initWebcam() {
	navigator = js.Global().Get("navigator").Get("mediaDevices")
	videoElement = js.Global().Get("document").Call("createElement", "video")
}

// ----------------------------------------------------------------------------
func setupCameraFeed() js.Value {
	user_media_params := map[string]interface{}{
		"video": true,
	}

	awaitVideoChannel = make(chan interface{})

	navigator.Call(
		"getUserMedia",
		user_media_params,
	).Call(
		"then",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			videoElement.Set("srcObject", args[0]) // media stream is active on args[0]
			videoElement.Call("addEventListener", "canplaythrough", js.FuncOf(startVideoFeed))
			close(awaitVideoChannel)

			return nil
		}),
		js.FuncOf(err),
	)
	<-awaitVideoChannel

	return videoElement
}

// ----------------------------------------------------------------------------
func err(this js.Value, args []js.Value) interface{} {
	js.Global().Get("console").Call("log", "error:", this, args)

	return nil
}

// ----------------------------------------------------------------------------
func startVideoFeed(this js.Value, args []js.Value) interface{} {
	videoElement.Call("play")

	outputCanvasWidth = videoElement.Get("videoWidth").Int()
	outputCanvasHeight = videoElement.Get("videoHeight").Int()

	imageData = outputContext2d.Call("createImageData", outputCanvasWidth, outputCanvasHeight)

	outputCanvas.Set("width", js.ValueOf(outputCanvasWidth))
	outputCanvas.Set("height", js.ValueOf(outputCanvasHeight))

	// mirror image
	outputContext2d.Call("translate", outputCanvasWidth, 0)
	outputContext2d.Call("scale", -1, 1)

	canProcess = true

	return nil
}
