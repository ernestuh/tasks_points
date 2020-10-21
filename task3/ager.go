package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//fmt.Printf("Args: %#v\n", os.Args)
	err := agerCmd()
	if err != nil {
		os.Exit(1000)
	}
}

func agerCmd() error {
	err := validateArgs()
	if err == nil {
		name := flag.Lookup("n").Value.String()
		age := flag.Lookup("a").Value.String()
		ager(name, age)
	}
	return err
}

func ager(name string, age string) {
	const dockerRunStr = `run|-i|--rm|busybox|sh|-c|echo Hi #arg1, age: #arg2;sleep #arg2;`
	const appName = "docker"

	// preparing the docker run command
	dockerRunCmd := strings.Replace(strings.Replace(dockerRunStr, "#arg1", name, 1), "#arg2", age, 2)
	appArg := strings.Split(dockerRunCmd, "|")
	fmt.Printf("Command to execute: %s %s\n", appName, strings.Join(appArg, " "))

	// executing docker run command
	out, err := exec.Command(appName, appArg...).CombinedOutput()

	if err == nil {
		fmt.Printf("%s", out)
	} else {
		fmt.Printf("Command Error: %s\n", err)
		os.Exit(1000)
	}
}

func validateArgs() error {
	// set parameters spec
	flag.String("n", "", "The name of the executor (required)")
	flag.Int("a", 0, "The age of the executor (required)")
	// check required parameters
	if len(os.Args[1:]) < 2 {
		return agerCmdHelpWithError("Invalid number of arguments")
	}

	flag.Parse()

	// extra validations
	nameValue := flag.Lookup("n").Value.String()
	if len(nameValue) == 0 {
		return agerCmdHelpWithError("-n value could not be empty")
	}
	// check if name is alphabetic
	var isStringAlphabetic = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	if !isStringAlphabetic(nameValue) {
		return agerCmdHelpWithError("-n value must be a alphabetic")
	}
	ageValue, _ := strconv.Atoi(flag.Lookup("a").Value.String())
	if ageValue == 0 {
		return agerCmdHelpWithError("-a value must be a greater than zero")
	}

	return nil
}

func agerCmdHelpWithError(err string) error {
	errStr := "Error: " + err
	println(errStr)
	agerCmdHelp()
	return errors.New(errStr)
}

func agerCmdHelp() {
	fmt.Printf("\nUsage: ager -n value -a value\n")
	fmt.Printf("\nAger: A simple program\n\n")
	fmt.Printf("Options:\n")
	flag.PrintDefaults()
}