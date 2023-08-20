package main

import "fmt"

func main() {
	state := "AP"
	gender := "M"
	height := 123
	age := 60

	var ticketStatus string //to store ticket status info

	switch state {
	case "Karnataka":
		switch gender {
		case "F":
			ticketStatus = "No ticket"
		case "M":
			if height < 110 && age < 5 {
				ticketStatus = "No ticket"
			} else {
				ticketStatus = "Full ticket"
			}
		}
	case "AP":
		switch gender {
		case "F":
			if height < 110 && age < 5 {
				ticketStatus = "No ticket"
			} else {
				ticketStatus = "Full ticket"
			}
		case "M":
			if height < 110 && age < 5 {
				ticketStatus = "No ticket"
			} else {
				ticketStatus = "Full ticket"
			}
		}
	case "Delhi":
		switch gender {
		case "F":
			ticketStatus = "No ticket"
		case "M":
			if height < 130 && age < 7 {
				ticketStatus = "No ticket"
			} else {
				ticketStatus = "Full ticket"
			}
		}
	case "UP":
		switch gender {
		case "F":
			if height < 120 && age < 6 {
				ticketStatus = "No ticket"
			} else {
				ticketStatus = "Full ticket"
			}
		case "M":
			if height < 120 && age < 6 {
				ticketStatus = "No ticket"
			} else {
				ticketStatus = "Full ticket"
			}
		}
	default:
		ticketStatus = "Full ticket"
	}

	fmt.Println("Ticket Status:", ticketStatus)
}
