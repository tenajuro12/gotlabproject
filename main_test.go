package main

import (
	"bufio"
	. "github.com/ovechkin-dm/mockio/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"showmeyourcode/changelog-generator/command"
	"showmeyourcode/changelog-generator/constant"
	"showmeyourcode/changelog-generator/custom_io_tool"
	"strings"
	"testing"
)

//
// INTEGRATION TESTS
//

func TestProgramExistsWithoutErrorWhen0IsPressed(t *testing.T) {
	reader = bufio.NewReaderSize(strings.NewReader("0"), 25)
	main()
}

func TestProgramExistsWithoutErrorWhenNonDigitCharactersArePressed(t *testing.T) {
	reader = bufio.NewReaderSize(strings.NewReader("rewrtegfds  oi"), 25)
	main()
}

func TestProgramExistsWithErrorWhenNonSupportedActionIsPressed(t *testing.T) {
	reader = bufio.NewReaderSize(strings.NewReader("5"), 25)
	require.Panics(t, func() { main() })
}

func TestProgramCreatesReleaseNotesWhen1IsPressed(t *testing.T) {
	releaseNotesFileName = "TEST-" + constant.ReleaseNotesFileName
	reader = bufio.NewReaderSize(strings.NewReader("1"), 25)

	main()

	assert.True(t, custom_io_tool.Exists(releaseNotesFileName), "The test release notes file should exist")
	assert.True(t, custom_io_tool.Remove(releaseNotesFileName), "The test file should be removed")
}

func TestProgramPrintsInfoWhen3IsPressed(t *testing.T) {
	SetUp(t)
	commandInfo = Mock[command.Command]()
	reader = bufio.NewReaderSize(strings.NewReader("3"), 25)

	main()

	Verify(commandInfo, Once()).Execute(Any[string]())
}

func TestUserInputIsFilteredWhenNonDigitsAreChosen(t *testing.T) {
	SetUp(t)
	commandInfo = Mock[command.Command]()
	reader = bufio.NewReaderSize(strings.NewReader("waergdfas 3asd asd"), 100)

	main()

	Verify(commandInfo, Once()).Execute(Any[string]())
}

func TestProgramCreatesChangelogWhen2IsPressed(t *testing.T) {
	changelogFileName = "TEST-" + constant.ChangelogFileName
	reader = bufio.NewReaderSize(strings.NewReader("2"), 25)

	main()

	assert.True(t, custom_io_tool.Exists(changelogFileName), "The test changelog file should exist")
	assert.True(t, custom_io_tool.Remove(changelogFileName), "The test file should be removed")
}
