package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

func main() {
	const tickets = 50
	var remainingTickets uint = 50
	fmt.Println("")
	// fmt.Println("Welcome to our booking portal")
	var bookingsSlice []string
	var userData = make(map[string]string)
	for {
		helper.GreetUsers(tickets, remainingTickets)

		firstname, lastname, email, userTickets := helper.AskUserData()

		fmt.Println("")

		//input validation
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateInput(firstname, lastname, email, userTickets, remainingTickets)
		// if userTickets > remainingTickets {
		// 	fmt.Println("Invalid ticket number")
		// 	fmt.Println("Please try again")
		// 	//instead of break , continue will start next itterartion
		// 	continue
		// }
		userData["firstname"] = firstname
		userData["lastname"] = lastname
		userData["email"] = email

		if isValidName && isValidEmail && isValidTicketNumber {

			fmt.Printf("User %v booked %v tickets \n", firstname, userTickets)

			remainingTickets = remainingTickets - userTickets
			bookingsSlice = append(bookingsSlice, firstname+" "+lastname)

			fmt.Printf("%v tickets remaning\n", remainingTickets)

			fmt.Printf("Thank you %v for booking %v tickets , you ll recieve a confirmation email at %v\n", firstname, userTickets, email)

			firstNames := []string{}

			//for loop for array/slice gets index and value
			//using _ tells the comp that we are not using it
			for _, booking := range bookingsSlice {
				//seprating at space
				var names = strings.Fields(booking)
				var firstname = names[0]
				firstNames = append(firstNames, firstname)
			}

			fmt.Printf("These are all the bookings %v\n", firstNames)
		} else {

			fmt.Printf("Invalid input , please try again!\n")
			continue

		}
		var noTicketsRemaning bool = remainingTickets == 0

		if noTicketsRemaning {
			// end the loop
			fmt.Println("House Full!!! , Come again next Year.")
			break
		}

	}
}
