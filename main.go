package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	// Initialize variables
	var names []string
	var attendance []bool
	var presentNames []string
	var absentNames []string
	var coordinators []string
	content := ""
	var batchName string

	// Open and read the members file
	file, err := os.Open("./members.txt")
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Read names from the file
	for scanner.Scan() {
		name := strings.TrimSpace(scanner.Text())
		names = append(names, name)
	}

	// Open and read the content file
	fileContent, err := os.ReadFile("./content.txt")
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	content = string(fileContent)

	// Prompt for the batch name
	fmt.Print("Enter the batch name: ")
	fmt.Scanln(&batchName)

	// Prompt for the current date or an optional selected date
	var selectedDate string
	fmt.Println("Enter today's date (or press enter to use the current date):")
	fmt.Scanln(&selectedDate)

	// If the user does not input anything, use the current date
	if selectedDate == "" {
		selectedDate = time.Now().Format("2006-01-02")
	}

	// Ask the user to select coordinators from the list of members
	fmt.Println("Select coordinators from the list below (press Enter when done):")
	for i, name := range names {
		fmt.Printf("%d. %s\n", i+1, name)
	}

	// Collect multiple coordinators
	for {
		var coordinatorIndex int
		fmt.Print("Enter the number of the coordinator (or press Enter to stop): ")
		// Check if input is a valid number
		_, err := fmt.Scanln(&coordinatorIndex)
		if err != nil {
			break // Break the loop if the user presses Enter without entering a number
		}

		// Validate the coordinator index and add to the list
		if coordinatorIndex > 0 && coordinatorIndex <= len(names) {
			coordinators = append(coordinators, names[coordinatorIndex-1])
		} else {
			fmt.Println("Invalid selection. Please choose a valid number.")
		}
	}

	// Ask for attendance for each member
	for _, name := range names {
		var isPresent string
		fmt.Printf("Is %s present today?: ", name)
		fmt.Scanln(&isPresent)

		// Treat empty input as "present"
		if isPresent == "" || strings.ToLower(isPresent) == "y" {
			attendance = append(attendance, true)
		} else {
			attendance = append(attendance, false)
		}
	}

	// Sort the names alphabetically
	sort.Strings(names)

	// Collect present and absentee names
	for i, present := range attendance {
		if present {
			presentNames = append(presentNames, names[i])
		} else {
			absentNames = append(absentNames, names[i])
		}
	}

	// Sort the present and absentee names as well
	sort.Strings(presentNames)
	sort.Strings(absentNames)

	// Output the formatted results
	fmt.Println("\nSession Report for", selectedDate)
	fmt.Println("------------------------------------------------")

	// Print the batch name and coordinators
	fmt.Println("Batch Name:", batchName)
	if len(coordinators) > 0 {
		fmt.Println("Coordinators: " + strings.Join(coordinators, ", "))
	} else {
		fmt.Println("No coordinators selected.")
	}
	fmt.Println("------------------------------------------------")

	// Print content after the content header
	fmt.Println(content)
	fmt.Println("------------------------------------------------")

	// Print present members
	fmt.Println("Present Members:")
	for i, name := range presentNames {
		fmt.Printf("%d. %s\n", i+1, name)
	}

	// Print absentees (if any)
	if len(absentNames) > 0 {
		fmt.Println("\nAbsentees:")
		for i, name := range absentNames {
			fmt.Printf("%d. %s\n", i+1, name)
		}
	} else {
		fmt.Println("\nNo absentees today!")
	}
}
