# aicommit

A cli tool to generate commit content.

## Prepare

1. Make sure you have installed the [Ollama](https://ollama.com/) on your machine.
2. Pull the model that will used to summarize commit content, for example: `ollama pull llama3`. The recommended models are:

- [codegemma](https://ollama.com/library/codegemma)
- [codeqwen](https://ollama.com/library/codeqwen)
- [codellama](https://ollama.com/library/codellama)

## Build

```bash
git clone https://github.com/stong1994/aicommit.git
cd aicommit
go build -o aicommit main.go
```

## Run

```bash
export AICOMMIT_MODEL=codegemma
export PATH=$PATH:$PWD # or move aicommit to the /usr/local/bin
# cd to the git repository
aicommit
```

There are two ways to change the LLM model:

1.  Using environment variable `AICOMMIT_MODEL` to specify the model to use.
2.  Using `--model` flag to specify the model to use.
