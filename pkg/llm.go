package pkg

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/viper"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func GetPrompt() string {
	prompt := DefaultPrompt()

	verbosity := viper.Get("VERBOSITY")

	if verbosity == 0 {
		prompt += "\nBe as succinct as possible."
	}
	if verbosity == 2 {
		prompt += "\nBe as detailed as possible."
	}

	return prompt
}

func CallLLM(message string, model string, temperature float64) string {
	llm, err := ollama.New(ollama.WithModel(model))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	completion, err := llm.Call(ctx, message,
		llms.WithTemperature(temperature),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)

	if err != nil {
		fmt.Println(checkOllamaError(err))
	}

	return completion
}

func checkOllamaError(err error) string {
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
		fmt.Println("Invalid model selected. Defaulting to llama3.2:1b.")
		return "llama3.2:1b"
	}
}

func getTemperature() float64 {
	tempStr := viper.GetString("TEMPERATURE")
	if tempStr == "" {
		return 0.1
	}

	temp, err := strconv.ParseFloat(tempStr, 64)
	if err != nil {
		fmt.Println("Invalid temperature selected (must be a float between 0.1 and 1.0; default 0.1). Defaulting to 0.1.")
		return 0.1
	}

	if temp < 0.1 || temp > 1.0 {
		fmt.Println("Temperature out of range (must be between 0.1 and 1.0). Defaulting to 0.1.")
		return 0.1
	}

	return temp
}
