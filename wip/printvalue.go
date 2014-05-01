package main

import "fmt"

type Blah struct {
	Bar   int
	Baz   string
	Baz1  string
	Baz2  string
	Baz3  string
	Baz4  string
	Baz5  string
	Baz6  string
	Baz7  string
	Baz8  string
	Baz9  string
	Baz10 string
	Baz11 string
	Baz12 string
	Baz13 string
	Baz14 string
	Baz15 string
	Baz16 string
	Baz17 string
	Baz18 string
	Baz19 string
	Baz20 string
	Baz21 string
}

func main() {
	bozz := &Blah{
		Bar: 8,
		Baz: "seven",
	}
	fmt.Printf("%#v\n", bozz)
}
