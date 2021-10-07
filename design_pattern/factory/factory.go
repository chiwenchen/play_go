package main

import "fmt"

// 用這個func當作工廠創造出各個實例，一律使用 interface 當作回傳
func newPublication(pubType string, name string, pages int, publisher string) (iPublication, error) {
	if pubType == "magazine" {
		return createMagazine(name, pages, publisher), nil
	}

	if pubType == "newspaper" {
		return createNewsPaper(name, pages, publisher), nil
	}
	return nil, fmt.Errorf("No publication type found")
}
