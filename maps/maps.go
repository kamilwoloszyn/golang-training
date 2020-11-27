package maps

import "fmt"

type Person struct {
	age       uint8
	eyesColor string
	height    uint16
	weight    int16
	gender    string
}

func run() {
	people := map[string]Person{
		"Chris": {
			30,
			"Blue",
			199,
			90,
			"Male",
		},
		"Vlad": {
			20,
			"Brown",
			160,
			60,
			"Male",
		},
		"Marika": {
			24,
			"Blue",
			177,
			55,
			"Female",
		},
	}

	for i, v := range people {
		// As you can see i is key id, and v element
		fmt.Printf("The name of %s and his/her personal data %v\n", i, v)
	}

	//Or you can access to the map more directly
	fmt.Println(people["Marika"])
	//search for specific person
	who, found := people["Chris"]
	fmt.Printf("Who: %v Found: %v\n", who, found)
	//Delete specific person
	delete(people, "Chris")
	//And his again
	who, found = people["Chris"]
	//As you can see, value of who is nil!
	fmt.Printf("Who: %v Found: %v", who, found)
}
