# AICommit

AICommit is a CLI tool that generates commit content using various AI models.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Environment Variables](#environment-variables)
- [Notice](#notice)

## Prerequisites

AICommit requires support from Language Learning Models (LLM). You can use either Ollama, 零一万物, or Github Copilot.

### Ollama

1. Install [Ollama](https://ollama.com/) on your machine.
2. Pull the model that will be used to summarize commit content, for example: `ollama pull llama3`. The recommended models are:
   - [codegemma](https://ollama.com/library/codegemma)
   - [codeqwen](https://ollama.com/library/codeqwen)
   - [codellama](https://ollama.com/library/codellama)

### 零一万物

Obtain the API Key of 零一万物 from [here](https://platform.lingyiwanwu.com/apikeys).

### Github Copilot

Ensure you are logged into Github Copilot. There should be a `host.json` in your config directory (e.g., `~/.config/github-copilot` on macOS), or you can specify the oauth token with the `GITHUB_OAUTH_TOKEN` environment variable.

## Installation

```bash
git clone https://github.com/stong1994/aicommit.git
cd aicommit
go mod tidy
go build -o aicommit main.go
```

To make the tool globally accessible, add it to your PATH:

```bash
export PATH=$PATH:$PWD # or move aicommit to the /usr/local/bin
```

## Usage

Run the tool in the git repository with the `aicommit` command.

The tool accepts the following arguments:

- `platform`: The LLM platform to use. Options are 'ollama', 'lingyi', or 'github'.
- `model`: The LLM model to use. For example: 'codegemma', 'codeqwen', 'codellama' (for Ollama), 'yi-large' (for Lingyi), 'gpt-4' (for github copilot).
- `quiet`: If set to true, AICommit will output the command directly. If false, the output will be streamed.
- `copy`: If set to true, AICommit will copy the command to the clipboard.

## Environment Variables

All the arguments can also be set with environment variables:

```bash
export AICOMMIT_MODEL=codegemma
export AICOMMIT_PLATFORM=ollama
export AICOMMIT_QUIET=true
export AICOMMIT_COPY=true
```

## Notice

AICommit uses `git diff --cached --diff-algorithm=minimal` to get the diff content. Ensure the output of `git diff --cached --diff-algorithm=minimal` is not empty.
