package pkg

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Init() {
	shell := normalizeShellName(GetShell())
	rc, err := determineShellRc(shell)
	if err != nil {
		log.Fatal(err)
	}

	appendAliasToRc(shell, rc)
}

func normalizeShellName(shell string) string {
	if strings.Contains(shell, "zsh") {
		return "zsh"
	}
	if strings.Contains(shell, "bash") {
		return "bash"
	}
	if strings.Contains(shell, "fish") {
		return "fish"
	}
	return shell
}

func determineShellRc(shell string) (string, error) {
	switch normalizeShellName(shell) {
	case "zsh":
		return ".zshrc", nil
	case "bash":
		return ".bashrc", nil
	default:
		return "", errors.New("Unable to determine shell rc file")
	}
}

func appendAliasToRc(shell string, rc string) {
	if shell == "bash" || shell == "zsh" {
		fmt.Println("Adding shell alias...")
		alias := `dammit() { go_dammit run "$(fc -ln -1 -1)" }`

		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Error occurred when getting the home directory: %s", err)
		}

		rc = filepath.Join(homeDir, rc)

		file, err := os.OpenFile(rc, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Fatalf("Error occurred when opening %s: %s", rc, err)
		}
		defer file.Close()

		_, err = file.WriteString("\n" + alias + "\n")
		if err != nil {
			log.Fatalf("Error occurred when writing to %s: %s", rc, err)
		}

		printSuccessInstructions(rc)
	}
}

func printSuccessInstructions(rc string) {
	fmt.Printf(`Alias successfully added to %s! 
	
Run "source %s" to apply the new alias or open a new terminal.

After doing so, you can run "dammit" from your terminal to troubleshoot
your most recent command. Note:
* The alias expects that "go_dammit" can be found in your path.
* Dammit expects that Ollama is currrently running on your system.

Dammit will use the "llama3.2:1b" model by default. You can configure the model and
output verbosity by setting the following environment variables:
* "DAMMIT_MODEL": The name of the Ollama model to use
* "DAMMIT_VERBOSITY": An integer between 0 and 2
`, rc, rc)
}
