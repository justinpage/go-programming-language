package display

import (
	"fmt"
	"reflect"
	"strconv"
)

const LIMIT = 25

func Display(name string, x interface{}) {
	var ctx int = 0
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x), ctx)
}

func display(path string, v reflect.Value, ctx int) {
	if ctx >= LIMIT {
		return
	}
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			ctx += 1
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), ctx)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			ctx += 1
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i), ctx)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			ctx += 1
			display(fmt.Sprintf("%s[%s]", path,
				formatKey(path, key, ctx)), v.MapIndex(key), ctx)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			ctx += 1
			display(fmt.Sprintf("(*%s)", path), v.Elem(), ctx)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			ctx += 1
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), ctx)
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

func formatKey(path string, v reflect.Value, ctx int) string {
	switch v.Kind() {
	case reflect.Struct, reflect.Array, reflect.Ptr:
		ctx += 1
		display(path, v, ctx)
		return v.Elem().Type().String()
	default:
		return formatAtom(v)
	}
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "Invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
		// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
