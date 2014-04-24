package main

import (
	"fmt"
	"reflect"
)

type CustomType int

const (
	FOO_BAR CustomType = 3
)

func main() {
	a := FOO_BAR
	v := reflect.ValueOf(a)
	fmt.Println(v.Elem())
}
