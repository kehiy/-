package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	inputData, err := ioutil.ReadFile("./words.txt")
	if err != nil {
		log.Fatal("error reading data:", err)
	}

	var buf bytes.Buffer

	gz := gzip.NewWriter(&buf)

	if _, err := gz.Write(inputData); err != nil {
		panic(err)
	}

	if err := gz.Close(); err != nil {
		panic(err)
	}

	inputSize := len(inputData)
	compressedSize := len(buf.Bytes())

	fmt.Printf("Input Size: %d bytes\n", inputSize)
	fmt.Printf("Compressed Size: %d bytes\n", compressedSize)
	/*
	output:(using words.txt):
	!Input Size: 4862992 bytes
	!Compressed Size: 1474416 bytes
	*/
}
