package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type SystemData struct {
	Arch       string
	OS         string
	Shell      string
	WorkingDir string
}

func GetShell() string {
	shell := os.Getenv("SHELL")
	return shell
}

func GetSystemData() *SystemData {
	data := SystemData{
		Arch:       runtime.GOARCH,
		OS:         normalizeOSName(runtime.GOOS),
		Shell:      GetShell(),
		WorkingDir: getWorkingDir(),
	}
	return &data
}

func ParseSystemData(sd *SystemData) string {
	return fmt.Sprintf("Arch: %s, OS: %s, Shell: %s, WorkingDir: %s",
		sd.Arch, sd.OS, sd.Shell, sd.WorkingDir)
}

func getWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return dir
}

func getLinuxDistro() (string, error) {
	out, err := exec.Command("lsb_release", "-d", "-s").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func getWindowsVersion() (string, error) {
	cmd := exec.Command("powershell", "Get-WmiObject", "-Class", "Win32_OperatingSystem")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func getMacOSVersion() (string, error) {
	out, err := exec.Command("sw_vers", "--productVersion").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func normalizeOSName(OS string) string {
	if OS == "darwin" {
		version, err := getMacOSVersion()
		if err != nil {
			return "macOS"
		}
		return "macOS " + version
	}
	if OS == "linux" {
		distro, err := getLinuxDistro()
		if err != nil {
			return "Linux"
		}
		return distro
	}
	if OS == "windows" {
		version, err := getWindowsVersion()
		if err != nil {
			return "Windows"
		}
		return version
	}
	return OS
}
