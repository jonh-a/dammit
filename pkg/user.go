package pkg

import (
	"os"
	"os/exec"
	"runtime"
)

type SystemData struct {
	OS   string
	Arch string
}

func GetShell() string {
	shell := os.Getenv("SHELL")
	return shell
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
		return "windows"
	}
	return OS
}

func GetSystemData() *SystemData {
	data := SystemData{OS: normalizeOSName(runtime.GOOS), Arch: runtime.GOARCH}
	return &data
}
