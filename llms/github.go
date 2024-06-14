package llms

import (
	"context"
	"log"
	"os"

	copilot "github.com/stong1994/github-copilot-api"
)

type Github struct {
	llm *copilot.Copilot
}

func NewGithub(model string) *Github {
	if model == "" {
		model = "gpt-4"
	}
	llm, err := copilot.NewCopilot(
		copilot.WithModel(model),
		copilot.WithGithubToken(os.Getenv("GITHUB_OAUTH_TOKEN")),
	)
	if err != nil {
		log.Fatal("error creating github copilot: ", err)
	}
	return &Github{
		llm: llm,
	}
}

func (l *Github) GenerateContent(
	ctx context.Context,
	prompt, diff string,
	streamingFn func(ctx context.Context, chunk []byte) error,
) (string, error) {
	response, err := l.llm.CreateCompletion(
		ctx,
		&copilot.CompletionRequest{
			Messages: []copilot.Message{
				{
					Role:    "system",
					Content: prompt,
				},
				{
					Role:    "user",
					Content: diff,
				},
			},
			StreamingFunc: streamingFn,
			Temperature:   0.1,
		},
	)
	if err != nil {
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}
