package chunker

import (
	"strings"
	"unicode/utf8"
)

type Chunk struct {
	File    string
	Content string
}

type Chunker struct {
	MaxTokens int
}

func New(max int) *Chunker {
	return &Chunker{MaxTokens: max}
}

// Rough token estimation (no tiktoken dependency)
func estimateTokens(s string) int {
	return utf8.RuneCountInString(s) / 3
}

func (c *Chunker) Split(file, content string) []Chunk {

	if estimateTokens(content) <= c.MaxTokens {
		return []Chunk{{File: file, Content: content}}
	}

	var chunks []Chunk
	var current strings.Builder

	for _, line := range strings.Split(content, "\n") {

		if estimateTokens(current.String()+line) > c.MaxTokens {

			chunks = append(chunks, Chunk{
				File:    file,
				Content: current.String(),
			})

			current.Reset()
		}

		current.WriteString(line + "\n")
	}

	if current.Len() > 0 {
		chunks = append(chunks, Chunk{
			File:    file,
			Content: current.String(),
		})
	}

	return chunks
}
