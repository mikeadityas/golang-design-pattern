package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sync"
	"time"
)

type piFunc func(int) float64

func wrapLogger(fn piFunc, logger *log.Logger) piFunc {
	return func(n int) float64 {
		now := time.Now()
		result := fn(n)
		logger.Printf("took=%v, n=%v, result=%v", time.Since(now), n, result)
		return result
	}
}

func wrapCache(fn piFunc, cache *sync.Map) piFunc {
	return func(n int) float64 {
		key := fmt.Sprintf("n=%d", n)
		val, ok := cache.Load(key)
		if ok {
			return val.(float64)
		}

		result := fn(n)
		cache.Store(key, result)
		return result
	}
}

func Pi(n int) float64 {
	ch := make(chan float64)

	for k := 0; k <= n; k++ {
		go func(ch chan float64, k float64) {
			ch <- 4 * math.Pow(-1, k) / (2*k + 1)
		}(ch, float64(k))
	}

	result := 0.0
	for k := 0; k <= n; k++ {
		result += <-ch
	}

	return result
}

func main() {
	piWithCache := wrapCache(Pi, &sync.Map{})
	piWithLogger := wrapLogger(piWithCache, log.New(os.Stdout, "[decorator_test] ", 1))

	piWithLogger(100000)
	piWithLogger(50000)
	piWithLogger(25000)

	// These should took 0s since we decorated the original function with caching
	piWithLogger(50000)
	piWithLogger(100000)
	piWithLogger(25000)
}
