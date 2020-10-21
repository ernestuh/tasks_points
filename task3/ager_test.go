package main

import (
	"flag"
	"os"
	"testing"
)

var oldArgs []string
var err error

func TestValidNoArguments(t *testing.T) {
	defer func() { os.Args = oldArgs }()

	// Test no arguments
	argHelper("", "")

	err = validateArgs()
	if err == nil {
		t.Fail() // .Errorf("-n argument must be validate to alphabetic value")
	}
}

func TestValidateName(t *testing.T) {
	defer func() { os.Args = oldArgs }()

	// Test invalid name
	argHelper("Ernest1", "5")

	err = validateArgs()
	if err == nil {
		t.Fail() // .Errorf("-n argument must be validate to alphabetic value")
	}
}

func TestValidateAge(t *testing.T) {
	defer func() { os.Args = oldArgs }()

	// Test age 0
	argHelper("Ernest", "0")

	err := validateArgs()
	if err == nil {
		t.Fail() // .Errorf("-n argument must be validate to alphabetic value")
	}
}

func argHelper(name string, age string) {
	os.Args = []string{"ager", "-n", name, "-a", age}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
}
