package services

import (
	"compress/gzip"
	"errors"
	"io"
	"os"
	"strings"
	"sync"
)

// PromoCodeFiles contains the paths to the promo code files
var PromoCodeFiles = []string{
	"backend-challenge/couponbase1.gz",
	"backend-challenge/couponbase2.gz",
	"backend-challenge/couponbase3.gz",
}

// ValidatePromoCode validates a promo code based on the rules:
// 1. It must be between 8 and 10 characters.
// 2. It must be present in at least two of the coupon files.
func ValidatePromoCode(code string) (bool, error) {
	// Rule 1: Length validation
	if len(code) < 8 || len(code) > 10 {
		return false, errors.New("invalid promo code length")
	}

	// Rule 2: Check presence in at least two files concurrently
	var wg sync.WaitGroup
	matchCount := 0
	mu := sync.Mutex{}
	done := make(chan struct{})

	for _, file := range PromoCodeFiles {
		wg.Add(1)
		go func(filepath string) {
			defer wg.Done()
			if containsCode(filepath, code, done) {
				mu.Lock()
				matchCount++
				// Signal to stop further processing if matchCount >= 2
				if matchCount >= 2 {
					close(done)
				}
				mu.Unlock()
			}
		}(file)
	}

	wg.Wait()

	if matchCount >= 2 {
		return true, nil
	}

	return false, errors.New("promo code not valid in at least two files")
}

// containsCode checks if a promo code exists in the given gzip file
// Stops reading further if the done channel is closed
func containsCode(filepath string, code string, done <-chan struct{}) bool {
	file, err := os.Open(filepath)
	if err != nil {
		return false
	}
	defer file.Close()

	// Create a new gzip reader
	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return false
	}
	defer gzReader.Close()

	// Create a buffered reader for line-by-line reading
	buf := make([]byte, 4096)
	for {
		select {
		case <-done:
			return false
		default:
		}

		n, err := gzReader.Read(buf)
		if err != nil && err != io.EOF {
			return false
		}
		if n == 0 {
			break
		}

		// Check if the promo code is in the current chunk
		if strings.Contains(string(buf[:n]), code) {
			return true
		}
	}

	return false
}
