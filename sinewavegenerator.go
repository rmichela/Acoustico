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
	s := newStereoSine(256, 256, sampleRate)
	defer s.Close()
	chk(s.Start())
	time.Sleep(2 * time.Second)
	chk(s.Stop())
}

type stereoSine struct {
	*portaudio.Stream
	stepL, phaseL float64
	stepR, phaseR float64
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

func newStereoSine(freqL, freqR, sampleRate float64) *stereoSine {
	s := &stereoSine{nil, freqL / sampleRate, 0, freqR / sampleRate, 0}
	var err error
	s.Stream, err = portaudio.OpenDefaultStream(0, 1, sampleRate, 0, s.processAudio)
	chk(err)
	return s
}

var t uint64

func (g *stereoSine) processAudio(out [][]float32) {

	// println(sine(t, 256))

	for i := range out[0] {
		// out[0][i] = float32(math.Sin(2 * math.Pi * g.stepL * float64(t+i)))
		var tc = timecode(t + uint64(i))

		// out[0][i] = float32(triangle(tc, 220))
		out[0][i] = float32(g2t(saw, 220)(tc))
		// out[1][i] = float32(sine(t+i, 500))
		// _, g.phaseL = math.Modf(g.phaseL + g.stepL)
		// out[1][i] = float32(math.Sin(2 * math.Pi * g.phaseR))
		// _, g.phaseR = math.Modf(g.phaseR + g.stepR)
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
