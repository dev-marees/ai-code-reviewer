package review

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ExtractIssues(aiText string) []Issue {

	var issues []Issue

	for _, l := range strings.Split(aiText, "\n") {

		// Expect format: LINE: 23: message
		if strings.HasPrefix(l, "LINE:") {

			parts := strings.SplitN(l, ":", 3)
			if len(parts) < 3 {
				continue
			}

			// naive parse
			line := 0
			fmt.Sscanf(parts[1], "%d", &line)

			issues = append(issues, Issue{
				Line:  line,
				Title: strings.TrimSpace(parts[2]),
				// Severity and suggestion extraction can be added here if needed
				// For now, we only have the line number and title from the AI response
				Severity:   "medium", // default severity
				Suggestion: "",       // no suggestion extracted in this simple parser
			})
		}
	}

	return issues
}

func ParseResult(raw string) (ReviewResult, error) {

	var r ReviewResult

	err := json.Unmarshal([]byte(raw), &r)
	return r, err
}
