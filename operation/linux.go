package command

import (
	"debug/elf"
	"os/exec"
)

type Linux struct {
}
type Command string

const (
	Apt Command = "apt"
	Dnf Command = "dnf"
)

var (
	pkgCommand Command
)

func init() {
	if commandIsExists(Apt) {
		pkgCommand = Apt
	}
	if commandIsExists(Dnf) {
		pkgCommand = Dnf
	}

}

func (linux Linux) Install(name string) error {

}

func commandIsExists(command Command) bool {
	_, err := exec.LookPath(string(command))
	return err == nil
}
