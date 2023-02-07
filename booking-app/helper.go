package main

import (
	"fmt"
	"strings"
)

func GetUserInput() UserData {
	var city string
	var userName string
	var userTicket int

	fmt.Println("Choose city:")
	fmt.Scan(&city)
	switch city {
	case "NewYork":
		fmt.Println("You select New York")
	case "Singapore":
		fmt.Println("You select Singapore")
	case "London", "Berlin":
		fmt.Println("VIP booking")
	default:
		fmt.Println("No valid city selected")
	}
	// ask for user name
	fmt.Println("Enter your name:")
	fmt.Scan(&userName)
	userName = strings.Replace(userName, "-", " ", -1)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTicket)

	var userData = UserData{
		city:       city,
		userName:   userName,
		userTicket: userTicket,
	}
	return userData
}
