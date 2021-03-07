package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

func getPallitre() []color.Color {
	rand.Seed(time.Now().UTC().UnixNano())
	return []color.Color{
		color.RGBA{
			uint8(rand.Intn(255-0) + 0),
			uint8(rand.Intn(255-0) + 0),
			uint8(rand.Intn(255-0) + 0),
			255,
		},
		color.RGBA{
			uint8(rand.Intn(255-0) + 0),
			uint8(rand.Intn(255-0) + 0),
			uint8(rand.Intn(255-0) + 0),
			255,
		}}
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	pallete := getPallitre()

	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64()
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func main() {
	file, _ := os.Create("./data/lissajous.gif")
	lissajous(file)
	file.Close()
}
