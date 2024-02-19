package main

import (
	"fmt"
	"time"
)

func main() {
	// Input
	var kenteken string
	fmt.Print("Voer uw kenteken in: ")
	fmt.Scan(&kenteken)

	// Lijst van hard-coded kentekens
	toegestaneKentekens := []string{"ABC123", "XYZ789", "123DEF"}

	// Checkt het ingevulde kenteken.
	if bevatKenteken(toegestaneKentekens, kenteken) {
		groet := bepaalGroet()
		bericht := fmt.Sprintf("%s! Welkom bij Fonteyn Vakantieparken", groet)
		fmt.Println(bericht)
	} else {
		fmt.Println(":U heeft helaas geen toegang tot het parkeerterrein")
	}
}

func bevatKenteken(kentekens []string, kenteken string) bool {
	for _, k := range kentekens {
		if k == kenteken {
			return true
		}
	}
	return false
}

func bepaalGroet() string {
	huidigeTijd := time.Now().Hour()

	if huidigeTijd >= 7 && huidigeTijd < 12 {
		return "Goedemorgen"
	} else if huidigeTijd >= 12 && huidigeTijd < 18 {
		return "Goedemiddag"
	} else if huidigeTijd >= 18 && huidigeTijd < 23 {
		return "Goedenavond"
	} else {
		return "Sorry, de parkeerplaats is ’s nachts gesloten"
	}
}
