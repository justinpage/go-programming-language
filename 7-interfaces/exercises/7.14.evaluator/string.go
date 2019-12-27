package eval

import (
	"bytes"
	"fmt"
)

func (v Var) String() string {
	return fmt.Sprint(string(v))
}

func (l literal) String() string {
	isIntegral := func(l literal) bool {
		return l == literal(int(l))
	}
	switch isIntegral(l) {
	case true:
		return fmt.Sprintf("%d", int(l))
	default:
		return fmt.Sprintf("%.2f", float64(l))
	}
}

func (u unary) String() string {
	return fmt.Sprintf("%s%s", string(u.op), u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %s %s)", b.x, string(b.op), b.y)
}

func (c call) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%s", c.fn)
	buf.WriteByte('(')
	for i, arg := range c.args {
		fmt.Fprintf(&buf, "%s", arg)
		if i < len(c.args)-1 {
			fmt.Fprint(&buf, ", ")
		}
	}
	buf.WriteByte(')')
	return buf.String()
}

func (a array) String() string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, arg := range a.args {
		fmt.Fprintf(&buf, "%s", arg)
		if i < len(a.args)-1 {
			fmt.Fprint(&buf, ", ")
		}
	}
	buf.WriteByte(']')
	fmt.Fprintf(&buf, ".%s", a.fn)
	return buf.String()
}
