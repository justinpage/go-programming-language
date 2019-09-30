package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		sha := flag.Int("h", 256, "hash")
		flag.Parse()

		input := flag.Args()[0]

		var message string

		switch *sha {
		case 384:
			hash := sha512.Sum384([]byte(input))
			message = fmt.Sprintf("message: %s\nsha%d: %x\n", input, *sha, hash)
		case 512:
			hash := sha512.Sum512([]byte(input))
			message = fmt.Sprintf("message: %s\nsha%d: %x\n", input, *sha, hash)

		default:
			hash := sha256.Sum256([]byte(input))
			message = fmt.Sprintf("message: %s\nsha%d: %x\n", input, *sha, hash)
		}

		fmt.Fprintf(os.Stdout, message)
	}
}
