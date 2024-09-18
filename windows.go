//go:build windows
// +build windows

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

func findPlayerPath(player string) (string, error) {
	path, err := findPathInRegistry(player)
	if err == nil {
		return path, nil
	}
	commonPaths := []string{
		filepath.Join(os.Getenv("ProgramFiles"), player, fmt.Sprintf("%s.exe", player)),
		filepath.Join(os.Getenv("ProgramFiles(x86)"), player, fmt.Sprintf("%s.exe", player)),
		filepath.Join(os.Getenv("LocalAppData"), "Programs", player, fmt.Sprintf("%s.exe", player)),
	}

	if player == "potplayer" {
		commonPaths = append(commonPaths,
			filepath.Join(os.Getenv("ProgramFiles"), "DAUM", "PotPlayer", "PotPlayerMini64.exe"),
			filepath.Join(os.Getenv("ProgramFiles(x86)"), "DAUM", "PotPlayer", "PotPlayerMini.exe"),
		)
	}

	for _, path := range commonPaths {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("%s not found", player)
}

func findPathInRegistry(player string) (string, error) {
	var key string
	var valueName string

	switch player {
	case "vlc":
		key = `SOFTWARE\VideoLAN\VLC`
		valueName = "InstallDir"
	case "potplayer":
		key = `SOFTWARE\DAUM\PotPlayer64`
		valueName = "ProgramPath"
	default:
		return "", fmt.Errorf("unsupported player: %s", player)
	}

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	path, _, err := k.GetStringValue(valueName)
	if err != nil {
		return "", err
	}

	if player == "vlc" {
		path = filepath.Join(path, "vlc.exe")
	}

	return path, nil
}
