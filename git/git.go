package git

import (
	"os/exec"
)

func GetAllTags() (string, error) {
	out, err := exec.Command("git", "tag").Output()
	return string(out), err
}

func GetTheMostRecentTag() (string, error) {
	out, err := exec.Command("git", "describe", "--tags", "--abbrev=0").Output()
	return string(out), err
}

func GetCommitsSinceTag(tag string) (string, error) {
	out, err := exec.Command("git", "log", tag+"..HEAD", "--pretty=format:'%h %s'.", "--oneline").Output()
	return string(out), err
}

func GetAllCommits() (string, error) {
	out, err := exec.Command("git", "log", "--pretty=format:'%h %s'.", "--oneline").Output()
	return string(out), err
}
