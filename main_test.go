package main

import (
	"reflect"
	"testing"
)

func TestValidateArgs(t *testing.T) {
	cases := []struct {
		name     string
		args     []string
		expected bool
	}{
		{"No Args", []string{}, false},
		{"One Arg", []string{"housesInput.csv"}, false},
		{"Two Args", []string{"housesInput.csv", "output.json"}, true},
		{"Invalid CSV File", []string{"input.txt", "output.json"}, false},
		{"Invalid JSON File", []string{"housesInput.csv", "output.txt"}, false},
		{"Correct Args", []string{"housesInput.csv", "output.json"}, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := validateArgs(c.args); got != c.expected {
				t.Errorf("validateArgs(%v) = %v; want %v", c.args, got, c.expected)
			}
		})
	}
}

func TestParseLine(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected Headers
		wantErr  bool
	}{
		{"Valid Input", "1,50000,25,5,3,1000,4", Headers{1, 50000, 25, 5, 3, 1000, 4}, false},
		{"Invalid Value", "abc,50000,25,5,3,1000,4", Headers{}, true},
		{"Invalid Value Last Item", "10000,50000,25,5,3,1000,A", Headers{}, true},
		{"Incomplete Data", "1,50000", Headers{}, true},
		{"Too Much Data", "1,50000, 50000, 4500, 2100, 325, 45748", Headers{}, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := parseLine(c.input)
			if (err != nil) != c.wantErr {
				t.Errorf("parseLine() error = %v, wantErr %v", err, c.wantErr)
				return
			}
			// DeepEqual used here to ensure that empty Header structs are returned
			if !reflect.DeepEqual(got, c.expected) {
				t.Errorf("parseLine() = %v, expected %v", got, c.expected)
			}
		})
	}
}
