// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import (
	"flag"
	"fmt"
)

const boilingF = 212.0

type Celsius float64
type Fahrenheit float64
type Kelvin float64

type celsiusFlag struct{ Celsius }

const (
	AboluteZeroC Celsius = -273.15
	FreezingC    Celsius = 0
	BoilingC     Celsius = 100

	AboluteZeroK Kelvin  = 0
	FreezingK    Kelvin  = 273.15
	BoilingK     Celsius = 373.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
