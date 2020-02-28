package arg

import (
	"testing"
)

func TestArgSet(t *testing.T) {
	args := ArgSet{}
	args.Add(StringArg)
	args.Add(BoolArg)

	if len(args.args) != 2 {
		t.Errorf("There should be 2 args: %#v", args.args)
	}

}
