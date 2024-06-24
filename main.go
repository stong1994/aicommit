package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	sllms "github.com/stong1994/aicommit/llms"
)

const prompt = `You are an AI programming assistant. 
Write commit message for the change with commitizen convention. Make sure the title has maximum 50 characters and message is wrapped at 72 characters. So, Your output will be executed directly, so make sure it is executable, and the command's formatter will be "git commit -m '{type}: {title}'". If the input content is huge, you can add more details with another '-m' tag. The type should follow rules below:
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
			`

var (
	model    string
	platform string // platform of llm, either ollam or lingyi
	quiet    bool
	needCopy bool
)

type LLM interface {
	GenerateContent(
		ctx context.Context,
		prompt, diff string,
		streamingFn func(ctx context.Context, chunk []byte) error,
	) (string, error)
}

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
		var streamingFn func(ctx context.Context, chunk []byte) error
		if !quiet {
			streamingFn = func(ctx context.Context, chunk []byte) error {
				_, err := os.Stdout.Write(chunk)
				if err != nil {
					return err
				}
				return nil
			}
		}

		var llm LLM
		switch platform {
		case "ollama":
			llm = sllms.NewOllama(model)
		case "lingyi":
			llm = sllms.NewLingyi(model)
		case "github":
			llm = sllms.NewGithub(model)
		default:
			log.Fatal("invalid platform, only suppport github, ollama and lingyi but found: ", platform)
		}
		command, err := llm.GenerateContent(context.Background(), prompt, diff, streamingFn)
		if err != nil {
			log.Fatal("error generating content: ", err)
		}

		// 输出总结结果
		if quiet {
			fmt.Println(command)
		}
		if needCopy {
			if err = clipboard.WriteAll(command); err != nil {
				log.Fatal("error copy to clipboard: ", err)
			}
		}
	},
}

func init() {
	model = os.Getenv("AICOMMIT_MODEL")
	rootCmd.Flags().StringVar(&model, "model", model, "AI model to use for summarizing git commit differences")
	platform = os.Getenv("AICOMMIT_PLATFORM")
	rootCmd.Flags().StringVar(&platform, "platform", platform, "platform to run llm")
	quiet = os.Getenv("AICOMMIT_QUIET") == "true"
	rootCmd.Flags().BoolVar(&quiet, "quiet", quiet, "if set, use text to repsponse, otherwise, use streaming to response")
	needCopy = os.Getenv("AICOMMIT_COPY") == "true"
	rootCmd.Flags().BoolVar(&needCopy, "copy", needCopy, "if set, the command will copy to clipboard automiticly")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var diffCached = []string{"diff", "--cached", "--diff-algorithm=minimal"}

func getDiffWithLastCommit() (string, error) {
	cmd := exec.Command("git", diffCached...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}
