package main

import (
	"github.com/jamesandariese/issh"
	"os"
)

func main() {
	key, err := issh.GetAuthorizedKey("fred")
	if err != nil {
		panic("Failed to create a key: " + err.Error())
	}
	os.Stdout.Write(key)
	stdout, exitcode, err := issh.Run("james", "localhost", 22, "fred")
	if err != nil {
		panic("Failed to execute remote command: " + err.Error())
	}
	
	os.Stdout.Write(stdout)
	os.Exit(exitcode)
}
