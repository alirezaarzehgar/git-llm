package commit

import (
	"fmt"

	"github.com/alirezaarzehgar/git-llm/internal/llm"
)

func FixCommitMessage(langModel llm.LanguageModel, message string) error {
	commitMessage, err := langModel.FixCommit(message)
	if err != nil {
		return fmt.Errorf("failed to connect LLM: %w", err)
	}

	err = gitCommit(commitMessage)
	if err != nil {
		return fmt.Errorf("failed to commit: %w", err)
	}

	return nil
}
