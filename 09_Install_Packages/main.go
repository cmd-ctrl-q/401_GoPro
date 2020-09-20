package main

import (
	"fmt"
	// run: `go install` to install dependencies in local pkg
	// change username
	helper "github.com/cmd-ctrl-q/401_GOPRO/helper/stringutil"
)

func main() {

	fmt.Println(helper.Reverse("pay pao"))

}
