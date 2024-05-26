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

var rootCmd = &cobra.Command{
	Use:   "git-commit-summarizer",
	Short: "A tool to summarize git commit differences using Ollama and Llama 3",
	Long:  `This tool retrieves the differences between the current working directory and the last git commit, and summarizes it using the Ollama service with the Llama 3 model.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取当前工作目录与最近一次提交的差异
		diff, err := getDiffWithLastCommit()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("got diff: ", diff)

		// 使用Ollama服务和Llama 3模型进行总结
		summary, err := summarizeWithOllama(context.TODO(), diff)
		if err != nil {
			log.Fatal(err)
		}

		// 输出总结结果
		fmt.Println("Summary of changes:", summary)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// getDiffWithLastCommit 获取当前工作目录与最近一次提交的差异
func getDiffWithLastCommit() (string, error) {
	cmd := exec.Command("git", "diff", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// summarizeWithOllama 使用Ollama服务进行总结
func summarizeWithOllama(ctx context.Context, diff string) (string, error) {
	// 初始化Ollama客户端
	client, err := ollama.New(ollama.WithModel("llama3"))
	if err != nil {
		return "", err
	}

	// 调用Ollama服务进行总结
	response, err := client.GenerateContent(ctx, []llms.MessageContent{
		{
			Role:  llms.ChatMessageTypeSystem,
			Parts: []llms.ContentPart{llms.TextPart("Write commit message for the change with commitizen convention. Make sure the title has maximum 50 characters and message is wrapped at 72 characters. Wrap the whole message in code block with language gitcommit.")},
		},
		{
			Role:  llms.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{llms.TextPart(diff)},
		},
	})
	if err != nil {
		return "", err
	}

	return response.Choices[0].Content, nil
}
