package main

import (
	"github.com/mjibson/go-dsp/wav"
)

// Sample represents PCM data as a TFunc
type Sample struct {
	wav    *wav.Wav
	buffer []float32
}

// NewSample creates a new Sample struct
func NewSample(wav *wav.Wav) (*Sample, error) {
	sample := new(Sample)
	sample.wav = wav

	buffer, error := wav.ReadFloats(wav.Samples)
	if error != nil {
		return nil, error
	}
	sample.buffer = buffer

	return sample, nil
}

// AsTFunc returns a TFunc expressing the underlying sample data.
// Zero is returned for timecodes beyond the length of the sample.
func (sample *Sample) AsTFunc() TFunc {
	return func(t Timecode) Amplitude {
		if int(t) >= sample.wav.Samples {
			return 0
		}
		return Amplitude(sample.buffer[t])
	}
}

// AsLoopingTFunc returns a TFunc expressing the underlying sample data in a loop.
func (sample *Sample) AsLoopingTFunc() TFunc {
	return func(t Timecode) Amplitude {
		tsample := uint64(t) % uint64(sample.wav.Samples)
		return Amplitude(sample.buffer[tsample])
	}
}
