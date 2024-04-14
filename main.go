package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Builder struct {
}

func printHelp() {
	fmt.Println("\nHelp Page:")
	fmt.Println("This tool is designed to convert cvs files to JSON format")
	fmt.Println("\nUsage:")
	fmt.Println("main.exe <first arguement> <second arguement")
	fmt.Println("The first argument should be your input in csv format")
	fmt.Println("The second argument should be your output filename ending in .json")
	fmt.Println("Example: ./main.exe inputfile.csv outputfile.json")
	fmt.Println("")
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

func inputFileCheck(inputFileName string, outputFileName string) (result bool) {
	result = false
	infile, inerr := os.Open(inputFileName)
	if inerr != nil {
		log.Fatal("Could not open input file", inerr)
		return
	}

	outfile, outerr := os.Create(outputFileName)
	if outerr != nil {
		log.Fatal("Could not open/create output file", outerr)
		defer infile.Close()
		return
	} else {
		defer infile.Close()
		defer outfile.Close()
		result = true
		return
	}
}

func establishFields(header string, lineOne string) {
	headerlist := NewGenericList[any]
	// firstline := NewGenericList[any]

	headerFields := strings.Split(header, ",")
	for _, title := range headerFields {
		headerlist().Insert(title)
	}

	for idx := range headerFields {
		fmt.Println(idx)
		fmt.Println(headerlist().ValueByIndex(idx))
	}

	firstLine := strings.Split(lineOne, ",")
	for _, value := range firstLine {
		fmt.Println(value)
	}
}

func processCSV(inputFileName string, outputfileName string) (result bool) {
	result = false
	infile, inerr := os.Open(inputFileName)
	if inerr != nil {
		log.Fatal("Could not open input file", inerr)
		return
	}

	fileScanner := bufio.NewScanner(infile)
	fileScanner.Split(bufio.ScanLines)

	var firsttwo [2]string
	var idx int = 0
	for fileScanner.Scan() {
		firsttwo[idx] = fileScanner.Text()
		idx++
		if idx >= 2 {
			break
		}
	}
	establishFields(firsttwo[0], firsttwo[1])
	fmt.Println(len(firsttwo[0]))
	fmt.Println(len(firsttwo[1]))

	fmt.Println(outputfileName)
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
	inputFileCheck(inputFileName, outputFileName)
	processCSV(inputFileName, outputFileName)
}
