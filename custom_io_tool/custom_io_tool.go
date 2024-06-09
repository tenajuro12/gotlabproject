package custom_io_tool

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"showmeyourcode/changelog-generator/command"
	"showmeyourcode/changelog-generator/path_finder"
)

type Directory struct {
	Name       string
	Path       string
	Files      []string
	Subfolders []Directory
}

func GetFolderContent(path string) Directory {
	var _, lastPathElem = path_finder.GetPathLastElement(path)
	var result = Directory{
		lastPathElem,
		path,
		[]string{},
		[]Directory{},
	}
	files, err := ioutil.ReadDir(path)
	command.StopExecutionIfError(err, fmt.Sprintf("Cannot read the directory content. Path: %s.", path))

	for _, f := range files {
		if f.IsDir() {
			fmt.Println("Processing folder: " + f.Name())
			result.Subfolders = append(result.Subfolders, GetFolderContent(path_finder.CombinePath(path, f.Name())))
		} else {
			fmt.Println("Found file: " + f.Name())
			result.Files = append(result.Files, f.Name())
		}
	}
	return result
}

func LoadFileContent(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	command.StopExecutionIfError(err, fmt.Sprintf("Cannot read file content. Path: %s.", path))

	return string(b)
}

func SaveFile(path string, fileContent string) error {
	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("Cannot create a file: %s. Error: %s\n", path, err.Error())
		return err
	}
	l, err := f.WriteString(fileContent)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return err
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Exists(path string) bool {
	f, err := os.Open(path)
	closeErr := f.Close()
	if closeErr != nil {
		fmt.Printf("\nCannot close the file: %s. Err: %s", path, closeErr.Error())
	}
	return err == nil
}

func Remove(fileName string) bool {
	err := os.Remove(fileName)
	isDeleted := err == nil
	if !isDeleted {
		fmt.Println("Cannot remove the file. Error: " + err.Error())
	}
	return isDeleted
}
