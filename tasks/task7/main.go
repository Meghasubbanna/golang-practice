//Implement this using if else State : Karnataka , AP, Delhi ,UP Gender: M,F Age: >0 Height:

//State	      Gender	Height	Age	     Ticket Status
//Karnataka	    F			               No ticket
//AP	        F	    <110cm	<5y	       No ticket
//Delhi     	F			               No Ticket
//UP	        F	    <120cm	<6y        No ticket
//Karnataka	    M	    <110cm	<5y        No ticket
//AP	        M	    <110cm  <5y        No ticket
//Delhi     	M	    <130cm	<7y	       No Ticket
//UP	        M	    <120cm	<6y        No ticket
//Other than the above table ,It is a full ticket

package main

import "fmt"

func main() {
	state := "UP"
	gender := "M"
	height := 107
	age := 3
	ticket(state, gender, height, age)

}

func ticket(state string, gender string, height int, age int) {
	if state == "Karnataka" && gender == "F" {
		fmt.Println("No ticket")
	} else if state == "AP" && gender == "F" && height <= 110 && age <= 5 {
		fmt.Println("No ticket")
	} else if state == "Delhi" && gender == "F" {
		fmt.Println("No ticket")
	} else if state == "UP" && gender == "F" && height <= 120 && age <= 6 {
		fmt.Println("No ticket")
	} else if state == "Karnataka" && gender == "M" && height <= 110 && age <= 5 {
		fmt.Println("No ticket")
	} else if state == "AP" && gender == "M" && height <= 110 && age <= 5 {
		fmt.Println("No ticket")
	} else if state == "Delhi" && gender == "M" && height <= 130 && age <= 7 {
		fmt.Println("No ticket")
	} else if state == "UP" && gender == "M" && height <= 120 && age <= 6 {
		fmt.Println("No ticket")
	} else {
		fmt.Println("Full ticket")
	}

}
