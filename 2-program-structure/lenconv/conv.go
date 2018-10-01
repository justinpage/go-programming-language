package lenconv

func FToM(c Feet) Meter { return Meter(c * 0.3048) }
func MToF(m Meter) Feet { return Feet(m / 0.3048) }
