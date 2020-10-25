package main

import (
	"bytes"
	"flag"
	"os"
	"testing"
)

func TestValidNoArguments(t *testing.T) {
	// Test no arguments
	argHelper("", "")

	err := validateArgs()
	if err == nil {
		t.Fail()
	}
}

func TestValidateName(t *testing.T) {
	// Test invalid name
	argHelper("Ernest1", "5")

	err := validateArgs()
	if err == nil {
		t.Fail()
	}
}

func TestValidateAge(t *testing.T) {
	// Test age 0
	argHelper("Ernest", "0")

	err := validateArgs()
	if err == nil {
		t.Fail()
	}
}

func TestArgCmdSucc(t *testing.T) {
	// Test age 0
	argHelper("Devops", "3")

	cmdOutput, err := agerCmd()

	// agerCmd execute without error
	if err != nil {
		t.Fail()
	}

	// Ensure output is correct
	if !bytes.Contains(cmdOutput, []byte("Hi Devops, age: 3")) {
		t.Fail()
	}
}

func argHelper(name string, age string) {
	os.Args = []string{"ager", "-n", name, "-a", age}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
}
