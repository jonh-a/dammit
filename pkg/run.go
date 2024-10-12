package pkg

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os/exec"
)

func Rerun(command string) string {
	shell := GetShell()

	cmd := exec.Command(shell, "-c", command)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Sprintf("Error: %s\nOutput: %s", err.Error(), string(output))
	}

	return string(output)
}

func Ask(command string) string {
	label := fmt.Sprintf("Rerun %s?", command)

	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	rerun, _ := prompt.Run()

	fmt.Println(rerun)

	return ""
}
