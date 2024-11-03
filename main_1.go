package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Function to initialize a 3D array with random temperature values
func initializeTemps(days, zones, hours int) [][3][24]float64 {
	rand.Seed(time.Now().UnixNano())
	var temperatures = make([][3][24]float64, days)

	for d := 0; d < days; d++ {
		for z := 0; z < zones; z++ {
			for h := 0; h < hours; h++ {
				temperatures[d][z][h] = 15 + rand.Float64()*10 // temperatures between 15 and 25°C
			}
		}
	}
	return temperatures
}

// Function to calculate the daily average temperature for each zone concurrently
func calculateDailyAverages(temps [][3][24]float64, wg *sync.WaitGroup, results chan<- [][3]float64) {
	defer wg.Done()
	dailyAverages := make([][3]float64, len(temps))

	// Calculate average temperature concurrently for each day
	for d := 0; d < len(temps); d++ {
		wg.Add(1)
		go func(day int) {
			defer wg.Done()
			for z, zoneData := range temps[day] {
				var sum float64
				for _, temp := range zoneData {
					sum += temp
				}
				dailyAverages[day][z] = sum / float64(len(zoneData))
			}
		}(d)
	}
	// Wait until all daily averages are calculated
	wg.Wait()
	results <- dailyAverages
}

// Function to find the hottest zone for each day concurrently
func hottestZonePerDay(averages [][3]float64, wg *sync.WaitGroup, results chan<- []int) {
	defer wg.Done()
	hottestZones := make([]int, len(averages))

	for d, dayAverages := range averages {
		wg.Add(1)
		go func(day int, avgs [3]float64) {
			defer wg.Done()
			maxTemp := avgs[0]
			zone := 0
			for z, avg := range avgs {
				if avg > maxTemp {
					maxTemp = avg
					zone = z
				}
			}
			hottestZones[day] = zone
		}(d, dayAverages)
	}
	wg.Wait()
	results <- hottestZones
}

func main() {
	// Initialize data
	days, zones, hours := 7, 3, 24
	temperatures := initializeTemps(days, zones, hours)

	// Channels and WaitGroup for managing concurrency
	dailyAveragesChannel := make(chan [][3]float64, 1)
	hottestZonesChannel := make(chan []int, 1)
	var wg sync.WaitGroup

	// Start calculating daily averages concurrently
	wg.Add(1)
	go calculateDailyAverages(temperatures, &wg, dailyAveragesChannel)

	// Retrieve daily averages from channel and calculate hottest zones concurrently
	wg.Add(1)
	go func() {
		dailyAverages := <-dailyAveragesChannel
		go hottestZonePerDay(dailyAverages, &wg, hottestZonesChannel)
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	close(dailyAveragesChannel)
	close(hottestZonesChannel)

	// Retrieve results and display them
	dailyAverages := <-dailyAveragesChannel
	fmt.Println("Daily Average Temperatures for Each Zone:")
	for d, averages := range dailyAverages {
		fmt.Printf("Day %d: Zone 1: %.2f°C, Zone 2: %.2f°C, Zone 3: %.2f°C\n", d+1, averages[0], averages[1], averages[2])
	}

	hottestZones := <-hottestZonesChannel
	fmt.Println("\nHottest Zone Each Day:")
	for d, zone := range hottestZones {
		fmt.Printf("Day %d: Zone %d\n", d+1, zone+1)
	}
}
