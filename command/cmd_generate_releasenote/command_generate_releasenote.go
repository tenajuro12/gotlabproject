package cmd_generate_releasenote

import (
	"bytes"
	"fmt"
	"gopkg.in/loremipsum.v1"
	"showmeyourcode/changelog-generator/command"
	"showmeyourcode/changelog-generator/constant"
	"showmeyourcode/changelog-generator/custom_io_tool"
	"time"
)

type CommandGenerateReleaseNote struct {
	FileName string
}

func (c *CommandGenerateReleaseNote) Execute(workingDirectory string) {
	fmt.Println("Executing a command with source: " + workingDirectory)

	fmt.Println("### The command is going to replace the old release notes file completely ###")

	var buffer bytes.Buffer
	var loremipsumGenerator = loremipsum.New()

	buffer.WriteString(constant.ReleaseNotestHeader + "\n\n")

	buffer.WriteString(loremipsumGenerator.Sentence() + "\n\n")

	buffer.WriteString(fmt.Sprintf("## ?.?.? (%s)\n\n", time.Now().Format("2006-1-2")))

	buffer.WriteString(loremipsumGenerator.Paragraphs(3) + "\n\n")

	err := custom_io_tool.SaveFile(c.FileName, buffer.String())
	if err != nil {
		command.StopExecutionIfError(err, "Cannot save the release notes.")
	}
}

func (c *CommandGenerateReleaseNote) PrintInfo() {
	fmt.Println("Finish preparing the latest release notes document!")
}
