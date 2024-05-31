package llms

import (
	"context"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/lingyi"
)

type Lingyi struct {
	llm *lingyi.LLM
}

func NewLingyi(model string) *Lingyi {
	if model == "" {
		model = "yi-large"
	}
	llm, err := lingyi.New(
		lingyi.WithAPIKey(os.Getenv("LINGYI_APIKEY")),
		lingyi.WithModel(model),
	)
	if err != nil {
		log.Fatal("error creating lingyi model: ", err)
	}
	return &Lingyi{
		llm: llm,
	}
}

func (l *Lingyi) GenerateContent(
	ctx context.Context,
	prompt, diff string,
	streamingFn func(ctx context.Context, chunk []byte) error,
) (string, error) {
	response, err := l.llm.GenerateContent(
		ctx,
		[]llms.MessageContent{
			{
				Role:  llms.ChatMessageTypeSystem,
				Parts: []llms.ContentPart{llms.TextPart(prompt)},
			},
			{
				Role:  llms.ChatMessageTypeHuman,
				Parts: []llms.ContentPart{llms.TextPart(diff)},
			},
		},
		llms.WithStreamingFunc(streamingFn),
	)
	if err != nil {
		return "", err
	}

	return response.Choices[0].Content, nil
}
