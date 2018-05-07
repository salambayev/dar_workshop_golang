package detector

import (
	"fmt"
	"os/exec"
	"strings"
)

type ExecuteCommand interface {
	ExecuteCommand(command string) ([]byte, error)
}

type executeCommand struct {
}

func NewExecuteCommand() ExecuteCommand {
	return &executeCommand{}
}

func (ec *executeCommand) ExecuteCommand(command string) ([]byte, error) {
	fmt.Println("Executing command: " + command)
	parts := strings.Split(command, " ")
	head := parts[0]
	tail := parts[1:len(parts)]
	out, err := exec.Command(head, tail...).Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}
