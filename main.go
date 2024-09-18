// main.go
package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: program <protocol://link>")
		os.Exit(1)
	}

	arg := os.Args[1]

	var player, protocol string
	if strings.HasPrefix(arg, "vlc://") {
		player = "vlc"
		protocol = "vlc://"
	} else if isWindows() && strings.HasPrefix(arg, "potplayer://") {
		player = "potplayer"
		protocol = "potplayer://"
	} else {
		fmt.Println("Invalid protocol. Use vlc:// or potplayer:// (PotPlayer is only supported on Windows)")
		os.Exit(1)
	}

	link := strings.TrimPrefix(arg, protocol)

	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") && !strings.HasPrefix(link, "file://") {
		fmt.Println("Invalid link. Must start with http://, https://, or file://")
		os.Exit(1)
	}

	isFileProtocol := strings.HasPrefix(link, "file://")
	if isFileProtocol {
		link = strings.TrimPrefix(link, "file://")
		link = strings.TrimPrefix(link, "/")
		if isWindows() {
			link = strings.Replace(link, "/", "\\", -1)
		}
		link, _ = url.QueryUnescape(link)
	}

	playerPath, err := findPlayerPath(player)
	if err != nil {
		fmt.Printf("Error finding %s path: %v\n", player, err)
		os.Exit(1)
	}

	var cmd *exec.Cmd
	switch player {
	case "vlc":
		if isFileProtocol {
			cmd = exec.Command(playerPath, link)
		} else {
			cmd = exec.Command(playerPath, "--open", link)
		}
	case "potplayer":
		cmd = exec.Command(playerPath, link)
	}

	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting player:", err)
		os.Exit(1)
	}
}

func isWindows() bool {
	return os.PathSeparator == '\\'
}
