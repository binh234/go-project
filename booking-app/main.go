package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var conferenceName = "Happy Concert"

const conferenceTickets = 50

var remainingTickets int = 50

var wg = sync.WaitGroup{}

type UserData struct {
	userName   string
	city       string
	userTicket int
}

func main() {
	var bookings = make([]UserData, 0)

	greetUser()

	for remainingTickets > 0 && len(bookings) < 50 {
		userData := GetUserInput()
		userTicket := userData.userTicket
		if userTicket > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, you can't book %v tickets\n", remainingTickets, userTicket)
			continue
		}
		remainingTickets = remainingTickets - userTicket
		bookings = append(bookings, userData)

		fmt.Printf("There are %v tickets left\n", remainingTickets)
		wg.Add(1)
		go sendTicket(userData.userName, userTicket)

		firstNames := getFirstNames(bookings)
		fmt.Printf("All bookings: %v\n", bookings)
		fmt.Printf("All bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is out of tickets. Please come back next year")
			break
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []UserData) []string {
	var firstNames = []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking.userName)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func sendTicket(userName string, userTicket int) {
	time.Sleep(50 * time.Second)
	fmt.Printf("%v tickets for %v\n", userTicket, userName)
	wg.Done()
}
