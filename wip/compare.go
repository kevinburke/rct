package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	f, err := ioutil.ReadFile("/Users/kevin/Applications/Wineskin/rct2.app/Contents/Resources/drive_c/GOG Games/RollerCoaster Tycoon 2 Triple Thrill Pack/Tracks/mymischief.td6")
	g, err := ioutil.ReadFile("/Users/kevin/code/go/src/github.com/kevinburke/rct-rides/rides/mischief.td6")

	if err != nil {
		fmt.Errorf(err.Error())
	}

	for i := range f {
		if f[i] != g[i] {
			fmt.Printf("byte %x differs, in my mischief it is %d, in orig it is %d\n", i, f[i], g[i])
		}
	}
}
