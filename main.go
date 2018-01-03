package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"math/rand"
  	"time"
)

func main() {
	printStartingMessage()

	// Loops forever until the user chooses the "exit" option
	var userOption = ""
	for {
		userOption = getUserOption()
		calculateResult(userOption)
	}
}

func printStartingMessage() {
	// Print instructions of the game
	fmt.Println("Welcome to Rock–Paper–Scissors!")
	fmt.Println("  To choose Rock, enter \"Rock\" or \"R\" in the terminal")
	fmt.Println("  To choose Paper, enter \"Paper\" or \"P\" in the terminal")
	fmt.Println("  To choose Scissors, enter \"Scissors\" or \"S\" in the terminal")
	fmt.Println("  To Exit the program, enter \"Exit\" or \"E\" in the terminal")
	fmt.Println("")
}

func getUserOption() string {
	// Get input from the command line
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your option: ")
	text, _ := reader.ReadString('\n')
	return text
}

func getComputerOption() string {
	// Generate a random choice from the three options of the game
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	options := []string{"Rock", "Paper", "Scissors"}
	randomInteger := r.Intn(3)
	return options[randomInteger]
}

func calculateResult(userOption string) {
	// With the first character entered by the user
	// we use a switch statement to determine what to do with the input
	firstCharacter := userOption[:1]
	switch strings.ToUpper(firstCharacter) {
	case "R":
		whoWon("Rock", getComputerOption())
	case "P":
		whoWon("Paper", getComputerOption())
	case "S":
		whoWon("Scissors", getComputerOption())
	case "E":
		fmt.Println("Goodbye.")
		os.Exit(0)
	default:
		fmt.Println("INVALID OPTION!")
	}
}

func whoWon(userOption string, computerOption string) {
	// options is a map made of maps, the outer maps have the name of the 9
	// possibilities in Rock-Paper-Scissors.
	// The inner maps have a simple ascii art of the possibility and the result of it.

	options := map[string]map[string]string{
		"Rock_vs_Rock": map[string]string{
			"ascii_art": `
    _______             _______
---'   ____)           (____   '---ma
      (_____)         (_____)
      (_____)   Vs.   (_____)
      (____)           (____)
---.__(___)             (___)__.---
`,
			"result": "It is a Tie!",
		},
		"Rock_vs_Paper": map[string]string{
			"ascii_art": `
    _______                  _______
---'   ____)            ____(____   '---
      (_____)          (______
      (_____)   Vs.   (_______
      (____)           (_______
---.__(___)              (__________.---
`,
			"result": "You Lost!",
		},
		"Rock_vs_Scissors": map[string]string{
			"ascii_art": `
    _______                  ______
---'   ____)            ____(____  '---
      (_____)          (______
      (_____)   Vs.   (_______
      (____)               (____)
---.__(___)                 (___)__.---
`,
			"result": "You Won!",
		},
		"Paper_vs_Rock": map[string]string{
			"ascii_art": `
    _______                  _______
---'   ____)____            (____   '---
          ______)          (_____)
          _______)   Vs.   (_____)
         _______)           (____)
---.__________)              (___)__.---
`,
			"result": "You Won!",
		},
		"Paper_vs_Paper": map[string]string{
			"ascii_art": `
    _______                       _______
---'   ____)____             ____(____  '---
          ______)           (______
          _______)   Vs.   (_______
         _______)           (_______
---.__________)              (__________.---
`,
			"result": "It is a Tie!",
		},
		"Paper_vs_Scissors": map[string]string{
			"ascii_art": `
    _______                       _______
---'   ____)____             ____(____  '---
          ______)           (______
          _______)   Vs.   (_______
         _______)               (____)
---.__________)                  (___)__.---
`,
			"result": "You Lost!",
		},
		"Scissors_vs_Rock": map[string]string{
			"ascii_art": `
    _______                  _______
---'   ____)____            (____   '---
          ______)          (_____)
          _______)   Vs.   (_____)
    	(____)              (____)
---.__(___)                  (___)__.---
`,
			"result": "You Lost!",
		},
		"Scissors_vs_Paper": map[string]string{
			"ascii_art": `
    _______                       _______
---'   ____)____             ____(____  '---
          ______)           (______
          _______)   Vs.   (_______
    	(____)              (_______
---.__(___)                  (__________.---
`,
			"result": "You Won!",
		},
		"Scissors_vs_Scissors": map[string]string{
			"ascii_art": `
    _______                       _______
---'   ____)____             ____(____   '---
          ______)           (______
          _______)   Vs.   (_______
    	(____)                (____)
---.__(___)                      (___)__.---
`,
			"result": "It is a Tie!",
		},
	}

	// Display the options choosed by the user and the computer, 
	// an ascii representation of the game and the result of the interaction.
	fmt.Printf("You choose %v. The computer choose %v.", userOption, computerOption)
	fmt.Println(options[userOption + "_vs_" + computerOption]["ascii_art"])
	fmt.Println(options[userOption + "_vs_" + computerOption]["result"])
	fmt.Println("")
}

/*
What did I learn with this Go project?
	For Loops
	Switch Statements
	Multiline Strings
	Nested Maps
	Generating Random Numbers
	Getting Input from The Terminal
*/