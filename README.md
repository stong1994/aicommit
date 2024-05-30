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

## Run

```bash
export AICOMMIT_MODEL=codegemma # or yi-large for 零一万物
export AICOMMIT_PLATFORM=ollama # or lingyi for 零一万物
export PATH=$PATH:$PWD # or move aicommit to the /usr/local/bin
# cd to the git repository
aicommit
```

There have two ways to change the LLM model:

1.  Using environment variable `AICOMMIT_MODEL` to specify the model to use.
2.  Using `--model` flag to specify the model to use.

There also have two ways to change the LLM platform:

1.  Using environment variable `AICOMMIT_PLATFORM` to specify the model to use.
2.  Using `--platform` flag to specify the model to use.

## Notice

The tool uses `git diff --cached --diff-algorithm=minimal` to get the diff content, so make sure the output of `git diff --cached --diff-algorithm=minimal` is not empty.
