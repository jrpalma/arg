package arg

import (
	"bytes"
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {

	show := &Cmd{
		Prefix: "users",
		Name:   "show",
		Help:   "Show users",
		Exec: func(args ExecArgs) error {
			return nil
		},
	}

	delete := &Cmd{
		Prefix: "users",
		Name:   "delete",
		Help:   "Deletes a user",
		Exec: func(args ExecArgs) error {
			return fmt.Errorf("Delete must fail")
		},
	}

	show.Options.OptBool("admin", "", "Include admin users")
	delete.Options.ReqInt64("id", "", "The ID of the user to be delted")

	testParser := func(args []string) error {
		output := &bytes.Buffer{}
		parser := NewParser(output)
		parser.AddCmd(show)
		parser.AddCmd(delete)
		return parser.Parse(false, args)
	}

	//Good call
	err := testParser([]string{"a.out", "users", "show"})
	if err != nil {
		t.Errorf("Parse failed: %v", err)
	}

	//Invalid call
	err = testParser([]string{"a.out", "invalid"})
	if err != ErrInvalidArgs {
		t.Errorf("Parse failed: %v", err)
	}
	err = testParser([]string{"a.out", "users", "invalid"})
	if err != ErrInvalidArgs {
		t.Errorf("Parse failed: %v", err)
	}

	//Invalid argument to flags
	err = testParser([]string{"a.out", "users", "delete", "name", "user1"})
	if err != ErrInvalidArgs {
		t.Errorf("Parse failed: %v", err)
	}

	//Good call with flags
	err = testParser([]string{"a.out", "users", "delete", "id", "3"})
	if err.Error() != "Delete must fail" {
		t.Errorf("Parse failed: %v", err)
	}

}
