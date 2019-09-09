package main

import (
	"fmt"
	"io/ioutil"
	"os"

	flag "github.com/spf13/pflag"
)

func main() {
	var argsFilePath *string = flag.String("args-file", "", "")
	var stdinFilePath *string = flag.String("stdin-file", "", "")
	var output *string = flag.String("output", "", "")
	var killGardenServer *bool = flag.Bool("kill-garden-server", false, "")
	var action *string = flag.String("action", "", "")
	var handle *string = flag.String("handle", "", "")

	flag.Parse()

	if *argsFilePath != "" {
		argsFile, err := os.OpenFile(*argsFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}
		defer argsFile.Close()
		if _, err := fmt.Fprintf(argsFile, "--action %s --handle %s\n", *action, *handle); err != nil {
			panic(err)
		}
	}

	if *stdinFilePath != "" {
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile(*stdinFilePath, input, 0600); err != nil {
			panic(err)
		}
	}

	if *killGardenServer && *action == "down" {
		proc, err := os.FindProcess(os.Getppid())
		if err != nil {
			panic(err)
		}

		if err := proc.Kill(); err != nil {
			panic(err)
		}
	}

	if *output != "" {
		fmt.Println(*output)
	}
}
