package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/ohler55/ojg/alt"
)

type Data struct {
	Raw    []byte
	JSON   interface{}
	Length int
	Valid  bool
}

func NewData() *Data {
	return &Data{}
}

func (d *Data) LoadRaw(path string) error {
	json, e := os.Open(path)
	if e != nil {
		return e
	}
	defer json.Close()

	bv, e := ioutil.ReadAll(json)
	if e != nil {
		return e
	}

	d.Valid = govalidator.IsJSON(string(bv))
	d.Length = len(bv)
	d.Raw = bv
	return nil
}

func (d *Data) Unmarshal() error {
	if a := d.IsValid(); a != nil {
		return a
	}

	var j interface{}
	e := json.Unmarshal(d.Raw, &j)
	if e != nil {
		return e
	}
	d.JSON = j

	return nil
}

func (d *Data) IsValid() error {
	if d.Valid != true {
		return errors.New("Not a valid JSON Type")
	}

	return nil
}

func (d *Data) DeepEqual(f *Data) bool {
	if !reflect.DeepEqual(d.JSON, f.JSON) {
		return false
	}
	return true
}

func (d *Data) Compare(f *Data) (string, error) {
	var result string

	// Check
	if !d.Valid || !f.Valid {
		return result, errors.New("One or both data are not valid")
	}

	if d.JSON == nil {
		return result, errors.New("(Source) - Need to Unmarshal first.")
	}
	if f.JSON == nil {
		return result, errors.New("(Comparator) - Need to Unmarshal first.")
	}

	// Compare
	result = "Comparison Result;\n"
	if d.Length != f.Length {
		l := fmt.Sprintf("---\nDifferent length between Source and Comparator.\nSource Length: %d\nComparator Length: %d\n---\n", d.Length, f.Length)
		result += l
	} else {
		l := fmt.Sprintf("---\nSame length between Source and Comparator.\nSource Length: %d\nComparator Length: %d\n---\n", d.Length, f.Length)
		result += l
	}
	de := fmt.Sprintf("reflect.DeepEqual(source, comparison) => %v\n---\n", d.DeepEqual(f))
	result += de

	return result, nil
}

func (d *Data) Diff(f *Data) {
	diffs := alt.Diff(d.JSON, f.JSON)
	sort.Slice(diffs, func(i, j int) bool {
		return 0 < strings.Compare(fmt.Sprintf("%v", diffs[j]), fmt.Sprintf("%v", diffs[i]))
	})

	fmt.Printf("diff: %+v\n", diffs)
}
