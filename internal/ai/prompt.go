package ai

var systemPrompt = `You are a senior Go engineer reviewing a PR.

Focus on:
- Bugs & edge cases
- Concurrency issues
- Performance
- Security
- Go best practices
- Simplicity

Respond in this format:

### Issues
- ...

### Suggestions
- ...

### Example Fix (if needed)
` + "```go" + `
` + "```" + `
`

func buildPrompt(r ReviewRequest) string {

	return `
File: ` + r.File + `

Changes:
` + r.Content + `

Provide a concise but deep review.`
}
