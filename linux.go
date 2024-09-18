//go:build linux
// +build linux

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func findPlayerPath(player string) (string, error) {
	if player == "potplayer" {
		return "", fmt.Errorf("PotPlayer is not supported on Linux")
	}

	commonPaths := []string{
		filepath.Join("/usr/bin", player),
		filepath.Join("/usr/local/bin", player),
		filepath.Join(os.Getenv("HOME"), "bin", player),
	}

	for _, path := range commonPaths {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("%s not found", player)
}
