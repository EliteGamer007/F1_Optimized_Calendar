package main

import "fmt"

func main() {
    // Create a map with numbers as keys and exact F1 race cities as values
    f1Cities := map[int]string{
		0:  "Sakhir",            // Bahrain Grand Prix
		1:  "Jeddah",            // Saudi Arabian Grand Prix
		2:  "Melbourne",         // Australian Grand Prix
		3:  "Baku",              // Azerbaijan Grand Prix
		4:  "Miami",             // Miami Grand Prix
		5:  "Imola",             // Emilia-Romagna Grand Prix (Italy)
		6:  "Monaco",            // Monaco Grand Prix
		7:  "Barcelona",         // Spanish Grand Prix
		8:  "Montreal",          // Canadian Grand Prix
		9:  "Spielberg",         // Austrian Grand Prix
		10: "Silverstone",       // British Grand Prix
		11: "Mogyoród",          // Hungarian Grand Prix
		12: "Spa-Francorchamps", // Belgian Grand Prix
		13: "Zandvoort",         // Dutch Grand Prix
		14: "Monza",             // Italian Grand Prix
		15: "Marina Bay",        // Singapore Grand Prix
		16: "Suzuka",            // Japanese Grand Prix
		17: "Lusail",            // Qatar Grand Prix
		18: "Austin",            // United States Grand Prix
		19: "Mexico City",       // Mexican Grand Prix
		20: "São Paulo",         // Brazilian Grand Prix
		21: "Las Vegas",         // Las Vegas Grand Prix
		22: "Yas Island",        // Abu Dhabi Grand Prix
		23: "Shanghai",          // Chinese Grand Prix

    for number, city := range f1Cities {
        fmt.Printf("%d: %s\n", number, city)
    }
}
