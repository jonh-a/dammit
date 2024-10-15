package pkg

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func GetPrompt() string {
	prompt := `Given the following system information and command input + output, 
dignose the problem and provide a solution. Do not repeat the system info.

If there area any obvious typos in the command, provide a suggested command 
at the end of your reply in plain text, prefaced with "Recommended command: ". The
command should be on the same line as "Recommended command: " and there should be no
text following it. Do not include any backticks surrounding the command and do not include
any lines of text after it.
`

	verbosity := viper.Get("VERBOSITY")

	if verbosity == 0 {
		prompt += "\nBe as succinct as possible."
	}
	if verbosity == 2 {
		prompt += "\nBe as detailed as possible."
	}

	return prompt
}

func CallLLM(message string) string {
	model := getModel()
	llm, err := ollama.New(ollama.WithModel(model))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	completion, err := llm.Call(ctx, message,
		llms.WithTemperature(0.1),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)

	if err != nil {
		fmt.Println(checkError(err))
	}

	return completion
}

func checkError(err error) string {
	e := fmt.Sprintf("An error occurred while calling the LLM: %s", err)
	if strings.Contains(e, "connect: connection refused") {
		return fmt.Sprintf("%s\nIt looks like Ollama may not be running.", e)
	}
	if strings.Contains(e, "try pulling it first") {
		return fmt.Sprintf("%s\nThe required Ollama model is not installed.", e)
	}
	return e
}

func getModel() string {
	if str, ok := viper.Get("MODEL").(string); ok {
		return str
	} else {
		log.Fatalf("Invalid model selected.")
		return ""
	}
}
