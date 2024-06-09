package command

import "log"

type Command interface {
	Execute(fileName string)
	PrintInfo()
}

func StopExecutionIfError(err error, message string) {
	if err != nil {
		log.Fatalf("%s Error: %v", message, err)
		// Use `panic(err)` if you don't need to log anything.
	}
}
