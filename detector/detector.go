package detector

import (
	"strings"

	"github.com/pkg/errors"
)

type DetectorCommand interface {
	ExecuteCommand(command string) ([]byte, error)
}

type detectorCommand struct {
	executor ExecuteCommand
}

func NewDetectorCommand(ex ExecuteCommand) DetectorCommand {
	return &detectorCommand{ex}
}

func (d *detectorCommand) ExecuteCommand(command string) ([]byte, error) {
	parts := strings.Split(command, " ")
	if parts[0] == "rm" {
		return nil, errors.New("You have not access to remove anything! ")
	}
	return d.executor.ExecuteCommand(command)
}
