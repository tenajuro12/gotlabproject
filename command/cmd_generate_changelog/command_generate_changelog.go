package cmd_generate_changelog

import (
	"bytes"
	"fmt"
	"showmeyourcode/changelog-generator/command"
	"showmeyourcode/changelog-generator/constant"
	"showmeyourcode/changelog-generator/custom_io_tool"
	"showmeyourcode/changelog-generator/git"
	"strings"
	"time"
)

type GitCommit struct {
	hash    string
	message string
}

type CommandGenerateChangelog struct {
	FileName string
}

func (c *CommandGenerateChangelog) Execute(workingDirectory string) {
	fmt.Println("Executing a command with source: " + workingDirectory)

	tagsAsString, errTags := git.GetAllTags()
	tags := removeEmptyStrings(splitByNewLine(tagsAsString))

	numberOfTags := len(tags)
	areGitTagsPresent := numberOfTags > 0

	if errTags != nil {
		fmt.Printf("An error occured when fetching tags. Details: %s", errTags.Error())
	}

	var commitsAsString string
	var commits []GitCommit
	var errCommits error

	if areGitTagsPresent {
		fmt.Printf("Found %d Git tags.\n", numberOfTags)
		tag, errTag := git.GetTheMostRecentTag()

		if errTag != nil {
			fmt.Printf("An error occured when fetching the most recent tag. Details: %s\n", errTag.Error())
		}

		fmt.Printf("The latest Git tag: %s", tag)

		commitsAsString, errCommits = git.GetCommitsSinceTag(tag)
		commits = parseCommits(splitByNewLine(commitsAsString))
	} else {
		commitsAsString, errCommits = git.GetAllCommits()
		commits = parseCommits(splitByNewLine(commitsAsString))
	}

	if errCommits != nil {
		fmt.Printf("An error occured when fetching commits. Details: %s\n", errCommits.Error())
	} else {
		appendToChangelogFile(commits, c.FileName)
	}
}

func (c *CommandGenerateChangelog) PrintInfo() {
	fmt.Println("Finish preparing the latest changelog document.")
}

func appendToChangelogFile(commits []GitCommit, fileName string) {
	fmt.Println("Processing the changelog file. File name: " + fileName)

	buffer := prepareNewChangelogRelease(commits)

	if custom_io_tool.Exists(fileName) {
		fmt.Printf("Found %s. Appending changes.\n", fileName)
		buffer.WriteString(strings.ReplaceAll(custom_io_tool.LoadFileContent(fileName), constant.ChangelogHeader, ""))
	}

	err := custom_io_tool.SaveFile(fileName, buffer.String())
	if err != nil {
		command.StopExecutionIfError(err, "Cannot save the changelog.")
	}
}

func prepareNewChangelogRelease(commits []GitCommit) bytes.Buffer {
	var buffer bytes.Buffer

	buffer.WriteString(constant.ChangelogHeader + "\n\n")

	buffer.WriteString(fmt.Sprintf("## ?.?.? (%s)\n\n", time.Now().Format("2006-1-2")))

	for _, commit := range commits {
		buffer.WriteString(fmt.Sprintf("- %s (%s)\n", commit.message, commit.hash))
	}

	buffer.WriteString("\n\n")

	return buffer
}

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func splitByNewLine(stringToSplit string) []string {
	return strings.Split(stringToSplit, "\n")
}

func parseCommits(commits []string) []GitCommit {
	commits = removeEmptyStrings(commits)
	parsedCommits := make([]GitCommit, len(commits))
	for i, v := range commits {
		splitBySpace := strings.Split(v, " ")
		parsedCommits[i] = GitCommit{
			hash:    splitBySpace[0],
			message: strings.Join(splitBySpace[1:], " "),
		}
	}
	return parsedCommits
}
