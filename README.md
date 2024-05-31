# aicommit

A cli tool to generate commit content.

## Usage

The tool needs LLM support, you can use ollama or 零一万物 for now.

### ollama

#### Prepare

1. Make sure you have installed the [Ollama](https://ollama.com/) on your machine.
2. Pull the model that will used to summarize commit content, for example: `ollama pull llama3`. The recommended models are:

- [codegemma](https://ollama.com/library/codegemma)
- [codeqwen](https://ollama.com/library/codeqwen)
- [codellama](https://ollama.com/library/codellama)

### 零一万物

#### Prepare

Make sure you got the API Key of 零一万物. You can get it from [there](https://platform.lingyiwanwu.com/apikeys).

## Build

```bash
git clone https://github.com/stong1994/aicommit.git
cd aicommit
go mod tidy
go build -o aicommit main.go
```

## set it to global

```bash
export PATH=$PATH:$PWD # or move aicommit to the /usr/local/bin
```

## Run

There are some arguments you can use:

- platform: the platform of llm, you can use ollama or lingyi
- model: the llma model you want to use, for example: codegemma, codeqwen, codellama (for ollma), yi-large (for lingyi)
- quiet: if true, aicommit will output the command directly, if false, the output will using be streaming

All the arguments can be set with environment:

```bash
export AICOMMIT_MODEL=codegemma
export AICOMMIT_PLATFORM=ollama
export AICOMMIT_QUIET=true
```

# cd to the git repository

aicommit

## Notice

The tool uses `git diff --cached --diff-algorithm=minimal` to get the diff content, so make sure the output of `git diff --cached --diff-algorithm=minimal` is not empty.

```

```
