package Foo

import (
	"fmt"
)

func DoStuffAndPrint() bool {
	fn := "Foo"
	fmt.Printf("Loading data from %s", fn)
	return true
}
