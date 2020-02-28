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

func (m *model) getCmdFlags(cmd *Cmd) map[string]string {
	allFlags := cmd.Flags.getFlags()
	cmdFlags := make(map[string]string)
	missing := make(map[string]struct{})

	for name, flag := range allFlags {
		if flag.Required() {
			missing[name] = struct{}{}
		}
	}

	for i := 0; i < len(m.args); {
		name := m.args[i]
		argsLeft := len(m.args) - (i + 1)

		if argsLeft <= 0 {
			return nil
		}

		flag, isFlag := allFlags[name]
		if !isFlag || flag == nil {
			return nil
		}

		args := m.args[i+1:]
		usedFlags := m.useFlags(flag, args)
		usedFlagsLen := len(usedFlags)

		if usedFlagsLen == 0 {
			return nil
		}

		cmdFlags[name] = strings.Join(usedFlags, ",")
		i = i + usedFlagsLen + 1
		delete(missing, name)
	}

	if len(missing) != 0 {
		return nil
	}

	return cmdFlags
}

func (m *model) useFlags(f *flag, args []string) []string {
	var used []string
	argsLen := len(args)

	if argsLen == 0 {
		return used
	}

	value := args[0]

	if f.Type() == stringFlag || f.Type() == enumFlag {
		used = append(used, value)
		return used
	} else if f.Type() == boolFlag {
		if _, err := strconv.ParseBool(value); err == nil {
			used = append(used, value)
		}
	} else if f.Type() == int64Flag {
		if _, err := strconv.ParseInt(value, 0, 64); err == nil {
			used = append(used, value)
		}
	} else if f.Type() == uint64Flag {
		if _, err := strconv.ParseUint(value, 0, 64); err == nil {
			used = append(used, value)
		}
	} else if f.Type() == float64Flag {
		if _, err := strconv.ParseFloat(value, 64); err == nil {
			used = append(used, value)
		}
	} else if f.Type() == argsFlag {
		if f.Count() <= uint(argsLen) {
			for i := uint(0); i < f.Count(); i++ {
				used = append(used, args[i])
			}
		}
	}

	return used
}
