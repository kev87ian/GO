package main

import (
	"fmt"
)

func main() {

	/*	casting from float to int */
	//accountAge := 2.6
	//accountAgeInt := int(accountAge)

	//fmt.Printf("The data type of acocuntAgeInt is %T\n", accountAgeInt)
	//fmt.Println("And it has existed for", accountAgeInt, "years")

	const premiumPlanName string = "Premium plan"
	const basicPlanName = "Basic Plan"

	const secondsInMinute = 60
	const minutesInAnHour = 60
	const secondsInHour = secondsInMinute * minutesInAnHour

	fmt.Println("Number of seconds in an hour: ", secondsInHour)

	/*Formatting Strings*/
	const name = "Saul Goodman"
	const openRate = 30.5
	msg := fmt.Sprintf("Hi, %s, your open rate is %.1f percent", name, openRate)
	println(msg)

}
