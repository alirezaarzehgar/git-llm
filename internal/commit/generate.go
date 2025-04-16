package commit

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/alirezaarzehgar/git-llm/internal/llm"
	"github.com/spf13/viper"
)

var defaultEditor = "nano"

func Generate(langModel llm.LanguageModel) error {
	diff, err := getCachedDiff()
	if err != nil {
		return err
	}

	commitMessage, err := langModel.GenerateCommitByDiff(diff)
	if err != nil {
		return fmt.Errorf("failed to connect LLM: %w", err)
	}

	commitMessage, err = getCommitFromEditor(commitMessage)
	if err != nil {
		return err
	}

	err = gitCommit(commitMessage)
	if err != nil {
		return err
	}

	return nil
}

func getCachedDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
	diff, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to run git diff: %w", err)
	}
	return string(diff), nil
}

func getCommitFromEditor(commitMessage string) (string, error) {
	f, err := os.CreateTemp("", "git-commit-*.txt")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}
	tempFileName := f.Name()
	f.WriteString(commitMessage)
	f.Close()
	defer os.Remove(tempFileName)

	editor := defaultEditor
	if ed := viper.GetString("EDITOR"); ed != "" {
		editor = ed
	}

	cmd := exec.Command(editor, tempFileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to open editor: %w", err)
	}

	data, err := os.ReadFile(tempFileName)
	if err != nil {
		return "", fmt.Errorf("failed to read temp file: %w", err)
	}

	return string(data), nil
}

func gitCommit(msg string) error {
	cmd := exec.Command("git", "commit", "-s", "-m", msg)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to commit: %w", err)
	}
	return nil
}
