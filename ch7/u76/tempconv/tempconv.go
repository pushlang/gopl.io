// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import (
	"flag"
	"fmt"
	t "gopl.io/ch2/u21/tempconv"
)

type celsiusFlag struct{ t.Celsius }
type fahrenheitFlag struct{ t.Fahrenheit }
type kelvinFlag struct{ t.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = t.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = t.FToC(t.Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = t.KToC(t.Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value t.Celsius, usage string) *t.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
