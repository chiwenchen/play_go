package main

import "fmt"

func main() {
	builder := newNotificationBuilder()

	builder.SetTitle("Hello World")
	builder.SetMessage("This is my first message")
	builder.SetPriority(10)
	builder.SetNotifyType("inapp")

	n, err := builder.Build()

	if err != nil {
		fmt.Println(err)
	}

	builder.SetPriority(5)
	n, err = builder.Build()

	fmt.Println(n)
}
