package main

import "fmt"

func main() {
	// Print "Hello, Go!" naar de console
	// Hint: Gebruik fmt.Println()

	// Declareer twee variabelen:
	// - naam (string), bijvoorbeeld jouw naam
	// - leeftijd (int), bijvoorbeeld jouw leeftijd

	// Roep de functie groet aan (hieronder gedefinieerd)

	// Loop over een slice van namen en roep de functie groet voor elke naam aan
	namen := []string{"Alice", "Bob", "Charlie"}

	// Loop over een array met ints en gebruik elke waarde als leeftijd voor groet
	// Hint: [n]type{x,x,x} maakt een array van vaste lengte n
	leeftijden := 
	for _, leeftijd := range leeftijden { // _ is de index die we niet gebruiken
		
	}

	// Roep checkLeeftijd aan met een negatieve leeftijd en handel de error
}

// groet print een persoonlijke begroeting met naam en leeftijd
func groet(naam string, leeftijd int) {
	fmt.Printf("Hallo %s, je bent %d jaar oud!\n", naam, leeftijd)
}

// checkLeeftijd geeft een error terug als de leeftijd niet geldig is
func checkLeeftijd(leeftijd int) error {
	if leeftijd < 0 {
		return errors.New("leeftijd kan niet negatief zijn")
	}
	return nil
}