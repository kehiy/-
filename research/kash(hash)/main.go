package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var bitset = []byte{0x17, 0x27, 0x37, 0x47, 0x57, 0x67, 0x77}

func kash(data string) int {
	hash := 0

	for i := 0; i < len(data); i++ {
		hash = hash + int(data[i]) >> 2
	}

	resultArr := []byte{}

	for hash > 0 {
		idx := hash % len(bitset)
		resultArr = append(resultArr, bitset[idx])
		hash = hash / len(bitset)
	}

	result := 0
	for _, j := range resultArr {
		result = result + int(j) | 1
	}

	return result
}

func main() {
	fileBytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	fileString := string(fileBytes)
	fmt.Println(kash(fileString))
}
