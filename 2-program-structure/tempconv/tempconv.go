// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

const boilingF = 212.0

type Celsius float64
type Fahrenheit float64
type Kelvin float64

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
