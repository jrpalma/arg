package arg

type model struct {
	name string
	pfx  *prefix
	args []string
}

func (m *model) parseArgs(cmd *Cmd) (ExecArgs, error) {

	var skip bool
	var operands []string
	execArgs := newExecArgs()
	missingRequired := cmd.getRequiredLongNames()

	for i := 0; i < len(m.args); {

		argStr := m.args[i]
		argVal := newPosixArg(argStr, cmd)

		opts, err := argVal.getOptions()
		if err != nil {
			return nil, err
		}

		if skip || argStr == "--" || len(opts) == 0 {
			operands = append(operands, argStr)
			skip = true
			i = i + 1
			continue
		}

		canUseNext := (i + 1) < len(m.args)
		for _, opt := range opts {
			if opt.dataType == none {
				execArgs.setOption(opt.short, opt.long)
			}
			if opt.dataType != none && len(opt.arg) != 0 {
				execArgs.addFlag(opt.short, opt.long, opt.arg)
				delete(missingRequired, opt.long)
			}
			if opt.dataType != none && len(opt.arg) == 0 && canUseNext {
				execArgs.addFlag(opt.short, opt.long, m.args[i+1])
				delete(missingRequired, opt.long)
				i = i + 1
			}
		}

		i = i + 1
	}

	if len(missingRequired) != 0 {
		return nil, ErrInvalidArgs
	}

	missingOperands := cmd.getOperands()

	for i, oper := range operands {
		execArgs.setOperand(i, oper)
		delete(missingOperands, i)
	}

	if len(missingOperands) != 0 {
		return nil, ErrInvalidArgs
	}

	return execArgs, nil
}
