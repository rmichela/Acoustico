package main

import (
	"time"

	"github.com/gordonklaus/portaudio"
)

const sampleRate = Frequency(44100)

var phasor SineGenerator
var phasor2 TriangleGenerator

func main() {

	phasor = NewGenerator(sampleRate)
	phasor2 = NewGenerator(sampleRate)

	portaudio.Initialize()
	defer portaudio.Terminate()
	// s := newStereoSine(256, 256, sampleRate)
	var stream *portaudio.Stream
	var err error
	stream, err = portaudio.OpenDefaultStream(0, 1, float64(sampleRate), 0, processAudio)
	chk(err)

	defer stream.Close()
	chk(stream.Start())
	time.Sleep(4 * time.Second)
	chk(stream.Stop())
}

var t uint64

func processAudio(out [][]float32) {
	for i := range out[0] {
		var tc = Timecode(t + uint64(i))

		df := phasor2.Triangle(tc, Frequency(0.5)) * 200

		out[0][i] = float32(G2T(phasor.Sine, 350+Frequency(df))(tc))
	}

	t = t + uint64(len(out[0]))
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
