package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/san-lab/sx402/facilitator"
	"golang.org/x/term"
)

func main() {
	withDemoStore := flag.Bool("demoStore", false, "starts the demo store under /store")
	flag.Parse()

	fmt.Print("Enter facilitator's keyfile password: ")

	passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println() // move to next line after input
	if err != nil {
		fmt.Println("Failed to read password: %w", err)
	}

	//*withDemoStore = true
	facilitator.Start(*withDemoStore, passwordBytes)

}
