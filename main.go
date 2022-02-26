//Ticket Booking APp
package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTicket int = 50
	var remainingTicket uint = 50
	fmt.Println("Welcome to ", conferenceName, "Booking Application")
	fmt.Println("We Have total of ", conferenceTicket, " Tickets and ", remainingTicket, " are still Available")
	fmt.Println("Get Your Tickets here to attend")

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
	remainingTicket = remainingTicket - userTickets

	// fmt.Println("Thank You",firstName,"",lastName,"for booking",email, "Booked", userTickets)
	fmt.Printf("Thank You %v %v for Booking %v tickets. You Will recive  a confirmation email at  %v \n", firstName, lastName, userTickets, email)
	fmt.Println(remainingTicket, "Tickets remaining for ", conferenceName)

}
