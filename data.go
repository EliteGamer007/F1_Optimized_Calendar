package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

type City struct {
	Name      string
	Latitude  float64
	Longitude float64
}

func Haversine(lat1, lon1, lat2, lon2 float64) int {
	const R = 6371
	dLat := (lat2 - lat1) * (math.Pi / 180)
	dLon := (lon2 - lon1) * (math.Pi / 180)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*(math.Pi/180))*math.Cos(lat2*(math.Pi/180))*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return int(R * c)
}

func buildDistanceMatrix(cities []City) [][]int {
	n := len(cities)
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			if i != j {
				dist[i][j] = Haversine(cities[i].Latitude, cities[i].Longitude,
					cities[j].Latitude, cities[j].Longitude)
			}
		}
	}
	return dist
}

func tspPath(dist [][]int, start, end int) (int, []int) {
	n := len(dist)
	const INF = 1 << 30
	dp := make([][]int, 1<<n)
	parent := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
		parent[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = INF
			parent[i][j] = -1
		}
	}
	dp[1<<start][start] = 0

	for mask := 0; mask < (1 << n); mask++ {
		for u := 0; u < n; u++ {
			if (mask>>u)&1 == 0 || dp[mask][u] == INF {
				continue
			}
			for v := 0; v < n; v++ {
				if (mask>>v)&1 == 0 {
					nextMask := mask | (1 << v)
					if dp[nextMask][v] > dp[mask][u]+dist[u][v] {
						dp[nextMask][v] = dp[mask][u] + dist[u][v]
						parent[nextMask][v] = u
					}
				}
			}
		}
	}

	// Reconstruct path
	fullMask := (1 << n) - 1
	path := []int{end}
	for mask, u := fullMask, end; u != start; {
		uPrev := parent[mask][u]
		path = append([]int{uPrev}, path...)
		mask ^= (1 << u)
		u = uPrev
	}

	return dp[fullMask][end], path
}

func printAllDistances(cities []City, dist [][]int) {
	for i, city := range cities {
		fmt.Printf("%s:\n", city.Name)
		for j, d := range dist[i] {
			if i != j {
				fmt.Printf("  -> %-20s: %4d km\n", cities[j].Name, d)
			}
		}
		fmt.Println()
	}
}

func printPathWithDistances(cities []City, dist [][]int, path []int, total int) {
	fmt.Println("Optimal Route:")
	sum := 0
	for i := 0; i < len(path)-1; i++ {
		from := cities[path[i]]
		to := cities[path[i+1]]
		d := dist[path[i]][path[i+1]]
		fmt.Printf("  %s -> %-20s: %4d km\n", from.Name, to.Name, d)
		sum += d
	}
	fmt.Printf("\nTotal Distance: %d km\n", sum)
}

func main() {
	mode := flag.String("mode", "path", "Choose mode: 'distances' or 'path'")
	flag.Parse()

	cities := []City{
		{"Sakhir", 26.0325, 50.5106},
		{"Jeddah", 21.4858, 39.1925},
		{"Melbourne", -37.8136, 144.9631},
		{"Baku", 40.4093, 49.8671},
		{"Miami", 25.7617, -80.1918},
		{"Imola", 44.3559, 11.7161},
		{"Monte Carlo", 43.7384, 7.4246},
		{"Barcelona", 41.3851, 2.1734},
		{"Montreal", 45.5017, -73.5673},
		{"Spielberg", 47.2172, 14.7649},
		{"Silverstone", 52.0786, -1.0169},
		{"Mogyoród", 47.6, 19.25},
		{"Spa-Francorchamps", 50.4542, 5.9714},
		{"Zandvoort", 52.3874, 4.6462},
		{"Monza", 45.619, 9.2813},
		{"Marina Bay", 1.2895, 103.8636},
		{"Suzuka", 34.8431, 136.5419},
		{"Lusail", 25.4848, 51.4503},
		{"Austin", 30.2672, -97.7431},
		{"Mexico City", 19.4326, -99.1332},
		{"São Paulo", -23.5505, -46.6333},
		{"Las Vegas", 36.1699, -115.1398},
		{"Yas Island", 24.4672, 54.6031},
		{"Shanghai", 31.2304, 121.4737},
	}

	cityIndex := make(map[string]int)
	for i, c := range cities {
		cityIndex[c.Name] = i
	}

	startCity := "Melbourne"
	endCity := "Yas Island"
	start := cityIndex[startCity]
	end := cityIndex[endCity]

	dist := buildDistanceMatrix(cities)

	switch strings.ToLower(*mode) {
	case "distances":
		printAllDistances(cities, dist)
	case "path":
		total, path := tspPath(dist, start, end)
		printPathWithDistances(cities, dist, path, total)
	default:
		fmt.Println("Invalid mode. Use -mode=distances or -mode=path")
		os.Exit(1)
	}
}
