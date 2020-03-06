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
	missingRequired := cmd.getRequiredNames()

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
				delete(missingRequired, string(opt.short))
				delete(missingRequired, opt.long)
			}
			if opt.dataType != none && len(opt.arg) == 0 && canUseNext {
				execArgs.addFlag(opt.short, opt.long, m.args[i+1])
				delete(missingRequired, string(opt.short))
				delete(missingRequired, opt.long)
				i = i + 1
			}
		}

		i = i + 1
	}

	if len(missingRequired) != 0 {
		return nil, ErrInvalidArgs
	}

	err := m.processOperands(execArgs, operands, cmd)
	if err != nil {
		return nil, err
	}

	return execArgs, nil
}

func (m *model) processOperands(ea *execArgs, operands []string, cmd *Cmd) error {

	cmdOperands := cmd.getOperands()

	//If cmd there are no operands
	if len(cmdOperands) == 0 && len(operands) == 0 {
		return nil
	}

	//If the operands counts are not the same
	if len(cmdOperands) != len(operands) {
		return ErrInvalidArgs
	}

	//Verify operand's position and type
	for i, oper := range operands {

		//Check the position
		cmdOper, ok := cmdOperands[i]
		if !ok {
			return ErrInvalidArgs
		}

		if !m.validOperandType(cmdOper, oper) {
			return ErrInvalidArgs
		}
	}

	return nil
}

func (m *model) validOperandType(cmdOper operand, value string) bool {
	var stat bool

	//Chip trick for validating
	//the string by using util.go's
	//functions such ags getBool
	var b bool
	var i int64
	var u uint64
	var f float64

	switch cmdOper.dataType {
	case String:
		stat = true
	case Bool:
		stat = getBool(value, &b)
	case Int64:
		stat = getInt64(value, &i)
	case Uint64:
		stat = getUint64(value, &u)
	case Float64:
		stat = getFloat64(value, &f)
	default:
		stat = false
	}
	return stat
}
