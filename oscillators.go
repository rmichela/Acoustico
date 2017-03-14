package main

import (
	"math"
	"math/cmplx"
)

const twoPi = 2 * math.Pi

// number of elapsed samples before stabilization occurs
const stabilizationPeriod = 500

// SineOscillator generates Sine waves
type SineOscillator interface {
	Sine(Timecode, Frequency) Amplitude
}

// CosineOscillator generates Cosine waves
type CosineOscillator interface {
	Cosine(Timecode, Frequency) Amplitude
}

// TrigOscillator generates Sine and Cosine waves
type TrigOscillator interface {
	SineOscillator
	CosineOscillator
}

// SawtoothOscillator generates Sawtooth waves
type SawtoothOscillator interface {
	Sawtooth(Timecode, Frequency) Amplitude
}

// SquareOscillator generates Square waves
type SquareOscillator interface {
	Square(Timecode, Frequency) Amplitude
}

// TriangleOscillator generates Triangle waves
type TriangleOscillator interface {
	Triangle(Timecode, Frequency) Amplitude
}

// Oscillator implements periodic waveform Oscillators using a recursively defined
// complex phasor.
// - All waveforms are in the range [-1, 1]
// - All waveforms (except cosine) are in phase with sine
// -- https://en.wikipedia.org/wiki/File:Waveforms.svg
type Oscillator struct {
	// The last known Oscillator timecode
	t Timecode
	// Time of last stabilization
	s Timecode
	// A phasor holding the current state of the trig Oscillator
	z complex128
	// The sample rate for generation
	sampleRate Frequency
}

// NewOscillator creates a new Oscillator
func NewOscillator(f Frequency) *Oscillator {
	gen := new(Oscillator)
	gen.t = 0
	gen.s = 0
	gen.sampleRate = f
	gen.z = complex(1, 0)
	return gen
}

// Advances the phasor counterclockwise according the desired frequency with
// respect to a change in time.
// http://dsp.stackexchange.com/a/1087
func (gen *Oscillator) advance(tʹ Timecode, f Frequency) {
	// determine how many samples we have to advance
	// typically this is one - more than one implies time dialation
	Δt := tʹ - gen.t

	// compute the angular frequency of the oscilator in radians
	ω := twoPi * f / sampleRate

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
}

// Sine generates a sine wave
func (gen *Oscillator) Sine(tʹ Timecode, f Frequency) Amplitude {
	// advance the phasor
	gen.advance(tʹ, f)
	// return the 'sine' component of the phasor
	return Amplitude(imag(gen.z))
}

// Cosine generates a sine wave
func (gen *Oscillator) Cosine(tʹ Timecode, f Frequency) Amplitude {
	// advance the phasor
	gen.advance(tʹ, f)
	// return the 'cosine' component of the phasor
	return Amplitude(real(gen.z))
}

// Sawtooth generates a sawtooth wave
func (gen *Oscillator) Sawtooth(tʹ Timecode, f Frequency) Amplitude {
	// advance the phasor
	gen.advance(tʹ, f)
	// return the 'phase' of the phasor, normalized between [-1, 1]
	return Amplitude(cmplx.Phase(gen.z) / math.Pi)
}

// Square generates a square wave
func (gen *Oscillator) Square(tʹ Timecode, f Frequency) Amplitude {
	// advance the phasor
	gen.advance(tʹ, f)
	// return the 'imaginary sign' component of the phasor
	if math.Signbit(imag(gen.z)) {
		return 1.0
	}
	return -1.0
}

// Triangle generates a sawtooth wave
func (gen *Oscillator) Triangle(tʹ Timecode, f Frequency) Amplitude {
	// advance the phasor
	gen.advance(tʹ, f)
	// return |Saw|, normalized between [-1, 1]
	// 'rotate' by 90 degrees to align phase with Sine
	saw := (cmplx.Phase(gen.z * complex(0, 1))) / math.Pi
	return Amplitude(2*math.Abs(saw) - 1)
}
