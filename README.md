# aicommit

A cli tool to generate commit content.

## Prepare

1. Make sure you have installed the [Ollama](https://ollama.com/) on your machine.
2. Pull the model that will used to summarize commit content, for example: `ollama pull llama3`. The recommended models are:

- [codegemma](https://ollama.com/library/codegemma)
- [codeqwen](https://ollama.com/library/codeqwen)
- [codellama](https://ollama.com/library/codellama)

## Run

```bash
go run main.go --model codegemma
```

default model is `codegemma` if not provided.
