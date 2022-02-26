//Ticket Booking APp
package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTicket int = 50
	var remainingTicket uint = 50
	fmt.Println("Welcome to ", conferenceName, "Booking Application")
	fmt.Println("We Have total of ", conferenceTicket, " Tickets and ", remainingTicket, " are still Available")
	fmt.Println("Get Your Tickets here to attend")

	var bookigns []string
	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter Your First Name:")
		fmt.Scanln(&firstName)

		fmt.Println("Enter your last Name:")
		fmt.Scanln(&lastName)

		fmt.Println("Enter your Email Address:")
		fmt.Scanln(&email)

		fmt.Println("Enter the Number of Tickets:")
		fmt.Scanln(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTickets := userTickets > 0 && userTickets <= remainingTicket

		if isValidName && isValidEmail && isValidTickets {

			if userTickets < remainingTicket {

				remainingTicket = remainingTicket - userTickets
				bookigns = append(bookigns, firstName+" "+lastName)
				// fmt.Printf("The Whole Slice %v\n", bookigns)
				// fmt.Printf("The First Value %v\n", bookigns[0])
				// fmt.Printf("Array Type: %T\n", bookigns)
				// fmt.Printf("Slice lenght %v\n", len(bookigns))

				// fmt.Println("Thank You",firstName,"",lastName,"for booking",email, "Booked", userTickets)
				fmt.Printf("Thank You %v %v for Booking %v tickets. You Will recive  a confirmation email at  %v \n", firstName, lastName, userTickets, email)
				fmt.Println(remainingTicket, "Tickets remaining for ", conferenceName)

				firstNames := []string{}
				for _, booking := range bookigns {
					var names = strings.Fields(booking)
					firstNames = append(firstNames, names[0])
				}
				fmt.Println("The first Names of Bookings are :\n", firstNames)

				noTicketsAvailable := remainingTicket == 0

				if noTicketsAvailable {
					fmt.Println("Our Conference is booked our. Come back next Year")
					break

				}

			} else {
				fmt.Println("we only have", remainingTicket, "Tickets remaining, so you can't book", userTickets)

			}

		} else {
			fmt.Println("Invalid Input Please Try Again")
		}

	}
}
