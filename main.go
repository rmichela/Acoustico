package main

import (
	"math"
	"time"

	"github.com/gordonklaus/portaudio"
)

const sampleRate = Frequency(44100)

var phasor *TrigGenerator
var phasor2 *TrigGenerator

func main() {
	phasor = NewTrigGenerator(sampleRate)
	phasor2 = NewTrigGenerator(sampleRate)

	portaudio.Initialize()
	defer portaudio.Terminate()
	// s := newStereoSine(256, 256, sampleRate)
	var stream *portaudio.Stream
	var err error
	stream, err = portaudio.OpenDefaultStream(0, 1, float64(sampleRate), 0, processAudio)
	chk(err)

	defer stream.Close()
	chk(stream.Start())
	time.Sleep(10 * time.Second)
	chk(stream.Stop())
}

var t uint64

func processAudio(out [][]float32) {
	for i := range out[0] {
		var tc = Timecode(t + uint64(i))

		df := phasor2.Sine(tc, Frequency(2)) * 100

		out[0][i] = float32(G2T(phasor.Sine, 350+Frequency(df))(tc))
	}

	t = t + uint64(len(out[0]))
}

var twoPi = 2 * math.Pi

func sine(t Timecode, f Frequency) Amplitude {
	var step = f / sampleRate
	return Amplitude(math.Sin(twoPi * float64(t) * float64(step)))
}

func saw(t Timecode, f Frequency) Amplitude {
	var step = f / sampleRate
	var x = float64(t) * float64(step)
	return Amplitude(2*(x-math.Floor(x)) - 1)
}

func square(t Timecode, f Frequency) Amplitude {
	if math.Signbit(float64(sine(t, f))) {
		return Amplitude(1)
	}
	return Amplitude(-1)
}

func triangle(t Timecode, f Frequency) Amplitude {
	return Amplitude(2*math.Abs(float64(saw(t, f))) - 1)
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
