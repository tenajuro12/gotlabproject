package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"showmeyourcode/changelog-generator/command"
	"showmeyourcode/changelog-generator/command/cmd_generate_changelog"
	"showmeyourcode/changelog-generator/command/cmd_generate_releasenote"
	"showmeyourcode/changelog-generator/command/cmd_info"
	"showmeyourcode/changelog-generator/constant"
	"showmeyourcode/changelog-generator/path_finder"
	"strconv"
	"strings"
)

var digitsOnlyRegex = regexp.MustCompile(`[^0-9]+`)

// override while running tests
var reader = bufio.NewReader(os.Stdin)
var releaseNotesFileName = constant.ReleaseNotesFileName
var changelogFileName = constant.ChangelogFileName
var commandInfo command.Command = &cmd_info.CommandInfo{}
var commandGenerateChangelog = func(name string) command.Command {
	return &cmd_generate_changelog.CommandGenerateChangelog{FileName: name}
}
var commandReleaseNote = func(name string) command.Command {
	return &cmd_generate_releasenote.CommandGenerateReleaseNote{FileName: name}
}

func main() {
	options := []string{"Release note", "Changelog", "Info"}
	fmt.Printf("Welcome in the documentation tool! Current version: %s\n", constant.Version)
	fmt.Println("If you want to exit the program, please press '0'.")
	fmt.Println("\nSelect an action:")
	for index, option := range options {
		fmt.Printf("%d. %s\n", index+1, option)
	}
	isActionChosen := false
	var sourcePath = ""
	var commandToExecute command.Command = nil
chooseActionLoop:
	for !isActionChosen {
		userChoice, _ := reader.ReadString('\n')
		userChoice = digitsOnlyRegex.ReplaceAllString(strings.TrimSuffix(userChoice, "\n"), "")
		choiceParsed, _ := strconv.Atoi(userChoice)
		switch choiceParsed {
		case 0:
			fmt.Println("Goodbye!")
			break chooseActionLoop
		case 1:
			commandToExecute = commandReleaseNote(releaseNotesFileName)
			sourcePath = path_finder.GetWorkingDirectoryPath()
			isActionChosen = true
		case 2:
			commandToExecute = commandGenerateChangelog(changelogFileName)
			sourcePath = path_finder.GetWorkingDirectoryPath()
			isActionChosen = true
		case 3:
			commandToExecute = commandInfo
			sourcePath = path_finder.GetWorkingDirectoryPath()
			isActionChosen = true
		default:
			panic("Action not supported")
		}
	}
	if commandToExecute != nil {
		commandToExecute.Execute(sourcePath)
		commandToExecute.PrintInfo()
	}
}
