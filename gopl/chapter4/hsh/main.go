package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

var algorithm = flag.String("a", "SHA256", "Hashing `algorithm` to use. Either SHA256, SHA384 or SHA512")

func main() {
	flag.Parse()
	var b bytes.Buffer
	b.ReadFrom(os.Stdin)
	switch *algorithm {
	case "SHA256":
		fmt.Printf("%x\n", sha256.Sum256(b.Bytes()))
	case "SHA512":
		fmt.Printf("%x\n", sha512.Sum512(b.Bytes()))
	case "SHA384":
		fmt.Printf("%x\n", sha512.Sum384(b.Bytes()))
	default:
		log.Printf("Invalid algorithm %#v\n", *algorithm)
		flag.PrintDefaults()
		os.Exit(1)
	}
	os.Exit(0)
}
