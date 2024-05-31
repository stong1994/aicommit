package llms

import (
	"context"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type Ollama struct {
	llm *ollama.LLM
}

func NewOllama(model string) *Ollama {
	if model == "" {
		model = "codegemma"
	}

	llm, err := ollama.New(ollama.WithModel(model))
	if err != nil {
		log.Fatal("error creating ollama model: ", err)
	}
	return &Ollama{
		llm: llm,
	}
}

func (o *Ollama) GenerateContent(
	ctx context.Context,
	prompt, diff string,
	streamingFn func(ctx context.Context, chunk []byte) error,
) (string, error) {
	response, err := o.llm.GenerateContent(
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
