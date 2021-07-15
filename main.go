package main

import (
	"fmt"
	"time"

	"practice.com/play_go/common"
)

// Magazine is cool

func main() {
	magazines := []common.Magazine{
		{Name: "About Happy", Price: 300, PublishedAt: time.Now()},
		{"About Happy", 200, time.Now()},
	}

	for _, m := range magazines {
		fmt.Println(m.PublishedAt)
	}

	fmt.Println("This magazine is", getMagazineName(magazines[0]))

	fmt.Println("Total Price is", addUpMagazinePrice(magazines...))
}

func getMagazineName(m common.Magazine) string {
	return m.Name
}

func addUpMagazinePrice(ms ...common.Magazine) int {
	total := 0

	for _, m := range ms {
		total += m.Price
	}

	return total
}

// everything in go is PASS BY VALUE
