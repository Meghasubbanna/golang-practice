// Create Company struct Id,Name, Website, Telehone, Fax, Address // Address must be a promoted filed

// Create Address struct City, Line1, Street, State, Country, PinCode , Map // Map must be a promoted field

// Create Map struct Lan, Lat

// -> Can you access Lan and Lap with the object of Company?

package main

import "fmt"

type Map struct {
	Lan float64
	Lat float64
}

type Address struct {
	City    string
	Line1   string
	Street  string
	State   string
	Country string
	PinCode string
	Map     // Promoted field
}

type Company struct {
	Id        int
	Name      string
	Website   string
	Telephone int
	Fax       int
	Address   // Promoted field
}

func main() {
	var megha map[Address]Map
	megha = make(map[Address]Map)
	obj := Company{Id: 1, Name: "VSCO", Website: "vsco.com", Telephone: 9886547634, Fax: 67643}
	obj.Map.Lat = 12.879
	obj.Map.Lan = 14.087
	megha[obj.Address] = obj.Map
	a := megha[obj.Address]
	fmt.Println(a)

}

// func main() {
// 	company := Company{
// 		Id:        1,
// 		Name:      "Megha",
// 		Website:   "www.victoria.com",
// 		Telephone: "123-456-7890",
// 		Fax:       "987-654-3210",
// 		Address: Address{
// 			City:    "Chamarajanagar",
// 			Line1:   "Ramasamudra",
// 			Street:  "Hallada street",
// 			State:   "Karnataka",
// 			Country: "India",
// 			PinCode: "12345",
// 			Map: Map{
// 				Lan: 12.3456,
// 				Lat: 78.9012,
// 			},
// 		},

// 	}

// fmt.Println("Lan:", company.Address.Map.Lan)
// fmt.Println("Lat:", company.Address.Map.Lat)
