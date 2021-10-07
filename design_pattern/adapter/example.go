package main

import "fmt"

func main() {
	sony := Sony{
		channel: 10,
		vol:     5,
		power:   false,
	}

	samsung := SamSung{
		currentChannel: 10,
		currentVol:     5,
		isTvOn:         false,
	}

	sony.tvOn(true)

	// This will not work
	// samsung.tvOn(true)

	sa := samsungAdapter{tv: &samsung}
	sa.tvOn(true)

	fmt.Println(sony)
}
