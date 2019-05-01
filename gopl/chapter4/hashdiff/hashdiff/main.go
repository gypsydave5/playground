package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/gypsydave5/playground/gopl/chapter4/hashDiff"
)

func main() {
	log.SetPrefix("hashdiff: ")
	log.SetFlags(0)

	if len(os.Args) != 3 {
		fmt.Println("usage: hashDiff [hash1] [hash2]")
		os.Exit(1)
	}

	h1, err := decode256byte(os.Args[1])
	if err != nil {
		log.Fatalf(err.Error())
	}

	h2, err := decode256byte(os.Args[2])
	if err != nil {
		log.Fatalf(err.Error())
	}

	diff := hashdiff.HashDiff(h1, h2)

	fmt.Printf("%d\n", diff)
}

func decode256byte(s string) ([32]byte, error) {
	result := [32]byte{}

	if len(s) != 64 {
		return result, fmt.Errorf("%v is not a hexadecimal representation of a 256 bit hash", s)
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		return result, err
	}

	copy(result[:], b)

	return result, err
}
