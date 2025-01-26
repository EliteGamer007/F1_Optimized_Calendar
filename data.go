package main

import (
	"fmt"
	"math"
)

type City struct {
	Name      string
	Latitude  float64
	Longitude float64
}

// Haversine function to calculate the distance between two points on the Earth
func Haversine(lat1, lon1, lat2, lon2 float64) int {
	const R = 6371 // Radius of the Earth in kilometers
	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLon := (lon2 - lon1) * math.Pi / 180.0
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180.0)*math.Cos(lat2*math.Pi/180.0)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	var dist int = int(R * c)
	return dist
}

func main() {
	// List of F1 host cities with their coordinates
	f1Cities := []City{
		{"Sakhir", 26.0325, 50.5106},           // Bahrain Grand Prix
		{"Jeddah", 21.4858, 39.1925},           // Saudi Arabian Grand Prix
		{"Melbourne", -37.8136, 144.9631},      // Australian Grand Prix
		{"Baku", 40.4093, 49.8671},             // Azerbaijan Grand Prix
		{"Miami", 25.7617, -80.1918},           // Miami Grand Prix
		{"Imola", 44.3559, 11.7161},            // Emilia-Romagna Grand Prix (Italy)
		{"Monte Carlo", 43.7384, 7.4246},       // Monaco Grand Prix
		{"Barcelona", 41.3851, 2.1734},         // Spanish Grand Prix
		{"Montreal", 45.5017, -73.5673},        // Canadian Grand Prix
		{"Spielberg", 47.2172, 14.7649},        // Austrian Grand Prix
		{"Silverstone", 52.0786, -1.0169},      // British Grand Prix
		{"Mogyoród", 47.6000, 19.2500},         // Hungarian Grand Prix
		{"Spa-Francorchamps", 50.4542, 5.9714}, // Belgian Grand Prix
		{"Zandvoort", 52.3874, 4.6462},         // Dutch Grand Prix
		{"Monza", 45.6190, 9.2813},             // Italian Grand Prix
		{"Marina Bay", 1.2895, 103.8636},       // Singapore Grand Prix
		{"Suzuka", 34.8431, 136.5419},          // Japanese Grand Prix
		{"Lusail", 25.4848, 51.4503},           // Qatar Grand Prix
		{"Austin", 30.2672, -97.7431},          // United States Grand Prix
		{"Mexico City", 19.4326, -99.1332},     // Mexican Grand Prix
		{"São Paulo", -23.5505, -46.6333},      // Brazilian Grand Prix
		{"Las Vegas", 36.1699, -115.1398},      // Las Vegas Grand Prix
		{"Yas Island", 24.4672, 54.6031},       // Abu Dhabi Grand Prix
		{"Shanghai", 31.2304, 121.4737},        // Chinese Grand Prix
	}

	//Adjacency list to store distances between cities
	adjacencyList := make(map[string]map[string]int)

	for i := 0; i < len(f1Cities); i++ {
		city1 := f1Cities[i]
		adjacencyList[city1.Name] = make(map[string]int)
		for j := 0; j < len(f1Cities); j++ {
			if i != j {
				city2 := f1Cities[j]
				distance := Haversine(city1.Latitude, city1.Longitude, city2.Latitude, city2.Longitude)
				adjacencyList[city1.Name][city2.Name] = distance
			}
		}
	}

	for city, connections := range adjacencyList {
		fmt.Printf("%s:\n", city)
		for connected, distance := range connections {
			fmt.Printf("  -> %s: %d km\n", connected, distance)
		}
	}
}
