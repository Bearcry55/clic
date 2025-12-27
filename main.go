package clic

import "fmt"

func main() {
	fmt.Println("welcome to clic a minimal cli in go ....")
	
	cli := New("mycli", "A simple CLI tool built with clic", "1.0.0")
	
	// Flag without value
	cli.Flag("-t", "Test flag to verify it works", func() {
		fmt.Println("This is working!")
	})
	
	// Flag with value
	cli.FlagWithValue("-n", "Set your name", func(value string) {
		fmt.Printf("Hello, %s!\n", value)
	})
	
	cli.FlagWithValue("-p", "Set port number", func(value string) {
		fmt.Printf("Port set to: %s\n", value)
	})
	
	// Parse and execute
	cli.Parse()
}