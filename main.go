package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size < 0 {
		log.Println("Размер слайса отрицательный")
		return nil
	}
	if size == 0 {
		return []int{} //
	}
	resultRand := make([]int, size)
	for i := 0; i < len(resultRand); i++ {
		resultRand[i] = rand.Int()
	}
	return resultRand
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		log.Println("Слайс пуст")
		return 0
	}
	max := data[0] // предполагаем, что первый элемент — максимальный
	for i := 1; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	n := len(data)
	if n == 0 {
		log.Println("Слайс пуст, невозможно найти максимум")
		return 0
	}
	chunkSize := n / CHUNKS
	var wg sync.WaitGroup
	maxima := make([]int, CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = n
		}
		go func(i, start, end int) {
			defer wg.Done()

			max := data[start]
			for j := start + 1; j < end; j++ {
				if data[j] > max {
					max = data[j]
				}
			}
			maxima[i] = max
		}(i, start, end)
	}

	wg.Wait()

	overallMax := maxima[0]
	for i := 1; i < len(maxima); i++ {
		if maxima[i] > overallMax {
			overallMax = maxima[i]
		}
	}

	return overallMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	resultRand := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(resultRand)
	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %v ms\n", max, elapsed.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(resultRand)
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %v ms\n", max, elapsed.Microseconds())
}
