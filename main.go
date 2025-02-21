// main.go
package main

import (
	"fmt"
	"log"
	"os/user"

	"eetf/cmd"
)

// This program requires the root privileges because it reads from /sys folder.
func isRoot() bool {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("Failed to get current user: %v", err)
	}
	return currentUser.Username == "root"
}

func main() {
	if !isRoot() {
		fmt.Println("This program requires root privileges to run.")
		return
	}

	cmd.Execute()
}
