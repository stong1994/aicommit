package llms

import (
	"context"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type Ollama struct {
	llm *ollama.LLM
}

func NewOllama(model string) *Ollama {
	llm, err := ollama.New(ollama.WithModel(model))
	if err != nil {
		log.Fatal("error creating ollama model: ", err)
	}
	return &Ollama{
		llm: llm,
	}
}

func (o *Ollama) GenerateContent(ctx context.Context, prompt, diff string) (string, error) {
	response, err := o.llm.GenerateContent(ctx, []llms.MessageContent{
		{
			Role:  llms.ChatMessageTypeSystem,
			Parts: []llms.ContentPart{llms.TextPart(prompt)},
		},
		{
			Role:  llms.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{llms.TextPart(diff)},
		},
	}, llms.WithStreamingFunc(
		func(ctx context.Context, chunk []byte) error {
			_, err := os.Stdout.Write(chunk)
			if err != nil {
				return err
			}
			return nil
		}),
	)
	if err != nil {
		return "", err
	}

	return response.Choices[0].Content, nil
}