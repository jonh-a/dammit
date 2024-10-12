package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func GetPrompt() string {
	prompt := `Given the following information, provide an explanation for what happened
and offer solutions.`
	return prompt
}

func CallLLM(message string, model string) string {
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
