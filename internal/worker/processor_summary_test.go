package worker

import (
	"context"
	"strings"
	"testing"

	"ai-code-reviewer/internal/config"
	"ai-code-reviewer/internal/dedup"
	"ai-code-reviewer/internal/github"
	"ai-code-reviewer/internal/mocks"
	"ai-code-reviewer/internal/observability"
	"ai-code-reviewer/internal/ratelimit"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type clientStub struct {
	files []github.PRFile
}

func (c *clientStub) GetPRFiles(ctx context.Context, repo string, pr int) ([]github.PRFile, error) {
	return c.files, nil
}

func (c *clientStub) GetPRDiff(ctx context.Context, repo string, pr int) (string, error) {
	return "", nil
}

func (c *clientStub) CreateComment(ctx context.Context, repo string, pr int, body string) error {
	return nil
}

func (c *clientStub) CreateLineComment(ctx context.Context, repo string, pr int, comment github.LineComment) error {
	return nil
}

func TestFormatSummaryComment_NoIssues(t *testing.T) {
	body := formatSummaryComment(reviewSummary{
		TotalIssues:      0,
		PostedComments:   0,
		SeverityCounters: map[string]int{"critical": 0, "high": 0, "medium": 0, "low": 0},
	})

	require.Contains(t, body, "No issues detected")
}

func TestProcessorHandle_PostsSummaryComment(t *testing.T) {
	ai := mocks.NewProvider(t)
	comments := mocks.NewCommentClient(t)
	client := &clientStub{
		files: []github.PRFile{
			{
				Filename: "main.go",
				Patch: "diff --git a/main.go b/main.go\n" +
					"--- a/main.go\n" +
					"+++ b/main.go\n" +
					"@@ -1,1 +1,2 @@\n" +
					"-old\n" +
					"+new\n",
			},
		},
	}

	ai.
		EXPECT().
		Review(mock.Anything, mock.Anything).
		Return(`{"issues":[{"line":1,"severity":"high","title":"nil check","suggestion":"add nil check"},{"line":2,"severity":"low","title":"style","suggestion":"rename var"}]}`, nil).
		Once()

	comments.
		EXPECT().
		CreateLineComment(mock.Anything, "acme/repo", 7, mock.Anything).
		Return(nil).
		Twice()

	comments.
		EXPECT().
		CreateComment(mock.Anything, "acme/repo", 7, mock.MatchedBy(func(body string) bool {
			return strings.Contains(body, "Total issues found: 2") &&
				strings.Contains(body, "Line comments posted: 2") &&
				strings.Contains(body, "High: 1") &&
				strings.Contains(body, "Low: 1")
		})).
		Return(nil).
		Once()

	p := NewProcessor(
		NewMemoryQueue(1),
		client,
		comments,
		dedup.NewMemory(),
		observability.NewLogger(&config.Config{LogLevel: "info"}),
		ai,
		ratelimit.New(100, 100),
	)

	p.handle(context.Background(), Job{Repo: "acme/repo", PR: 7})
}
