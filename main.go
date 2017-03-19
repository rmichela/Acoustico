package main

import (
	"os"
	"time"

	"github.com/gordonklaus/portaudio"
	"github.com/mjibson/go-dsp/wav"
)

const sampleRate = Frequency(44100)

var phasor TrigOscillator
var phasor2 TrigOscillator

var helloSample *Sample

func main() {

	phasor = NewOscillator(sampleRate)
	phasor2 = NewOscillator(sampleRate)

	file, error := os.Open("samples/hello.wav")
	chk(error)
	wav, error := wav.New(file)
	chk(error)
	helloSample, error = NewSample(wav)

	portaudio.Initialize()
	defer portaudio.Terminate()

	//sampler := monoSample(phasorFunc)
	sampler := stereoSample(helloSample.AsLoopingTFunc(), helloSample.AsLoopingTFunc())
	channels := 2
	stream, err := portaudio.OpenDefaultStream(0, channels, float64(sampleRate), 0, sampler)
	chk(err)

	defer stream.Close()
	chk(stream.Start())
	time.Sleep(4 * time.Second)
	chk(stream.Stop())
}

func phasorFunc(t Timecode) Amplitude {
	df := TMap(G2T(phasor2.Sine, FreqFunc(Hz(0.5))),
		func(a Amplitude) Amplitude {
			return 500 + a*200
		})

	return Amplitude(G2T(phasor.Sine, df)(t))
}

func monoSample(tfunc TFunc) func([][]float32) {
	var t uint64
	return func(out [][]float32) {
		for i := range out[0] {
			var tc = Timecode(t + uint64(i))
			out[0][i] = float32(tfunc(tc))
		}
		t = t + uint64(len(out[0]))
	}
}

func stereoSample(left TFunc, right TFunc) func([][]float32) {
	var t uint64
	return func(out [][]float32) {
		for i := range out[0] {
			var tc = Timecode(t + uint64(i))
			out[0][i] = float32(left(tc))
			out[1][i] = float32(right(tc))
		}
		t = t + uint64(len(out[0]))
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
