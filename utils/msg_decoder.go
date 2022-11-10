package utils

import (
	"IroBot/config"
	"errors"
	"strings"
)

func IsIroCommand(msg string) bool {
	if len(msg) <= len(config.Config().Basic.Key) {
		return false
	}
	if strings.HasPrefix(msg, config.Config().Basic.Key+" ") {
		return true
	}
	return false
}

func DecodeIroCommand(msg string) (string, error) {
	if !IsIroCommand(msg) {
		return "", errors.New("is not a command")
	}
	return strings.TrimPrefix(msg, config.Config().Basic.Key+" "), nil
}
