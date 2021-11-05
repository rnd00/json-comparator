package main

import (
	"errors"
	"fmt"
	"os"
)

func init() {
	checkData("data/1.json")
	checkData("data/2.json")
}

func main() {

	var d, f Data
	errorHandler(d.LoadRaw("data/1.json"))
	errorHandler(f.LoadRaw("data/2.json"))

	errorHandler(d.Unmarshal())
	errorHandler(f.Unmarshal())

	r, e := d.Compare(&f)
	errorHandler(e)

	fmt.Println(r)
	d.Diff(&f)
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func checkData(path string) {
	stat, err := os.Stat(path)
	errorHandler(err)

	if stat.IsDir() {
		errorHandler(errors.New(fmt.Sprintf("Error, specified file %s is a dir\n", stat.Name())))
	}

	if stat.Size() < 1 {
		errorHandler(errors.New(fmt.Sprintf("Error while checking %s filesize\n", stat.Name())))
	}
}
