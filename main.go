package main

import (
	"fmt"
	"os"
	"strings"
)

func printHelp() {
	fmt.Println("\nHelp Page:")
	fmt.Println("This tool is designed to convert cvs files to JSON format")
	fmt.Println("\nUsage:")
	fmt.Println("main.exe <first arguement> <second arguement")
	fmt.Println("The first argument should be your input in csv format")
	fmt.Println("The second argument should be your output filename ending in .json")
	fmt.Println("Example: ./main.exe inputfile.csv outputfile.json\n")

}

func validateArgs(args []string) (result bool) {
	result = false
	if len(args) != 2 {
		printHelp()
		return
	}
	if len(args[0]) < 4 || len(args[1]) <= 5 {
		printHelp()
		return
	}
	if !strings.HasSuffix(args[0], ".csv") || !strings.HasSuffix(args[1], ".json") {
		printHelp()
		return
	}
	result = true
	return
}

func main() {
	var validInput bool = validateArgs(os.Args[1:])
	if !validInput {
		return
	}
	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

}