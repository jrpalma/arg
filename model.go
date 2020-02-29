package arg

import (
	"strconv"
	"strings"
)

type model struct {
	name string
	pfx  *prefix
	args []string
}

func (m *model) getCmdOptions(cmd *Cmd) map[string]string {
	allOptions := cmd.Options.getOptions()
	cmdOptions := make(map[string]string)
	missing := make(map[string]struct{})

	for name, option := range allOptions {
		if option.required {
			missing[name] = struct{}{}
		}
	}

	for i := 0; i < len(m.args); {
		name := m.args[i]
		argsLeft := len(m.args) - (i + 1)

		if argsLeft <= 0 {
			return nil
		}

		option, isOption := allOptions[name]
		if !isOption || option == nil {
			return nil
		}

		args := m.args[i+1:]
		usedOptions := m.useOptions(option, args)
		usedOptionsLen := len(usedOptions)

		if usedOptionsLen == 0 {
			return nil
		}

		cmdOptions[name] = strings.Join(usedOptions, ",")
		i = i + usedOptionsLen + 1
		delete(missing, name)
	}

	if len(missing) != 0 {
		return nil
	}

	return cmdOptions
}

func (m *model) useOptions(f *option, args []string) []string {
	var used []string
	argsLen := len(args)

	if argsLen == 0 {
		return used
	}

	value := args[0]

	if f.optionType == stringOption || f.optionType == enumOption {
		used = append(used, value)
		return used
	} else if f.optionType == boolOption {
		if _, err := strconv.ParseBool(value); err == nil {
			used = append(used, value)
		}
	} else if f.optionType == int64Option {
		if _, err := strconv.ParseInt(value, 0, 64); err == nil {
			used = append(used, value)
		}
	} else if f.optionType == uint64Option {
		if _, err := strconv.ParseUint(value, 0, 64); err == nil {
			used = append(used, value)
		}
	} else if f.optionType == float64Option {
		if _, err := strconv.ParseFloat(value, 64); err == nil {
			used = append(used, value)
		}
	} else if f.optionType == argsOption {
		if len(f.args) <= argsLen {
			for i := 0; i < len(f.args); i++ {
				used = append(used, args[i])
			}
		}
	}

	return used
}
