package cmd_generate_changelog

import (
	"github.com/stretchr/testify/assert"
	"showmeyourcode/changelog-generator/constant"
	"showmeyourcode/changelog-generator/custom_io_tool"
	"testing"
)

func TestExecuteCommandGenerateChangelogWhenIsValid(t *testing.T) {
	cmdReleaseNote := &CommandGenerateChangelog{FileName: constant.ChangelogFileName}
	cmdReleaseNote.Execute("../../")

	assert.True(t, custom_io_tool.Exists(constant.ChangelogFileName), "The test changelog file should exist")
	assert.True(t, custom_io_tool.Remove(constant.ChangelogFileName), "The test file should be removed")
}

func TestPrintInfoCommandGenerateChangelogWhenIsValid(t *testing.T) {
	cmdReleaseNote := &CommandGenerateChangelog{}
	cmdReleaseNote.PrintInfo()
}
