package main

import (
	"image"
	"image/color"
	"math/cmplx"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"math"
)

func mandelbrot(c complex128) color.Color {
	const iter = 100
	z := complex(0, 0)
	var n uint8
	for n = 0; n < iter; n++ {
			z = z*z + c
			if cmplx.Abs(z) > 4 {
					break
			}
	}

	if n == iter {
			return color.Black
	}

	// Création d'un dégradé de couleur basé sur le nombre d'itérations
	red := uint8(128 + 127*math.Sin(float64(n)*math.Pi/16))
	green := uint8(128 + 127*math.Sin(float64(n)*math.Pi/8))
	blue := uint8(128 + 127*math.Sin(float64(n)*math.Pi/4))

	return color.RGBA{R: red, G: green, B: blue, A: 255}
}


func drawMandelbrot(img *image.RGBA, scale, xOffset, yOffset float64) {
	for x := 0; x < 800; x++ {
		for y := 0; y < 800; y++ {
			zx := (float64(x)-400)/scale + xOffset
			zy := (float64(y)-400)/scale + yOffset
			z := complex(zx, zy)
			img.Set(x, y, mandelbrot(z))
		}
	}
}

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Fractal Zoom")

	img := image.NewRGBA(image.Rect(0, 0, 800, 800))
	canvasImage := canvas.NewImageFromImage(img)
	canvasImage.FillMode = canvas.ImageFillOriginal

	w.SetContent(canvasImage)
	w.Resize(fyne.NewSize(800, 800))

	scale := 200.0
	xOffset, yOffset := -0.0, 1.0

	go func() {
		for {
			drawMandelbrot(img, scale, xOffset, yOffset)
			canvasImage.Refresh()
			time.Sleep(100 * time.Millisecond) // Adjust the speed of zoom here
			scale *= 1.1 // Adjust the zoom factor here
		}
	}()

	w.ShowAndRun()
}
