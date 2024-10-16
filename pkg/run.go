package pkg

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

func Run(command string) {
	// Prompt user to re-run the last command.
	// If "no" is selected, then we can't capture stdout/stderr
	// and the completion will have to be generated only by the command.
	label := fmt.Sprintf("Rerun %s", command)

	rerunPrompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	sd := GetSystemData()
	dataString := ParseSystemData(sd)

	shouldRerun, _ := rerunPrompt.Run()

	out := dataString + "\nCommand: " + command

	if strings.ToLower(shouldRerun) == "y" {
		out += "\n" + runCommand(command)
	}
	fmt.Println(out)

	// Get advice from the LLM. The LLM /should/ provide a recommended
	// alternative command, if one is available. We can parse this from
	// the response.
	completion := CallLLM(GetPrompt() + out)
	recommendedCommand := getRecommendedCommand(completion)

	if recommendedCommand == "" {
		return
	}

	// Ask the user to run the suggested command, if one is available.
	label = fmt.Sprintf("Run %s", recommendedCommand)

	runPrompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	shouldRun, _ := runPrompt.Run()
	newOut := recommendedCommand

	if strings.ToLower(shouldRun) == "y" {
		newOut += "\n" + runCommand(recommendedCommand)
		fmt.Println(newOut)
	}
}

func runCommand(command string) string {
	shell := GetShell()

	cmd := exec.Command(shell, "-c", command)
	output, err := cmd.Output()

	if err != nil {
		return fmt.Sprintf("Error: %s\nOutput: %s", err.Error(), string(output))
	}

	return string(output)
}

func getRecommendedCommand(response string) string {
	lines := strings.Split(response, "\n")
	lastLine := lines[len(lines)-1]

	if !strings.HasPrefix(lastLine, "Recommended command: ") {
		return ""
	}

	command := strings.Split(lastLine, "Recommended command: ")[1]

	return command
}
