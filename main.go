package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	InitDB()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n== CLI AUTH SYSTEM ==")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("0. Exit")
		fmt.Print("Choose option: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("Enter username: ")
			scanner.Scan()
			username := strings.TrimSpace(scanner.Text())

			fmt.Print("Enter password: ")
			scanner.Scan()
			password := strings.TrimSpace(scanner.Text())

			Register(username, password)

		case "2":
			fmt.Print("Enter username: ")
			scanner.Scan()
			username := strings.TrimSpace(scanner.Text())

			fmt.Print("Enter password: ")
			scanner.Scan()
			password := strings.TrimSpace(scanner.Text())

			Login(username, password)

		case "0":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
