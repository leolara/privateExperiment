package main

import (
	"fmt"
	"github.com/leolara/privateExperiment/other"
)

func main () {
	fmt.Println(other.GetGetter().Get().PublicField)
}
