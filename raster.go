package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"

	"github.com/mjibson/go-dsp/spectral"
	"github.com/wayneashleyberry/terminal-dimensions"
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
	termWidth := termWidth()

	segmentLength := termWidth*2 - 2 // FFT results in segmentLength/2+1 output samples

	// stores the rendered output
	lineQueueDepth := 8
	lineQueue := make(chan []float64, lineQueueDepth)
	go printer(lineQueue, segmentLength, colormap)

	options := newOptions(segmentLength)

	return func(out [][]float64) {
		// adjust PSD options if they are too agressive for the input data
		if segmentLength > len(out[0]) {
			options = newOptions(len(out[0]))
		}

		// compute raster values from the inner rasterizer
		inner(out)

		// divide the rasterized sound into [segmentLength] sample segments
		var segments [][]float64
		for _, chanBuf := range out {
			chanSegs := spectral.Segment(chanBuf, options.NFFT, options.Noverlap)
			segments = append(segments, chanSegs...)
		}

		// subtract the mean of each segment from each segment's sample
		// https://www.mathworks.com/matlabcentral/answers/267658-why-am-i-getting-huge-values-on-low-frequencies-of-my-psd#comment_342452
		for i := range segments {
			segment := segments[i]
			var avg float64
			for j := range segment {
				avg += segment[j]
			}
			avg /= float64(len(segment))
			for j := range segment {
				segment[j] -= avg
			}
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

		// normalize the PSD
		for i := range psd {
			// divide by number of segments to get average
			psd[i] /= float64(len(segments))
			// linearize the PSD
			psd[i] *= math.Sqrt(freqs[i])
		}

		// skip line rendering if there are too many lines waiting to be rendered
		if len(lineQueue) < lineQueueDepth {
			lineQueue <- psd
		}
	}
}

func newOptions(segmentLength int) *spectral.PwelchOptions {
	overlap := segmentLength / 2
	options := new(spectral.PwelchOptions)
	options.Scale_off = true
	options.Noverlap = overlap
	options.NFFT = segmentLength
	options.Pad = segmentLength
	return options
}

// Renders printed lines from a queue
func printer(lineQueue chan []float64, segmentLength int, colormap Colormap) {
	outLine := bytes.NewBuffer(make([]byte, 30*segmentLength, 30*segmentLength))
	for {
		psd := <-lineQueue
		outLine.Reset()
		for i := range psd {
			r, g, b := colormap(psd[i])
			outLine.WriteString("\x1b[48;2;")
			outLine.WriteString(strconv.Itoa(int(r)))
			outLine.WriteString(";")
			outLine.WriteString(strconv.Itoa(int(g)))
			outLine.WriteString(";")
			outLine.WriteString(strconv.Itoa(int(b)))
			outLine.WriteString("m \x1b[0m")
		}
		fmt.Println(outLine.String())
	}
}

// get the width of the terminal - always even
func termWidth() int {
	w, e := terminaldimensions.Width()
	if e != nil {
		return 80
	}

	// make w even
	if w%2 != 0 {
		w--
	}
	return int(w)
}
