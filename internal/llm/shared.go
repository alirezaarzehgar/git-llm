package llm

type LanguageModel interface {
	GenerateCommitByDiff(diff string) (string, error)
	FixCommit(commitMessage string) (string, error)
}
