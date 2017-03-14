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

// G2T converts a GFunc to a TFunc by currying frequency
func G2T(g GFunc, f Frequency) TFunc {
	return func(t Timecode) Amplitude {
		return g(t, f)
	}
}
