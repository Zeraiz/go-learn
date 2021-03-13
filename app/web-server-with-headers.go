package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Method: %q, Url: %q, Proto %q\n", r.Method, r.URL, r.Proto)

	headers := make([]string, 0, len(r.Header))
	for k, _ := range r.Header {
		headers = append(headers, k)
	}

	sort.Strings(headers)

	for _, v := range headers {
		_, _ = fmt.Fprintf(w, "Header[%q] = %q\n", v, r.Header[v])
	}
	fmt.Fprintf(w, "Host: %q\n", r.Host)
	fmt.Fprintf(w, "Remote addrs: %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form field [%q] = %q\n", k, v)
	}
}

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

func lissajous(out io.Writer, setCycles int) {
	var (
		cycles  = 5.0
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	if setCycles > 0 {
		cycles = float64(setCycles)
	}
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
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), 1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func lassangeHandler(w http.ResponseWriter, r *http.Request) {
	cycles, _ := strconv.Atoi(r.FormValue("cycles"))
	lissajous(w, cycles)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/lassange", lassangeHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil))
}
