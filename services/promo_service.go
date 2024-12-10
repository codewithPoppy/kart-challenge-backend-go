package services

import (
	"compress/gzip"
	"errors"
	"io"
	"os"
	"strings"
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

	// Rule 2: Check presence in at least two files
	matchCount := 0
	for _, file := range PromoCodeFiles {
		if containsCode(file, code) {
			matchCount++
		}
		if matchCount >= 2 {
			return true, nil
		}
	}

	return false, errors.New("promo code not valid in at least two files")
}

// containsCode checks if a promo code exists in the given gzip file
func containsCode(filepath string, code string) bool {
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

	// Read the file
	content, err := io.ReadAll(gzReader)
	if err != nil {
		return false
	}

	// Check if the promo code is in the file
	return strings.Contains(string(content), code)
}
