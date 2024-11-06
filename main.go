package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const conferenceTickets int = 50
var ConferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

 type userData struct {
	 firstName string
	 lastName string
	 email string
	 numberOfTicktes uint
	 
}

var wg = sync.WaitGroup{}

func main(){
	 

	greetUsers()

	 

		
		firstName, lastName, email, userTickets := getUserInput()		 
		isValidName ,isValidemail , isvalidTicketNumber := validateUserInput(firstName, lastName, email, userTickets )


		if isValidName && isValidemail && isvalidTicketNumber {
				
			bookTiket( userTickets, firstName, lastName, email )

			wg.Add(1)
			go sendTickte(userTickets, firstName, lastName, email)

			firstNames := getfirstName()
			fmt.Printf("These first names of bookings are: %v\n", firstNames)

			if   remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				//break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is to short")
			}
			if !isValidemail{
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isvalidTicketNumber{
				fmt.Println("number of tickets you entered is invalid")
			}				 
	}	
	wg.Wait()	 
}


func greetUsers(){
	fmt.Printf("Welcome to %v booking application.\n",ConferenceName)
	fmt.Printf("we have total of %v tickets and %v are stil available.\n", conferenceTickets,  remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getfirstName() []string{
	firstNames := []string{}
	for _, booking := range bookings { 
		firstNames = append(firstNames , booking.firstName)
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint )(bool, bool, bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidemail := strings.Contains(email,"@")
	isvalidTicketNumber:= userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidemail , isvalidTicketNumber
}

func getUserInput( )(string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name 
	fmt.Println("Enter your First name: ")
	fmt.Scan(&firstName)
	
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter  number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
func bookTiket( userTickets uint, firstName string, lastName string, email string) {
	remainingTickets=  remainingTickets - userTickets

	//create a map for user
	var userData = userData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTicktes: userTickets,
	}

	bookings  = append(bookings, userData)
	fmt.Printf("List of booking is %v\n", bookings)


	fmt.Printf("Thank you %v %v for booking %v tickets. you will reacive a conformation email at %v\n",firstName,lastName,userTickets,email)
	fmt.Printf("%v ticktes  remaining for %v\n",  remainingTickets,ConferenceName)
}
func sendTickte(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v ticktes for %v %v", userTickets, firstName, lastName)
	fmt.Println("#########################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#########################")
	wg.Done()
}