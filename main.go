package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Headers struct {
	Value    int
	Income   float64
	Age      int
	Rooms    int
	Bedrooms int
	Pop      int
	HH       int
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

func writeFile(data Headers, outfile string) {
	file, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
	}

	f, err := os.OpenFile(outfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer f.Close()

	if _, err := f.Write(file); err != nil {
		log.Fatal("Error writing to file:", err)
	}

	if _, err := f.WriteString("\n"); err != nil {
		log.Fatal("Error writing newline to file:", err)
	}
}

func processLine(input string, outfile string) {
	splitInput := strings.Split(input, ",")
	var curHeader Headers
	value, err := strconv.Atoi(splitInput[0])
	if err != nil {
		value2, err2 := strconv.ParseFloat(splitInput[0], 64)
		if err2 != nil {
			log.Fatal("Could not interpret input within first field", err)
		} else {
			value = int(value2)
		}
	}
	curHeader.Value = value

	income, err := strconv.ParseFloat(splitInput[1], 64)
	if err != nil {
		panic(err)
	}
	curHeader.Income = income

	curHeader.Age, err = strconv.Atoi(splitInput[2])
	if err != nil {
		panic(err)
	}
	curHeader.Rooms, err = strconv.Atoi(splitInput[3])
	if err != nil {
		panic(err)
	}
	curHeader.Bedrooms, err = strconv.Atoi(splitInput[4])
	if err != nil {
		panic(err)
	}
	curHeader.Pop, err = strconv.Atoi(splitInput[5])
	if err != nil {
		panic(err)
	}
	curHeader.HH, err = strconv.Atoi(splitInput[6])
	if err != nil {
		panic(err)
	}
	writeFile(curHeader, outfile)
}

func processCSV(inputFileName string, outputFileName string) (result bool) {
	result = false
	infile, inerr := os.Open(inputFileName)
	if inerr != nil {
		log.Fatal("Could not open input file", inerr)
		return
	}

	outfile, outerr := os.Open(outputFileName)
	if outerr != nil {
		log.Fatal("Could not open output file", outerr)
		defer infile.Close()
		return
	}

	fileScanner := bufio.NewScanner(infile)
	fileScanner.Split(bufio.ScanLines)

	var idx int = 0
	for fileScanner.Scan() {
		if idx == 0 {
			idx++
			continue
		} else {
			var line string = fileScanner.Text()
			processLine(line, outputFileName)
		}

	}

	defer infile.Close()
	defer outfile.Close()

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
