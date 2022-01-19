package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

//package level variable declration

var conferenceName = "Go Conference" // if var not specified, : can be used
const conferenceTicket = 50          // : wont work in constant
var remainingTicket uint = 50        //type declartion
// var bookings = []string{}            // slice (not array ) dynamic
// var bookings = make([]map[string]string, 0) // list of map (not array ) dynamic
// initializing with 1 empty map

var bookings = make([]UserData, 0) // list of Userdata (not array ) dynamic

type UserData struct {
	firstName  string
	lastName   string
	email      string
	userTicket uint
}

func main() {

	greetUser()
	// fmt.Printf("the type of conferenceName is : %T \n", conferenceName)

	// array declaration
	// var bookings = [50]string{"kannan"}
	// var bookings = [50]string{}
	// var bookings [50]string //fixed size
	// var bookings []string // slice (not array ) dynamic

	for {

		firstName, lastName, email, userTicket := getUserDetails()
		// userData := getUserDetails()

		isValidUserData := helper.ValidateUserData(firstName,
			lastName, email)

		// userTicket := strconv.FormatUint(string(userData["userTicket"]))
		isValidTicketCounts := validateTicketCounts(userTicket)
		if !isValidUserData && !isValidTicketCounts {
			continue
		} else {
			remainingTicket = remainingTicket - userTicket

			// bookings[0] = firstName + " " + lastName //index for array
			// bookings = append(bookings, firstName+" "+lastName)
			// userData := make(map[string]string)
			// userData["firstName"] = firstName
			// userData["lastName"] = lastName
			// userData["email"] = email
			// userData["userTicket"] = strconv.FormatUint(uint64(userTicket), 10)

			// bookings = append(bookings, userData)

			var userData = UserData{
				firstName:  firstName,
				lastName:   lastName,
				email:      email,
				userTicket: userTicket,
			}

			bookings = append(bookings, userData)
			fmt.Printf("the map : %v\n", bookings)
			//

			go sendTicket(userData)
			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a email confirmation at %v\n", firstName, lastName, userTicket, email)
			fmt.Printf("%v tickets are remaining for %v\n", remainingTicket, conferenceName)

			printUserFirstName()
			// fmt.Printf("%v", bookings[0]) //printing array elements
			// fmt.Printf("%T", bookings)    // type of the array
			// fmt.Printf("%v", bookings)    // printing full array values

		}

		// var noTicketRemaining bool = remainingTicket == 0
		noTicketRemaining := remainingTicket == 0
		if noTicketRemaining {
			fmt.Println("Tickets are booked out!, please try next year!!")
			break
		}
	}
}

func getUserDetails() (string, string, string, uint) {
	// func getUserDetails() map[string]string {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email id: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of ticket: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
	// userData := make(map[string]string)

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["userTicket"] = strconv.FormatUint(uint64(userTicket), 10)
	// return userData
}

func greetUser() {

	fmt.Printf("Hello, Welcome to our %v booking application!\n", conferenceName)
	fmt.Println("Total Ticket:", conferenceTicket, "and Available Ticket:", remainingTicket)
	fmt.Println("Get your ticket here to attend")
}

func printUserFirstName() {
	firstNames := []string{}
	// fmt.Printf("inside function: %v", bookings)
	for _, booking := range bookings {
		// fmt.Printf("key: %v and booking : %v\n", key, booking)
		// firstNames = append(firstNames, booking["firstName"]) //map reader
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("These are booking list: %v\n", firstNames)
}

func validateTicketCounts(userTicket uint) bool {
	if userTicket > remainingTicket {
		fmt.Printf("you cannot Book %v tickets. Only %v tickets are remaining\n", userTicket, remainingTicket)
		return false
	}
	return true
}

func sendTicket(userData UserData) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v \n", userData.userTicket, userData.firstName, userData.lastName)

	fmt.Println("##################")
	fmt.Printf("Sending Tickets:\n %v \n to email id: %v\n", ticket, userData.email)
	fmt.Println("##################")
}
