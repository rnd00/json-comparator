package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func init() {}

func main() {
	json1 := loadFile("data/1.json")
	json2 := loadFile("data/2.json")

	fmt.Println(equal(json1, json2))
}

func loadFile(path string) []byte {
	json, e := os.Open(path)
	errorHandler(e)
	defer json.Close()

	bv, e := ioutil.ReadAll(json)
	errorHandler(e)

	return bv
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
