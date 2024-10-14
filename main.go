package main

import (
	"fmt"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(firstName, lastName, email, userTickets)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("These first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out.  Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("Your name is invalid.  Please try again.\n")
			}
			if !isValidEmail {
				fmt.Printf("Your email is invalid.  Please try again.\n")
			}
			if !isValidTicketNumber {
				fmt.Printf("Your ticket number is invalid.  Please try again.\n")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("How many tickets would you like to book? ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets.  You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("#######################################\n")
	fmt.Printf("Sending ticket: %v\nTo email address %v\n", ticket, email)
	fmt.Printf("#######################################\n")
}
