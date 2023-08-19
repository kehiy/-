package main

import (
	"fmt"
	"io/ioutil"
)

func kash(data string) string {
	symbols := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_-+={}[]/"
	hash := 0
	for i := 0; i < len(data); i++ {
		hash = hash + int(data[i])
	}
	result := ""
	for hash > 0 {
		idx := hash % len(symbols)
		result = string(symbols[idx]) + result
		hash = hash / len(symbols)
	}
	return result
}

func main() {
	fileBytes, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	fmt.Println(kash(fileString))
}
