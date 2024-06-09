package cmd_info

import (
	"fmt"
)

type CommandInfo struct {
}

func (c *CommandInfo) Execute(workingDirectory string) {
	fmt.Printf("\nPROGRAM INFO\nWorking directory: %s\n\n", workingDirectory)
}

func (c *CommandInfo) PrintInfo() {

}
