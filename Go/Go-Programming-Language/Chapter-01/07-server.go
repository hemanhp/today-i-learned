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
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/counter", counter)
	http.HandleFunc("/generator", generator)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generator(writer http.ResponseWriter, request *http.Request) {
	lissajouss(writer)
}

func counter(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Fprintln(writer, count)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(writer, "Hello World!")
}

func lissajouss(out io.Writer) {

	var pallette = []color.Color{color.White, color.Black}

	const (
		whiteIndex = 0
		blackIndex = 1
	)
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
