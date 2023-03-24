package controller

import (
	"IroBot/config"
	"IroBot/model"
	"errors"
	"strings"
)

/*
 * @brief Check if the message is a command by prefix.
 * @param msg: The message to check.
 * @return bool: Whether the message is a command.
 */
func isIroCommand(msg string) bool {
	if len(msg) <= len(config.Config().Basic.Key) {
		return false
	}
	if strings.HasPrefix(msg, config.Config().Basic.Key+" ") {
		return true
	}
	return false
}

/*
 * @brief Check prefix and remove it if can.
 * @param msg: The message to check.
 * @return string: The message without prefix.
 */
func decodeIroCommand(msg string) (string, error) {
	if !isIroCommand(msg) {
		return "", errors.New("is not a command")
	}
	return strings.TrimSpace(strings.TrimPrefix(msg, config.Config().Basic.Key)), nil
}

/*
 * @brief Parse the message to a command.
 * @param msg: The message to be parsed.
 * @return model.Instruction: The parsed instruction.
 */
func Parse(msg string) (model.Instruction, error) {
	msg, err := decodeIroCommand(msg)
	if err != nil {
		return model.Instruction{}, err
	}

	ret := model.Instruction{}

	msg = strings.TrimLeft(msg, " ")
	cmdWithArgs := strings.Split(msg, " ")[0]
	cmd := strings.Split(cmdWithArgs, "[")[0]
	ret.OptionName = cmd

	if cmd == cmdWithArgs {
		ret.HasArg = false
	} else {
		args := strings.TrimPrefix(cmdWithArgs, cmd)
		if strings.HasPrefix(args, "[") && strings.HasSuffix(args, "]") {
			args = strings.TrimPrefix(args, "[")
			args = strings.TrimSuffix(args, "]")
			ret.HasArg = true
			ret.Args = args
		} else {
			return ret, errors.New("invalid instruction")
		}
	}
	content := strings.TrimPrefix(msg, cmdWithArgs)
	content = strings.TrimPrefix(content, " ")
	content = strings.TrimPrefix(content, "\n")
	ret.Content = content
	return ret, nil
}
