package main

import (
	"errors"
	"fmt"
)

var (
	twinErr1 = errors.New("I am twin error")
	twinErr2 = errors.New("I am twin error")
)

func main() {
	fmt.Println(twinErr1 == twinErr2)
	fmt.Println(twinErr1.Error() == twinErr2.Error())
}
