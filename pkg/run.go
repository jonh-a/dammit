package pkg

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

func rerunCommand(command string) string {
	shell := GetShell()

	cmd := exec.Command(shell, "-c", command)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Sprintf("Error: %s\nOutput: %s", err.Error(), string(output))
	}

	return string(output)
}

func Run(command string) string {
	label := fmt.Sprintf("Rerun %s", command)

	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	systemData := GetSystemData()
	dataString := fmt.Sprintf("Running %s", systemData.OS)

	shouldRerun, _ := prompt.Run()

	out := dataString + "\nCommand: " + command

	if strings.ToLower(shouldRerun) == "y" {
		out += "\n" + rerunCommand(command)
	}

	fmt.Println(out)

	completion := CallLLM(GetPrompt() + out)

	return completion
}
