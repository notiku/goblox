package utils

import (
	"fmt"
)

// GetCdnUrl returns the CDN URL for the given asset.
func GetCdnUrl(hash string) (string, error) {
	if hash == "" {
		return "", fmt.Errorf("hash cannot be empty")
	}

	var i int = 31

	for _, char := range hash {
		i = i ^ int(char)
	}

	return fmt.Sprintf("https://t%d.rbxcdn.com/%s", i%8, hash), nil
}
