package main

import (
	"fmt"
	"math"

	"github.com/mjibson/go-dsp/spectral"
)

// Rasterizer fills one or more buffers with discreet audio samples
type Rasterizer func([][]float64)

// RenderMono rasterizes a TFunc into a mono Portaudio channel
func RenderMono(tfunc TFunc) Rasterizer {
	var t uint64
	return func(out [][]float64) {
		for i := range out[0] {
			var tc = Timecode(t + uint64(i))
			out[0][i] = float64(tfunc(tc))
		}
		t = t + uint64(len(out[0]))
	}
}

// RenderStereo rasterizes two TFuncs into a stereo Portaudio channel
func RenderStereo(left TFunc, right TFunc) Rasterizer {
	var t uint64
	return func(out [][]float64) {
		for i := range out[0] {
			var tc = Timecode(t + uint64(i))
			out[0][i] = float64(left(tc))
			out[1][i] = float64(right(tc))
		}
		t = t + uint64(len(out[0]))
	}
}

// Visualize renders a Rasterizer's power spectral density to the screen
func Visualize(sampleRate Frequency, colormap Colormap, inner Rasterizer) Rasterizer {
	options := new(spectral.PwelchOptions)
	options.Scale_off = true
	options.Pad = 256

	var max float64

	return func(out [][]float64) {
		// compute raster values from the inner rasterizer
		inner(out)

		// divide the rasterized sound into 256 sample segments
		var segments [][]float64
		for _, chanBuf := range out {
			chanSegs := spectral.Segment(chanBuf, 256, 0)
			segments = append(segments, chanSegs...)
		}

		// compute the averge PSD for all segments
		psd := make([]float64, options.Pad/2+1) // output of FFT is len(time domain)/2+1
		for _, seg := range segments {
			// compute the PSD for each segment
			pxx, _ := spectral.Pwelch(seg, float64(sampleRate), options)
			for i := range pxx {
				psd[i] += pxx[i]
			}
			// println()
		}
		// divide by number of segments to get average, then print
		for i := range psd {
			psd[i] /= float64(len(segments))

			max = math.Max(max, psd[i])

			printColor(psd[i]/max, colormap)
		}
		println()
	}
}

func printColor(value float64, colormap Colormap) {
	r, g, b := colormap(value)
	fmt.Printf("\x1b[48;2;%d;%d;%dm \x1b[0m", r, g, b)
}
