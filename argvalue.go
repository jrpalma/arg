package arg

import (
	"strings"
)

type argValue interface {
	getOptions() ([]option, error)
}

func newPosixArg(value string, cmd *Cmd) *posixArg {
	return &posixArg{runes: getRunes(value), cmd: cmd}
}

type posixArg struct {
	cmd   *Cmd
	runes []rune
}

func (as *posixArg) getOptions() ([]option, error) {
	if len(as.runes) < 2 || as.runes[0] != '-' {
		return nil, nil
	}

	if as.runes[0] == '-' && as.runes[1] == '-' {
		return as.getLongOption()
	}

	return as.getShortOptions()
}

func (as *posixArg) getLongOption() ([]option, error) {
	var opt option
	var opts []option

	rhs := string(as.runes[2:])
	tokens := strings.Split(rhs, "=")

	cmdOpt, exist := as.cmd.longOption(tokens[0])
	if !exist {
		return nil, ErrInvalidArgs
	}

	//Check for argument
	opt.long = tokens[0]
	opt.short = cmdOpt.short
	opt.dataType = cmdOpt.dataType
	opt.required = cmdOpt.required
	if len(tokens) > 1 {
		opt.arg = tokens[1]
	}

	opts = append(opts, opt)
	return opts, nil
}
func (as *posixArg) getShortOptions() ([]option, error) {

	var opts []option
	rhs := as.runes[1:]

	//Process first option
	cmdOpt, exist := as.cmd.shortOption(rhs[0])
	if !exist {
		return nil, ErrInvalidArgs
	}

	opts = append(opts, option{
		dataType: cmdOpt.dataType,
		required: cmdOpt.required,
		long:     cmdOpt.long,
		short:    rhs[0],
	})

	//Start processing subsequent options
	for i := 1; i < len(rhs); i++ {
		//previous option has this argument
		cmdOpt, exist := as.cmd.shortOption(rhs[i])
		if !exist {
			opts[i-1].arg = string(rhs[i:])
			break
		}

		//another good option
		opts = append(opts, option{
			dataType: cmdOpt.dataType,
			required: cmdOpt.required,
			long:     cmdOpt.long,
			short:    rhs[i],
		})
	}

	return opts, nil
}
