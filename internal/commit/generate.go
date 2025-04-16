package commit

import (
	"fmt"
	"os/exec"
)

var Editor = "vim"

func Generate() error {
	cmd := exec.Command("git", "diff", "--cached")
	diff, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to run git diff: %w", err)
	}

	fmt.Println(diff)
	return err
}
