package main

import (
	"fmt"
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
	if size <= 0 {
		return []int{}
	}

	res := make([]int, size)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < size; i++ {
		res[i] = r.Int()
	}

	return res
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]

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
		return 0
	}

	if CHUNKS <= 1 || n == 1 {
		return maximum(data)
	}

	chunkSize := (n + CHUNKS - 1) / CHUNKS

	type pair struct {
		idx int
		val int
	}

	maxParts := make([]int, 0, CHUNKS)
	maxPartsMu := sync.Mutex{}

	var wg sync.WaitGroup

	for c := 0; c < CHUNKS; c++ {

		start := c * chunkSize

		if start >= n {
			break
		}

		end := start + chunkSize

		if end > n {
			end = n
		}

		wg.Add(1)

		go func(s, e int) {
			defer wg.Done()
			localMax := maximum(data[s:e])
			maxPartsMu.Lock()
			maxParts = append(maxParts, localMax)
			maxPartsMu.Unlock()
		}(start, end)
	}

	wg.Wait()

	return maximum(maxParts)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	start := time.Now()
	data := generateRandomElements(SIZE)
	genElapsed := time.Since(start).Microseconds()

	fmt.Printf("Генерация заняла: %d μs\n\n", genElapsed)

	fmt.Println("Ищем максимальное значение в один поток")
	start = time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
