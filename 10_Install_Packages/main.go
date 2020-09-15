package main

import (
	"fmt"
	// run: `go install` to install dependencies in local pkg
	helper "github.com/cmd-ctrl-q/401GoProject/helper/stringutil"
)


func main() {

	fmt.Println(helper.Reverse("pay pao"))
}