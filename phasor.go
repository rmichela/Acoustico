package main

import (
	"math"
	"math/cmplx"
)

// TrigGenerator implements a trigonometric generator using a recursively updated complex
// phasor
type TrigGenerator struct {
	t          Timecode
	phasor     complex128
	sampleRate Frequency
}

// NewTrigGenerator creates a new TrigGenerator
func NewTrigGenerator(f Frequency) *TrigGenerator {
	gen := new(TrigGenerator)
	gen.t = 0
	gen.sampleRate = f
	gen.phasor = complex(1, 0)
	return gen
}

// Sine emits a sine wave for a given frequency and timecode
// http://dsp.stackexchange.com/a/1087
func (gen *TrigGenerator) Sine(tʹ Timecode, f Frequency) Amplitude {
	// determine how many samples we have to advance
	// typically this is one - more than one implies time dialation
	Δt := tʹ - gen.t

	// compute the angular frequency of the oscilator in radians
	ω := 2 * math.Pi * f / sampleRate

	// compute the complex angular coeficient
	Ω := cmplx.Exp(complex(0, ω))

	// advance the phasor Δt units
	for ; Δt > 0; Δt-- {
		gen.phasor = gen.phasor * Ω
	}
	gen.t = tʹ

	// return the 'sine' component of the phasor
	return Amplitude(imag(gen.phasor))
}
