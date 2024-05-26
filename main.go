package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

var model string

var rootCmd = &cobra.Command{
	Use:   "aicommit",
	Short: "A tool to summarize git commit differences using Ollama",
	Long:  `This tool retrieves the differences between the current working directory and the last git commit, and summarizes it using the Ollama service with the Llama 3 model.`,
	Run: func(cmd *cobra.Command, args []string) {
		diff, err := getDiffWithLastCommit()
		if err != nil {
			log.Fatal(err)
		}
		// log.Println("got diff: ", diff)

		// 使用Ollama服务和Llama 3模型进行总结
		summary, err := summarizeWithOllama(context.TODO(), diff)
		if err != nil {
			log.Fatal(err)
		}

		_ = summary
		// 输出总结结果
		// fmt.Println(summary)
	},
}

func init() {
	model = os.Getenv("AICOMMIT_MODEL")
	if model == "" {
		model = "codegemma"
	}
	rootCmd.PersistentFlags().StringVar(&model, "model", model, "AI model to use for summarizing git commit differences")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getDiffWithLastCommit() (string, error) {
	cmd := exec.Command("git", "diff", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func summarizeWithOllama(ctx context.Context, diff string) (string, error) {
	client, err := ollama.New(ollama.WithModel(model))
	if err != nil {
		return "", err
	}

	response, err := client.GenerateContent(ctx, []llms.MessageContent{
		{
			Role: llms.ChatMessageTypeSystem,
			Parts: []llms.ContentPart{llms.TextPart(`You are an AI programming assistant. 
				Your duty is to generating a git commit command with the diff content. 
				Please adhere to the following guidelines:
				1. The commit should include one title, one type and mutli details. 
				2. The title should be a summary of the commit content and should be less than 50 characters.
				3. The details should be a more detailed summary of the commit content, and the length of each detail should not exceed 72 characters.
				4. Choose a type following the rules below:
			     - docs: Documentation only changes,
			     - style:Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc),
			     - perf: A code change that improves performance,
			     - test: Adding missing tests or correcting existing tests,
					 - build: Changes that affect the build system or external dependencies,
					 - ci: Changes to our CI configuration files and scripts,
					 - chore: "Other changes that don't modify src or test files",
					 - revert: Reverts a previous commit,
			     - feat: A new feature,
			     - fix: A bug fix, 
			     - refactor: A code change that neither fixes a bug nor adds a feature,
				5. The format of command should be "git commit -m '{type}: {title}' -m '{detail1}' -m '{detail2}'", the count of -m tag depends on the details count.
				6. The details can be ignore If the commit content is simple and not important.
				7. Make sure the tile and details are concise.
				8. Your entire response will be passed directly into shell to execute, so make sure it's executable.
			`)},
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
