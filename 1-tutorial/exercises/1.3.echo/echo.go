package echo

import (
	"fmt"
	"strings"
)

func Join(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func Range(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func Loop(args []string) {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}
