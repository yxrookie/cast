package main

import (
	"fmt"

	"github.com/spf13/cast"
)

func main() {
	var s string = "hello"
	newbool := cast.ToInt(s)
	fmt.Println(newbool)
}