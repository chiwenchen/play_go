package main

import (
	"fmt"
	"time"

	"practice.com/play_go/common"
)

func main() {
	m := common.Magazine{Name: "Issue 2021 June", Price: 200, PublishedAt: time.Now()}

	fmt.Println(m.GetName())
}
