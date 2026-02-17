package review

type ReviewResult struct {
	Issues []Issue `json:"issues"`
}

type Issue struct {
	Line       int    `json:"line"`
	Severity   string `json:"severity"`
	Title      string `json:"title"`
	Suggestion string `json:"suggestion"`
}
