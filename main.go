package main

import "fmt"

func main() {
	var smsSendingLimit int
	var costPerSms float64
	var hasPermission bool
	var userName string

	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimit,
		costPerSms,
		hasPermission,
		userName,
	)
}