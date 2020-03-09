package arg

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
)

// ErrInvalidArgs Error to signal invalid arguments
var ErrInvalidArgs = errors.New("Invalid Arguments")

// NewParser creates a new Parser object. A Parser object is used
// to parse the arguments and execute the command that matches
// the arguments. A Parser will write to the writer on a failure
// if the writer is not nil.
func NewParser(writer io.Writer) *Parser {
	parser := &Parser{
		pfxs:   make(map[string]*prefix),
		output: writer,
	}

	return parser
}

// Parser represents a parser object used to parse the arguments and
// execute the command that matches the arguments.
type Parser struct {
	args   []string
	output io.Writer
	pfxs   map[string]*prefix
}

// AddCmd adds a command that will execute if the Parser's Parse
// function can match command against the arguments.
func (p *Parser) AddCmd(cmd *Cmd) {
	cmdPfx := newPrefix(cmd.Prefix)
	pfx, ok := p.pfxs[cmdPfx.str]

	if !ok {
		p.pfxs[cmdPfx.str] = cmdPfx
		cmdPfx.addCmd(cmd)
		return
	}

	pfx.addCmd(cmd)
}

// Parse parses the arguments and execute the appropriate command.
// It is assumed that args contains the name of the executable as
// the first element in the args slice. A Parser will return
// ErrInvalidArgs if len(args) < 2.
// A Parser will print to the writer the commands usage, or the any
// non-nil error that is returned by a comand.
//
// Argument exitOnFailure
//
//
// A Parser's Parse function will exit the program if: "the parser
// cannot find a command to execute, or the command failed to
// execute and returns a non nil error".
//
// A Parser's Parse function will return: "ErrInvalidArgs if the
// parser cannot find a command to execute, or the command's
// returned error"
func (p *Parser) Parse(exitOnError bool, args []string) error {
	var cmd *Cmd
	var err error
	var execArgs ExecArgs

	if len(args) < 2 {
		return ErrInvalidArgs
	}

	execName := args[0]
	models := p.getModels(args[1:])

	for _, model := range models {
		pfx, ok := p.pfxs[model.pfx.str]

		if !ok || !pfx.equal(model.pfx) {
			continue
		}

		match, ok := pfx.getCmd(model.name)
		if !ok {
			continue
		}

		execArgs, err = model.parseArgs(match)
		if err != nil {
			p.cmdHelp(execName, pfx.str, match)
			return err
		}

		cmd = match
		break
	}

	if cmd == nil {
		p.usage(execName)
		if !exitOnError {
			return ErrInvalidArgs
		}
		os.Exit(1)
	}

	if cmd.Exec == nil {
		return nil
	}

	err = cmd.Exec(execArgs)
	if err != nil {
		p.print(err.Error())
		if !exitOnError {
			return err
		}
		os.Exit(1)
	}

	return nil
}
func (p *Parser) getModels(execArgs []string) []model {
	var models []model
	var prefix string

	for i, token := range execArgs {
		models = append(models, model{
			pfx:  newPrefix(prefix),
			args: execArgs[i+1:],
			name: token,
		})
		prefix += " " + token
	}

	return models
}
func (p *Parser) usage(execName string) {
	p.print("Usage:\n%v <command>\n\n", execName)
	p.print("COMMANDS:\n")
	for _, pfx := range p.sortedPrefixes() {
		for _, cmd := range pfx.sortedCmds() {
			p.cmdUsage(pfx.str, cmd)
		}
	}
}
func (p *Parser) cmdUsage(pfx string, cmd *Cmd) {

	var synopsis string
	help := tab + tab + cmd.Description + nl + nl

	if len(pfx) == 0 {
		synopsis += tab + cmd.Name
	} else if len(pfx) > 0 {
		synopsis += tab + pfx + space + cmd.Name
	}
	if cmd.hasOptions() {
		synopsis += space + "[OPTIONS]..."
	}
	for _, op := range cmd.sortedOperands() {
		synopsis += space + "<" + op.name + ">"
	}

	p.print(synopsis + nl + help)

}
func (p *Parser) cmdHelp(execName string, pfx string, cmd *Cmd) {
	synopsis := "SYNOPSIS:" + nl + tab
	help := "NAME:" + nl + tab + cmd.Description + nl

	if len(pfx) == 0 {
		synopsis += execName + space + cmd.Name
	} else if len(pfx) > 0 {
		synopsis += execName + space + pfx + space + cmd.Name
	}
	if cmd.hasOptions() {
		synopsis += space + "[OPTIONS]..."
	}
	for _, op := range cmd.sortedOperands() {
		synopsis += space + "<" + op.name + ">"
	}

	p.print(help + nl + synopsis + nl + nl)
	p.optsHelp(cmd)
}
func (p *Parser) optsHelp(cmd *Cmd) {

	longOptions := cmd.longOptions()
	shortOptions := cmd.sortedShortOptions()

	p.print("OPTIONS:\n")
	for _, shortOpt := range shortOptions {
		str := tab + "-" + string(shortOpt.short)
		if shortOpt.long != "" {
			delete(longOptions, shortOpt.long)
			str += ", --" + shortOpt.long
		}
		p.print(str + nl + tab + tab + shortOpt.description + nl + nl)
	}
	for _, longOpt := range longOptions {
		str := tab
		if longOpt.short != 0 {
			str += "-" + string(longOpt.short) + ", "
		}
		str += "--" + longOpt.long
		p.print(str + nl + tab + tab + longOpt.description + nl + nl)
	}

}
func (p *Parser) sortedPrefixes() []*prefix {
	var pfxs []*prefix
	var names []string
	for name := range p.pfxs {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, pfxName := range names {
		if pfx, ok := p.pfxs[pfxName]; ok {
			pfxs = append(pfxs, pfx)
		}
	}
	return pfxs
}
func (p *Parser) print(msg string, args ...interface{}) {
	if p.output != nil {
		p.output.Write([]byte(fmt.Sprintf(msg, args...)))
		//fmt.Printf(msg, args...)
	}
}

const (
	space = " "
	nl    = "\n"
	tab   = "    "
)
