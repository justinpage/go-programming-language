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
		hash := flag.Int("hash", 256, "hash algorithm")

		flag.Parse()

		message := flag.Args()[0]

		switch *hash {
		case 384:
			sha384 := sha512.Sum384([]byte(message))
			fmt.Printf("SHA384: %x\n", sha384)
		case 512:
			sha512 := sha512.Sum512([]byte(message))
			fmt.Printf("SHA512: %x\n", sha512)
		default:
			sha256 := sha256.Sum256([]byte(message))
			fmt.Printf("SHA256: %x\n", sha256)
		}
	}
}
