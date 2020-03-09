![](doc/mascot.png)

# A[rg] Command Line Argument Parser for Go
[![CircleCI status](https://circleci.com/gh/jrpalma/arg.png?circle-token=:circle-token "CircleCI status")](https://circleci.com/gh/jrpalma/arg)
[![GoDoc](https://godoc.org/github.com/jrpalma/arg?status.svg)](https://godoc.org/github.com/jrpalma/arg)
[![Go Report Card](https://goreportcard.com/badge/github.com/jrpalma/arg)](https://goreportcard.com/report/github.com/jrpalma/arg)



# Table of Contents
- [Documentation](#documenation)
- [Overview](#Overview)
- [Installing](#Installing)
- [Concepts](#concetps)
  * [Commands](#commands)
  * [Arguments](#arguments)
  * [Options](#options)
  * [Operands](#operands)
  * [Parser](#parser)
- [Contributing](#contributing)

# Documentation
You can find the package documentation at https://godoc.org/github.com/jrpalma/arg

# Overview
Arg is a Go package that provides a simple Application Programming Interface to parse command line arguments. Arg's API allows users to
specify Command Line Interfaces in a simple and declarative way. Arg provides the following features:

* Simple declarative API
* Short and long options support
* Man page style command line help
* Simple type based verification

# Installing
You can install arg under your GOPATH if your version of Go does not support modules. Run the following command to install arg under
your GOPATH:
```sh
go get github.com/jrpalma/arg
```
Arg also support Go modules. All you have to do is initialize your project and import arg into your project. If your project has not been initialized, all you have to do is 
run the following command under your project:
```sh
# You will have to specify <project path>
# if you are outside your GOPATH
go mod init <project path>
```

# Concepts
At its core, arg uses Cmd objects to describe a CLI and execute actions. A Parser object use commands to parse command line arguments and execute commands that match the arguments.
The combination of commands and parser allow users to specify actions that get triggered when the correct arguments are passed to the application through the command line.
Arg's CLIs can be easily described as `APP PREFIX COMMAND OPTIONS OPERANDS`. The following list describes this in more details.
- APP: The compiled application
- PREFIX: The command's prefix. This can be any string including the empty string.
- COMMAND: This is the command's name. The name can be empty for the root command.
- OPTIONS: Arguments that start with - or --
- OPERANDS: Arguments passed to the command after all the options.

## Commands
Commands describe the CLI and execute actions when the arguments match the command. The command has the following fields: Prefix, Name, Description, and Exec.
The Prefix can be any string including the empty string. Name is the command's name. Description is a short description for the command and is used to render the command's help. The Exec method is called when the parser matches the argument against. The following sample code creates a command that prints a version number:
```go
verCmd := &arg.Cmd {
	Prefix: "engine",
	Name: "version",
	Description:  "Prints the engine version",
	Exec: func(argv arg.ExecArgs) error {
		fmt.Printf("Version: v1.0.0\n")
		return nil
	},
}
```
## Arguments
A command's Exec function receives its options and operands through the ExecArgs interface. ExecArgs is a convenience interface that allows quick access to options and operands. The following example shows how ExecArgs can be used to check options, get option values, and get operands value:
```go
cp := &arg.Cmd{
	Name:        "",
	Description: "cp - Copy SRC to DEST",
	Exec: func(ea arg.ExecArgs) error {
		var src, dest, t string
		recursive := ea.HasOption("r")
		if e.GetOption("t", &t) {
			return fmt.Errorf("Invalid t option")
		}
		if ea.GetOperand(0, &src) {
			return fmt.Errorf("Invalid SRC operand")
		}
		if ea.GetOperand(1, &dest) {
			return fmt.Errorf("Invalid DEST operand")
		}
		...
		return nil
	},
}
cp.Option('r', "", "Recursive")
cp.ReqString('t', "", "Type")
cp.StringOperand(0, "SRC")
cp.StringOperand(1, "DEST")

```

## Options
Options can be short, long, or both. Options can require a parameter or not. Finally, Options themselves can be required or not.
The parser will not call the command's Exec function if the arguments do not match the command's options. For example, the parser
will not call a command's Exec function if a required option is missing or is the incorrect type. The following sample code illustrates these statements:
```go
cmd := &arg.Cmd{Name: "commit"}

//Required short option '-b' with a string paremeter
cmd.ReqString('b', "", "Branch name")

//Required long option '--branch" with a string parameter
cmd.ReqString(0, "branch", "Branch name")

//Required '-b' or '--branch' option with a string paremeter
cmd.ReqString('b', "branch", "Branch name")

//Optional '-f' or '--format' option with a string parameter
cmd.OptString('f', "format", "The format used to display")

//Optional '-f' or '--force' without a paremter
cmd.Option('f', "force", "Force the operation")
```

## Operands
Operands are arguments that are passed to commands after all the options are parsed or after '--' argument marker.
Operands have a position and a type. The parser validates the operand position and type before calling Exec method on a command.
The parser will not call the command Exec function if the arguments do not match the command's operands. The following sample code
illustrates how operands can be used:
```go
cmd := &arg.Cmd{
	Name: "cp",
	Descriontion: "cp copies SOURCE to DEST",
}
cmd.StringOperand(0, "SOURCE")
cmd.StringOperand(1, "DEST")
```

## Parser
Parser objects match the command line arguments to commands. If the arguments match a command, then the command's Exec function is called.
The following code example demonstrate the use of a parser object:
```go
func main() {
	cp := &arg.Cmd{
		Name:        "",
		Description: "cp - Copy SRC to DEST",
		Exec: func(ea arg.ExecArgs) error {
			var src, dest, t string
			recursive := ea.HasOption("r")
			if e.GetOption("t", &t) {
				return fmt.Errorf("Invalid t option")
			}
			if ea.GetOperand(0, &src) {
				return fmt.Errorf("Invalid SRC operand")
			}
			if ea.GetOperand(1, &dest) {
				return fmt.Errorf("Invalid DEST operand")
			}
			...
			return nil
		},
	}
	cp.Option('r', "", "Recursive")
	cp.ReqString('t', "", "Type")
	cp.StringOperand(0, "SRC")
	cp.StringOperand(1, "DEST")

	parser := arg.NewParser(os.Stdout)
	parser.AddCmd(cp)

	err := parser.Parse(false, os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
```

# Contributing
1. Fork it
2. Clone it `git clone https://github.com/user_name/arg && cd arg`)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Make changes and add them (`git add .`)
5. Commit your changes (`git commit -m 'Add some feature'`)
6. Push to the branch (`git push origin my-new-feature`)
7. Create new pull request
