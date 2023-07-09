package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func GzipFunc(data []byte) []byte {
	var buf bytes.Buffer

	gz := gzip.NewWriter(&buf)

	if _, err := gz.Write(data); err != nil {
		panic(err)
	}

	if err := gz.Close(); err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func UnGzipFunc(data []byte) []byte {
	buf := bytes.NewReader(data)

	gz, err := gzip.NewReader(buf)
	if err != nil {
		panic(err)
	}
	defer gz.Close()

	uncompressedData, err := ioutil.ReadAll(gz)
	if err != nil {
		panic(err)
	}

	return uncompressedData
}
