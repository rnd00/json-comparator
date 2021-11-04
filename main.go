package main

import (
	"fmt"
)

func init() {}

func main() {

	var d, f Data
	if e := d.LoadRaw("data/1.json"); e != nil {
		errorHandler(e)
	}
	if e := f.LoadRaw("data/2.json"); e != nil {
		errorHandler(e)
	}

	if e := d.Unmarshal(); e != nil {
		errorHandler(e)
	}
	if e := f.Unmarshal(); e != nil {
		errorHandler(e)
	}

	r, e := d.Compare(&f)
	if e != nil {
		errorHandler(e)
	}

	fmt.Println(r)
	d.Diff(&f)
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
