package main

import (
	"os"
	"os/exec"
	"log"
)

func main() {
	// exit on missing arg
	if len(os.Args) < 2 {
		log.Fatal("Error: arg missing") // calls os.Exit 1 after logging
	}

	base := "https://golang.org/pkg/"
	arg := os.Args[1]
	url := base + arg

	exec.Command("open", url).Run()
}
