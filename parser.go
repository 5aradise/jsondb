package jsondb

import (
	"errors"
	"strings"
)

func parseArgs(arg string) ([]string, error) {
	args := strings.Split(arg, ".")
	for i, arg := range args {
		arg = strings.TrimSpace(arg)
		if arg == "" {
			return nil, errors.New("invalid field name")
		}
		args[i] = arg
	}
	return args, nil
}
