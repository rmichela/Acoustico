package main

import (
	"math"
	"time"

	"github.com/gordonklaus/portaudio"
)

const sampleRate = 44100

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()
	// s := newStereoSine(256, 256, sampleRate)
	var stream *portaudio.Stream
	var err error
	stream, err = portaudio.OpenDefaultStream(0, 1, sampleRate, 0, processAudio)
	chk(err)

	defer stream.Close()
	chk(stream.Start())
	time.Sleep(2 * time.Second)
	chk(stream.Stop())
}

type timecode uint64
type frequency float64
type amplitude float64

// Time function
type tFunc func(timecode) amplitude

// Generative function
type gFunc func(timecode, frequency) amplitude

func g2t(g gFunc, f frequency) tFunc {
	return func(t timecode) amplitude {
		return g(t, f)
	}
}

var t uint64

func processAudio(out [][]float32) {

	// println(sine(t, 256))

	for i := range out[0] {
		var tc = timecode(t + uint64(i))

		// var df = frequency(10 * sine(tc, 2))
		// var da = square(tc, 2)*0.5 + 1

		out[0][i] = float32(g2t(sine, 220)(tc))

	}

	t = t + uint64(len(out[0]))
}

var twoPi = 2 * math.Pi

func sine(t timecode, f frequency) amplitude {
	var step = f / sampleRate
	return amplitude(math.Sin(twoPi * float64(t) * float64(step)))
}

func saw(t timecode, f frequency) amplitude {
	var step = f / sampleRate
	var x = float64(t) * float64(step)
	return amplitude(2*(x-math.Floor(x)) - 1)
}

func square(t timecode, f frequency) amplitude {
	if math.Signbit(float64(sine(t, f))) {
		return amplitude(1)
	}
	return amplitude(-1)
}

func triangle(t timecode, f frequency) amplitude {
	return amplitude(2*math.Abs(float64(saw(t, f))) - 1)
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
