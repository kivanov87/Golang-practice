package main

func main() {
	var degreese int
	var timeofDay, outfit, shoes string
	switch timeofDay {
	case "Morning":
		if degreese <= 10 || degreese <= 18 {
			outfit = "Sweatshirt"
			shoes = "Sneakers"
		}
	case "Afternoon":
		if degreese > 18 || degreese <= 24 {
			outfit = "Shirt"
			shoes = "Moccasins"
		}

	case "Evening":

	}
}
