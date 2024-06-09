package custom_io_tool

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFolderContent(t *testing.T) {
	const path = "../examples"
	var folderContent = GetFolderContent(path)
	assert.Equal(t, "examples", folderContent.Name)
	assert.Equal(t, path, folderContent.Path)
	assert.Equal(t, 2, len(folderContent.Subfolders))
}

func TestLoadFileContent(t *testing.T) {
	const path = "../examples/changelogs/template.md"
	var content = LoadFileContent(path)
	assert.NotNil(t, content)
}

func TestSaveFile(t *testing.T) {
	const path = "./test-save-output/test.txt"
	err := SaveFile(path, "Something something\nSomething new line")

	var content = LoadFileContent(path)
	assert.NotNil(t, content)
	assert.Nil(t, err)
}

func TestExists(t *testing.T) {
	const path = "./test-save-output/test.txt"
	exists := Exists(path)

	assert.True(t, exists)
	assert.True(t, Remove(path))
}

func TestRemoveWhenFileExists(t *testing.T) {
	const fileToRemove = "file-to-remove.txt"
	assert.Nil(t, SaveFile(fileToRemove, "Content"))

	isDelete := Remove(fileToRemove)

	assert.True(t, isDelete)
}

func TestDoesNotRemoveWhenFileDoesNotExist(t *testing.T) {
	const path = "./empty-file-never-created.txt"
	isDelete := Remove(path)

	assert.False(t, isDelete)
}
