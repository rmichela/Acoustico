package main

import (
	"math"

	"bytes"

	"strconv"

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
	lowerAudible := Hz(20)
	upperAudible := KHz(20)

	segmentLength := 256

	// stores the rendered output
	outLine := bytes.NewBuffer(make([]byte, 30*segmentLength, 30*segmentLength))
	var r, g, b byte

	options := new(spectral.PwelchOptions)
	options.Scale_off = true
	options.Pad = segmentLength

	return func(out [][]float64) {
		// compute raster values from the inner rasterizer
		inner(out)

		// divide the rasterized sound into [segmentLength] sample segments
		var segments [][]float64
		for _, chanBuf := range out {
			chanSegs := spectral.Segment(chanBuf, segmentLength, 0)
			segments = append(segments, chanSegs...)
		}

		// compute the averge PSD for all segments
		psd := make([]float64, options.Pad/2+1) // output of FFT is len(time domain)/2+1
		var freqs []float64
		for _, seg := range segments {
			var pxx []float64
			// compute the PSD for each segment
			pxx, freqs = spectral.Pwelch(seg, float64(sampleRate), options)
			for i := range pxx {
				psd[i] += pxx[i]
			}
			// println()
		}

		// render samples to the output buffer
		outLine.Reset()
		for i := range psd {
			// only render samples in the audible range
			if float64(lowerAudible) < freqs[i] && freqs[i] < float64(upperAudible) {
				// divide by number of segments to get average
				p := psd[i] / float64(len(segments))
				// linearize the PSD
				p = p * math.Sqrt(freqs[i]) * 0.5

				// render a 'pixel'
				r, g, b = colormap(p)
				outLine.WriteString("\x1b[48;2;")
				outLine.WriteString(strconv.Itoa(int(r)))
				outLine.WriteString(";")
				outLine.WriteString(strconv.Itoa(int(g)))
				outLine.WriteString(";")
				outLine.WriteString(strconv.Itoa(int(b)))
				outLine.WriteString("m \x1b[0m")
			}
		}
		println(outLine.String())
	}
}

// func printColor(value float64, colormap Colormap) {
// 	r, g, b := colormap(value)
// 	fmt.Printf("\x1b[48;2;%d;%d;%dm \x1b[0m", r, g, b)
// }
