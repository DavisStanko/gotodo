package main

import (
	"bufio"   // For reading input from the user
	"fmt"     // For printing output
	"os"      // For interacting with the operating system (stdin)
	"strings" // For trimming input and handling text
)

// Struct to store each to-do item.
type Task struct {
    Title string
}

func main() {
    // A slice (like a dynamic array) to hold tasks.
    tasks := []Task{}

    // Scanner to read input from the user line by line.
    scanner := bufio.NewScanner(os.Stdin)

    // Main loop 
    for {
        fmt.Println("\nCommands: add, list, quit")
        fmt.Print("> ")

        // Read input
        scanner.Scan()
        input := strings.TrimSpace(scanner.Text()) // Clean whitespace

        // Handle input
        switch input {
        case "add":
            // Ask the user for the task description
            fmt.Print("Enter task: ")
            scanner.Scan()
            title := scanner.Text()

            // Add the new task to our list
            tasks = append(tasks, Task{Title: title})

        case "list":
            // Go through each task and print it out
            for i, task := range tasks {
                fmt.Printf("[%d] %s\n", i+1, task.Title)
            }

        case "quit":
            // Exit 
            fmt.Println("Goodbye!")
            return

        default:
            fmt.Println("Unknown command")
        }
    }
}
