package bookstore

import (
	"math"
)

var discounts = map[int]float64{
	1: 1.00,
	2: 0.95,
	3: 0.90,
	4: 0.80,
	5: 0.75,
}

const bookPrice = 800

func findOptimalCost(bookCounts []int, currentCost int) int {
	// Remove zero counts for easy processing
	var books []int
	for _, count := range bookCounts {
		if count > 0 {
			books = append(books, count)
		}
	}

	if len(books) == 0 {
		return currentCost
	}

	minTotalCost := math.MaxInt

	// Try to form groups of different sizes (5, 4, 3, etc.)
	for size := 5; size > 0; size-- {
		if len(books) < size {
			continue
		}

		// Create a copy of bookCounts and reduce the selected group
		temp := make([]int, len(books))
		copy(temp, books)
		count := 0
		for i := 0; i < len(temp) && count < size; i++ {
			if temp[i] > 0 {
				temp[i]--
				count++
			}
		}

		// Recursively find the minimum cost
		totalCost := findOptimalCost(temp, currentCost+int(float64(size*bookPrice)*discounts[size]))
		if totalCost < minTotalCost {
			minTotalCost = totalCost
		}
	}

	return minTotalCost
}

func Cost(books []int) int {
	bookCounts := make([]int, 5)
	for _, book := range books {
		bookCounts[book-1]++
	}
	return findOptimalCost(bookCounts, 0)
}
