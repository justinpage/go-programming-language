package unitconv

func ClToFh(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FhToCl(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func FtToMt(f Foot) Meter { return Meter(f / 3.281) }
func MtToFt(m Meter) Foot { return Foot(m * 3.281) }

func LbToKg(p Pound) Kilogram { return Kilogram(p * 0.45359237) }
func KgToLb(k Kilogram) Pound { return Pound(k / 0.45359237) }
