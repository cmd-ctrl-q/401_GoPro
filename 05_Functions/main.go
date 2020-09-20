package main

import (
	"fmt"

	"github.com/cmd-ctrl-q/401_GOPRO/05_Functions/myfunc"
)

func main() {

	rn1, rn2, diff := myfunc.CalcDiff(myfunc.PickRandom, 0, 10)

	fmt.Printf("min:\t\t%d\nmax:\t\t%d\ndifference:\t%d\n", rn1, rn2, diff)
}
