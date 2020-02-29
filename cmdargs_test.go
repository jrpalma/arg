package arg

import (
	"testing"
)

func TestCmdArgs(t *testing.T) {
	args := CmdArgs{}
	args.Types(StringArg, BoolArg)

	if len(args.args) != 2 {
		t.Errorf("There should be 2 args: %#v", args.args)
	}
	if args.args[0] != StringArg {
		t.Errorf("The first string arg should be string")
	}
	if args.args[1] != BoolArg {
		t.Errorf("The second string arg should be bool")
	}

}
