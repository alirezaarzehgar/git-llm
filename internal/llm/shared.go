package llm

type LanguageModel interface {
	GenerateCommitByDiff(diff string) (string, error)
	FixComment(commentMessage string) (string, error)
}
