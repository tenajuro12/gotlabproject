package cmd_generate_releasenote

import (
	"github.com/stretchr/testify/assert"
	"showmeyourcode/changelog-generator/constant"
	"showmeyourcode/changelog-generator/custom_io_tool"
	"testing"
)

func TestExecuteReleaseNoteCommandWhenIsValid(t *testing.T) {
	cmdReleaseNote := &CommandGenerateReleaseNote{FileName: constant.ReleaseNotesFileName}
	cmdReleaseNote.Execute("../../")

	assert.True(t, custom_io_tool.Exists(constant.ReleaseNotesFileName), "The test release notes file should exist")
	assert.True(t, custom_io_tool.Remove(constant.ReleaseNotesFileName), "The test file should be removed")
}

func TestPrintInfoReleaseNoteCommandWhenIsValid(t *testing.T) {
	cmdReleaseNote := &CommandGenerateReleaseNote{FileName: constant.ReleaseNotesFileName}
	cmdReleaseNote.PrintInfo()
}
