package llms

import (
	"context"
	"log"
	"os"
	"strconv"

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
		copilot.WithCompletionModel(model),
		copilot.WithTemperature(getTemperature()),
		copilot.WithGithubOAuthToken(os.Getenv("GITHUB_OAUTH_TOKEN")),
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
		},
	)
	if err != nil {
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}

func getTemperature() float64 {
	temperatureStr := os.Getenv("AICOMMIT_TEMPERATURE")
	if temperatureStr == "" {
		return 1.0
	}
	temperature, _ := strconv.ParseFloat(temperatureStr, 64)
	if temperature == 0 {
		return 1.0
	}
	return temperature
}
