package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func GetPrompt() string {
	prompt := `Given the following system information and command input + output, 
dignose the problem and provide a solution. Do not repeat the system info.`

	verbosity := viper.Get("VERBOSITY")

	if verbosity == 0 {
		prompt += "\nBe as succinct as possible."
	}
	if verbosity == 2 {
		prompt += "\nBe as detailed as possible."
	}

	return prompt
}

func getModel() string {
	if str, ok := viper.Get("MODEL").(string); ok {
		return str
	} else {
		log.Fatalf("Invalid model selected.")
		return ""
	}
}

func CallLLM(message string) string {
	model := getModel()
	llm, err := ollama.New(ollama.WithModel(model))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	completion, err := llm.Call(ctx, message,
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)

	if err != nil {
		log.Fatal(err)
	}

	return completion
}
