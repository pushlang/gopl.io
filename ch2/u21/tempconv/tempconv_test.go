package tempconv

import (
	"math"
	"testing"
)

func TestFToC(t *testing.T) {
	var tests = []struct {
		f    Fahrenheit
		want Celsius
	}{
		{0, -17.78},
		{100, 37.78},
	}

	for _, test := range tests {
		if got := r(float64(FToC(test.f)), 2); Celsius(got) != test.want {
			t.Errorf("FToC(%v) = %v, want %v", test.f, got, test.want)
			continue
		}
	}
}

func TestCToF(t *testing.T) {
	var tests = []struct {
		f    Celsius
		want Fahrenheit
	}{
		{-17.78, 0},
		{37.78, 100},
	}

	for _, test := range tests {
		if got := r(float64(CToF(test.f)), 2); Fahrenheit(got) != test.want {
			t.Errorf("FToC(%v) = %v, want %v", test.f, got, test.want)
			continue
		}
	}
}

func TestKToC(t *testing.T) {
	var tests = []struct {
		f    Kelvin
		want Celsius
	}{
		{0, 273.15},
		{-273.15, 0},
	}

	for _, test := range tests {
		if got := r(float64(KToC(test.f)), 2); Celsius(got) != test.want {
			t.Errorf("FToC(%v) = %v, want %v", test.f, got, test.want)
			continue
		}
	}
}

func TestCToK(t *testing.T) {
	var tests = []struct {
		f    Celsius
		want Kelvin
	}{
		{0, -273.15},
		{273.15, 0},
	}

	for _, test := range tests {
		if got := r(float64(CToK(test.f)), 2); Kelvin(got) != test.want {
			t.Errorf("FToC(%v) = %v, want %v", test.f, got, test.want)
			continue
		}
	}
}

func r(f float64, n float64) float64 {
	p := math.Pow(10, n)
	return math.Round(f*p) / p
}
