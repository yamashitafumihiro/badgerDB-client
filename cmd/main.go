package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/yamashitafumihiro/badgerDB-client/internal/db"
)

func main() {
	database, err := db.InitDB("./data")
	if err != nil {
		log.Fatalf("failed to initialize DB: %v", err)
	}
	defer database.Close()

	fmt.Println("Welecome to BadgerDB CLI")
	fmt.Println("Type 'help' to see the list of available commands")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "exit" {
			fmt.Println("Exiting CLI. Goodbye!")
			break
		}

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		command := strings.ToLower(args[0])
		switch command {
		case "put":
			if len(args) != 3 {
				fmt.Println("Usage: put <key> <value>")
				continue
			}
			key := args[1]
			value := args[2]
			if err := db.WriteData(database, key, value); err != nil {
				log.Printf("failed to write data: %v", err)
			} else {
				fmt.Printf("Successfully wrote key: %s, value: %s\n", key, value)
			}

		case "get":
			if len(args) != 2 {
				fmt.Println("Usage: get <key>")
				continue
			}
			key := args[1]
			value, err := db.ReadData(database, key)
			if err != nil {
				log.Printf("failed to read data: %v", err)
			} else {
				fmt.Printf("Key: %s, Value: %s\n", key, value)
			}

		case "delete":
			if len(args) != 2 {
				fmt.Println("Usage: delete <key>")
				continue
			}
			key := args[1]
			if err := db.DeleteData(database, key); err != nil {
				log.Printf("failed to delete data: %v", err)
			} else {
				fmt.Printf("Successfully deleted key: %s\n", key)
			}

		case "help":
			fmt.Println("Available commands:")
			fmt.Println("  put <key> <value> - Insert or update a key-value pair")
			fmt.Println("  get <key>         - Retrieve the value for a key")
			fmt.Println("  delete <key>      - Delete a key-value pair")
			fmt.Println("  exit              - Exit the CLI")

		default:
			fmt.Println("Unknown command. Type 'help' for available commands.")
		}
	}
}
