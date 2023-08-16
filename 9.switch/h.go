package main

func main() {
	switch day := 4; day {
	case 1:
		println("sunday")
	case 2:
		println("monday")
	case 3:
		println("tuesday")
	case 4:
		println("wednesday")
	case 5:
		println("thursday")
	case 6:
		println("friday")
	case 7:
		println("saturday")
	default:
		println("no day")

	}
	println("another type of switch")
	num := 500

	switch { // This is called empty switch
	case num > 0 && num < 50:
		println("Num is between the range of 0-50")
	case num >= 50 && num < 100:
		println("Num is between the range of 50-100")
	case num >= 100 && num < 200:
		println("Num is between the range of 100-200")
	default:
		println("Num is 200 or greater")
	}

	char := 'A'
	switch char {
	case 65, 69, 73, 79, 85:
		println(string(char), "is a upper case vowel")
	case 'a', 'e', 'i', 'o', 'u':
		println(string(char), "is a lower case vowel")
	default:
		println(string(char), "is either a consonent or a special character")
	}
}
