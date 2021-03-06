package main

import (
	"os"
	"time"

	"github.com/gordonklaus/portaudio"
	"github.com/mjibson/go-dsp/wav"
)

const sampleRate = Frequency(44100)
const channels = 1
const framesPerBuffer = 512

var phasor TrigOscillator
var phasor2 TrigOscillator

var sample *Sample

func main() {

	phasor = NewOscillator(sampleRate)
	phasor2 = NewOscillator(sampleRate)

	file, error := os.Open("samples/anewbeginning.wav")
	chk(error)
	wav, error := wav.New(file)
	chk(error)
	sample, error = NewSample(wav)

	portaudio.Initialize()
	defer portaudio.Terminate()

	// sampler := Visualize(sampleRate, Inferno, RenderStereo(
	// 	G2T(phasor.Sine, FreqFunc(Hz(20))),
	// 	G2T(phasor2.Sine, FreqFunc(KHz(20)))))
	sampler := Visualize(sampleRate, Inferno, RenderMono(sample.AsTFunc()))
	// sampler := Visualize(sampleRate, Inferno, RenderStereo(helloSample.AsLoopingTFunc(), helloSample.AsLoopingTFunc()))

	stream, err := portaudio.OpenDefaultStream(0, channels, float64(sampleRate), framesPerBuffer, downsample(sampler))
	chk(err)

	defer stream.Close()
	chk(stream.Start())
	time.Sleep(120 * time.Second)
	chk(stream.Stop())
}

func downsample(inner Rasterizer) func([][]float32) {
	// allocate a float64 buffer once to rasterize into
	out64 := make([][]float64, channels)
	for i := range out64 {
		out64[i] = make([]float64, framesPerBuffer)
	}

	return func(out32 [][]float32) {
		inner(out64)
		// downsample each float64 into a float32 for compatibility with portaudio
		for i := range out32 {
			for j := range out32[i] {
				out32[i][j] = float32(out64[i][j])
			}
		}
	}
}

func phasorFunc(t Timecode) Amplitude {
	df := TMap(G2T(phasor2.Sine, FreqFunc(Hz(0.5))),
		func(a Amplitude) Amplitude {
			return 10000 + a*9000
		})

	return Amplitude(G2T(phasor.Sine, df)(t))
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
