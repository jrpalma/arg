package arg

import (
	"bytes"
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {

	show := getShowCmd()
	delete := getDelCmd()

	//show.Option('h', "human", "Human format")
	show.Option('a', "admin", "Include admin users")
	show.ReqInt64('i', "id", "The ID of the user to be delted")
	show.Operand(0, "department", String)

	delete.ReqInt64('i', "id", "The ID of the user to be delted")

	testParser := func(args []string) error {
		output := &bytes.Buffer{}
		parser := NewParser(output)
		parser.AddCmd(show)
		parser.AddCmd(delete)
		return parser.Parse(false, args)
	}

	//Good call

	err := testParser([]string{"a.out", "users", "show", "--admin", "--id", "3", "clothing"})
	if err != nil {
		t.Errorf("Parse failed: %v", err)
	}
	err = testParser([]string{"a.out", "users", "show", "-a", "-i", "3", "clothing"})
	if err != nil {
		t.Errorf("Failed to show: %v", err)
	}
	err = testParser([]string{"a.out", "users", "show", "-ai3", "clothing"})
	if err != nil {
		t.Errorf("Failed to show: %v", err)
	}
	err = testParser([]string{"a.out", "users", "show", "-ai3", "clothing"})
	if err != nil {
		t.Errorf("Failed to show: %v", err)
	}
	err = testParser([]string{"a.out", "users", "show", "-a", "--id=3", "clothing"})
	if err != nil {
		t.Errorf("Failed to show: %v", err)
	}
	//Bad call
	err = testParser([]string{"a.out", "users", "show", "-t"})
	if err == nil {
		t.Errorf("Should fail with -t option")
	}
	err = testParser([]string{"a.out", "users", "show", "--table"})
	if err == nil {
		t.Errorf("Should fail with -t option")
	}
	err = testParser([]string{"a.out", "users", "show", "-"})
	if err == nil {
		t.Errorf("Should fail with dash")
	}
	err = testParser([]string{"a.out", "users", "bad"})
	if err == nil {
		t.Errorf("Should fail with bad command")
	}
	err = testParser([]string{"a.out"})
	if err == nil {
		t.Errorf("Should fail with no args")
	}
	err = testParser([]string{"a.out", "users", "delete", "-i3"})
	if err == nil {
		t.Errorf("Delete must fail")
	}
}

func getShowCmd() *Cmd {
	show := &Cmd{
		Prefix: "users",
		Name:   "show",
		Help:   "Show users",
		Exec: func(args ExecArgs) error {
			var i int64
			if args.HasOption("") {
				return fmt.Errorf("Empty options should fail")
			}
			if !args.HasOption("a") {
				return fmt.Errorf("Admin option should be provided")
			}
			if !args.HasOption("admin") {
				return fmt.Errorf("Admin option should be provided")
			}
			if args.HasOption("bad") {
				return fmt.Errorf("Bad option should fail")
			}
			if args.GetOperand(0, nil) {
				return fmt.Errorf("Invalid position must fail")
			}
			if args.GetFlag("i", nil) {
				return fmt.Errorf("Invalid flag value must fail")
			}
			if !args.GetFlag("i", &i) {
				return fmt.Errorf("ID flag must succeed")
			}
			return nil
		},
	}

	return show
}
func getDelCmd() *Cmd {
	delete := &Cmd{
		Prefix: "users",
		Name:   "delete",
		Help:   "Deletes a user",
		Exec: func(args ExecArgs) error {
			return fmt.Errorf("Delete must fail")
		},
	}
	return delete
}
