package input

import (
	"fmt"
	"strconv"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

func Parse(line string) (Command, error) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return Command{}, fmt.Errorf("command is empty")
	}
	return Command{Name: strings.ToLower(parts[0]), Args: parts[1:]}, nil
}

func IntArg(command Command, position int) (int, error) {
	if position < 0 || position >= len(command.Args) {
		return 0, fmt.Errorf("argument %d is missing", position)
	}
	value, err := strconv.Atoi(command.Args[position])
	if err != nil {
		return 0, fmt.Errorf("argument %d must be an integer", position)
	}
	return value, nil
}

func RequireArgs(command Command, count int) error {
	if len(command.Args) < count {
		return fmt.Errorf("%s requires %d arguments", command.Name, count)
	}
	return nil
}

func Help() string {
	return "status | move DX DY | scan RADIUS | fire SLOT ASTEROID | upgrade SLOT | turn | save | quit"
}
