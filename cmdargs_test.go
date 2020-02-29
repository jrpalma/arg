package arg

import (
	"testing"
)

func TestCmdArgs(t *testing.T) {
	args := CmdArgs{}
	args.Add("url", StringArg)
	args.Add("dir", StringArg)

	if len(args.args) != 2 {
		t.Errorf("There should be 2 args: %#v", args.args)
	}
	if args.args[0].typ != StringArg {
		t.Errorf("The first string arg should be string")
	}
	if args.args[1].typ != StringArg {
		t.Errorf("The second string arg should be string")
	}

}
