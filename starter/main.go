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

	// Maak een map[string]int waarbij de key (string) een naam uit de namen lijst bevat en de value (int) een leeftijd uit de leeftijden array is. 
	var leeftijdMap <TYPE>;

	// Implementeer de logica om de namen en leeftijden aan elkaar te koppelen.
	for index, naam := range namen {
		// Gebruik de index van namen om de leeftijd op te halen uit leeftijden
	}


	// Loop door de leeftijdMap heen en print de naam en leeftijd uit de map.
	for VARIABLE, VARIABLE := range leeftijdMap {

		// Tip: Zoek Sprintf op in pkg.go.dev!
		message := fmt.Sprintf("%s is %d jaar oud", VARIABLE, VARIABLE) // naam, leeftijd

		// Print message
	}
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
