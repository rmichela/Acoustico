package main

import (
	"math"
	"math/cmplx"
)

// number of elapsed samples before stabilization occurs
const stabilizationPeriod = 500

// TrigGenerator implements a trigonometric generator using a recursively updated complex
// phasor
type TrigGenerator struct {
	// The last known generator timecode
	t Timecode
	// Time of last stabilization
	s Timecode
	// A phasor holding the current state of the trig generator
	z complex128
	// The sample rate for generation
	sampleRate Frequency
}

// NewTrigGenerator creates a new TrigGenerator
func NewTrigGenerator(f Frequency) *TrigGenerator {
	gen := new(TrigGenerator)
	gen.t = 0
	gen.s = 0
	gen.sampleRate = f
	gen.z = complex(1, 0)
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
	for i := Δt; i > 0; i-- {
		gen.z = gen.z * Ω
	}

	// stabilize the phasor's amplitude every once in a while
	// the amplitude can drift due to rounding errors
	// since z is a unity phasor, adjust its amplitude back towards unity
	if gen.s > stabilizationPeriod {
		a := real(gen.z)
		b := imag(gen.z)
		c := (3 - math.Pow(a, 2) - math.Pow(b, 2)) / 2
		gen.z = gen.z * complex(c, 0)
		gen.s = 0
	}

	// advance time
	gen.t += Δt
	gen.s += Δt
	// return the 'sine' component of the phasor
	return Amplitude(imag(gen.z))
}
