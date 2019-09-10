// Package tempconv performs Celsius and Fahrenheit conversions.
package unitconv

import "fmt"

type Celsius float64
type Fahrenheit float64

type Foot float64
type Meter float64

type Pound float64
type Kilogram float64

func (c Celsius) String() string    { return fmt.Sprintf("%.3g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.3g°F", f) }

func (f Foot) String() string  { return fmt.Sprintf("%.3g ft", f) }
func (m Meter) String() string { return fmt.Sprintf("%.3g m", m) }

func (p Pound) String() string    { return fmt.Sprintf("%.3g lb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%.3g kg", k) }
