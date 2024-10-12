# Dammit

Dammit is a remake of [thefuck](https://github.com/nvbn/thefuck) that uses a (configurable) LLM to troubleshoot terminal commands.

## Requirements

### Ollama

[Ollama](https://ollama.com/) should be installed in the current environment. Dammit uses the `llama3.2:1b` model by default, but other models can be used and specified with the `DAMMIT_MODEL` environment variable.

## Local development

Behavior can be mimicked by running `go run main.go run "$(fc -ln -1 -1)"` locally.