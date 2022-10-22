package helper

//package needs to be defined
import (
	"fmt"
	"strings"
)

// return value types
func ValidateInput(firstname string, lastname string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {

	isValidName := len(firstname) >= 2 && len(lastname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets

	//in go we can return any values from a funct
	return isValidName, isValidEmail, isValidTicketNumber
}

func GreetUsers(tickets uint, remainingTickets uint) {
	fmt.Println("Welcome to our booking portal")
	fmt.Printf("%v tickets in total\n", tickets)
	fmt.Printf("%v tickets remaning\n", remainingTickets)
	fmt.Println("Book your tickets here.")
	fmt.Printf("\n")

}

func AskUserData() (string, string, string, uint) {
	//ask user for their name

	var firstname string
	var lastname string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstname) //add a pointer to the var

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastname) //add a pointer to the var

	fmt.Println("Enter your emai:")
	fmt.Scan(&email) //add a pointer to the var

	fmt.Println("no of tickets you'd like to buy?")
	fmt.Scan(&userTickets)

	return firstname, lastname, email, userTickets
}
