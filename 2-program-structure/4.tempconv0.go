package main

import "fmt"

const boilingF = 212.0

type Celsius float64
type Fahrenheit float64

const (
	AboluteZeroC Celsius = -273.15
	FreezingC    Celsius = 0
	BoilingC     Celsius = 100
)

func main() {
	fmt.Println(BoilingC)
}

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func fToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
