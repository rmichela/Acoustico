package main

// Timecode is an index in time
type Timecode uint64

// Frequency is an oscilation in hertz
type Frequency float64

// Hz returns a Frequency in hertz
func Hz(f float64) Frequency {
	return Frequency(f)
}

// KHz returns a Frequency in kilohertz
func KHz(f float64) Frequency {
	return Frequency(1000 * f)
}

// Amplitude is the intensity of a value
type Amplitude float64

// TFunc provides a value at a given timecode
type TFunc func(Timecode) Amplitude

// GFunc generates impulses at a given frequency
type GFunc func(Timecode, Frequency) Amplitude

// ConstFunc returns a time invariant constant TFunc
func ConstFunc(c float64) TFunc {
	return func(t Timecode) Amplitude {
		return Amplitude(c)
	}
}

// FreqFunc returns a time invariant constant TFunc
func FreqFunc(c Frequency) TFunc {
	return func(t Timecode) Amplitude {
		return Amplitude(c)
	}
}

// TMap applies function g to every value of TFunc f
func TMap(f TFunc, g func(Amplitude) Amplitude) TFunc {
	return func(t Timecode) Amplitude {
		return (g(f(t)))
	}
}

// G2T converts a GFunc to a TFunc by currying frequency
func G2T(g GFunc, f TFunc) TFunc {
	return func(t Timecode) Amplitude {
		return g(t, Frequency(f(t)))
	}
}
