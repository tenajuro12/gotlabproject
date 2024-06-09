package path_finder

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCombinePath(t *testing.T) {
	var releaseNoteFolderOutputPath = CombinePath("folder1", "folder2")
	assert.True(t, strings.Contains(releaseNoteFolderOutputPath, "folder2"))
}

func TestGetWorkingDirectoryPath(t *testing.T) {
	var workingDirectoryPath = GetWorkingDirectoryPath()
	assert.True(t, strings.Contains(workingDirectoryPath, "path_finder"))
}
