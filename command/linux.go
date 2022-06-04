package command

import "os/exec"

var (
	Apt            Command = newCommand("apt")
	Dnf            Command = newCommand("dnf")
	packageManager Command = newCommand("")
)

func installPackage(packageNmae ...string) error {
	if packageManager.command == "" {
		if commandIsExists(Apt) {
			packageManager = Apt
		}
		if commandIsExists(Dnf) {
			packageManager = Dnf
		}
	}
}

func commandIsExists(command Command) bool {
	_, err := exec.LookPath(command.command)
	return err == nil
}
