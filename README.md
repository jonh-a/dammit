# Dammit

Dammit is a remake of [thefuck](https://github.com/nvbn/thefuck) that uses a (configurable) LLM to troubleshoot terminal commands.

## Requirements

### Ollama

[Ollama](https://ollama.com/) should be installed in the current environment. Dammit uses the `llama3.2:1b` model by default, but other models can be used and specified with the `DAMMIT_MODEL` environment variable.

## Installation

Run `go_dammit init` to add the `dammit` function to your shell's .rc file.

With this function added, you can now run `dammit` in your terminal to fetch the last command, (optionally) re-run it to capture terminal output, and ask your preferred LLM what's going on.

## Configuration

Dammit can be configured, either via environment variables or a `.dammit.yaml` file.

The following variables can be set:

- MODEL (env. `DAMMIT_MODEL`): The name of the Ollama model to use (default `llama3.2:1b`)
- VERBOSITY (env. `DAMMIT_VERBOSITY`): An integer between 0-2 representing the response verbosity (default `1`).
- TEMPERATURE (env. `DAMMIT_TEMPERATURE`): A float between 0.1-1.0 representing the response temperature (default `0.1`).

## Local development

Behavior can be mimicked by running `go run main.go run "$(fc -ln -1 -1)"` locally.