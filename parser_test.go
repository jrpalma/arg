package arg

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDebug(t *testing.T) {
	cmd := &Cmd{Name: "name"}
	cmd.Option('y', "verbose", "Verbose output")
	err := testCmd(cmd, []string{"a.out", "name", "x"})
	if err == nil {
		t.Errorf("Command with no exec should not fail: %v", err)
	}
	err = testCmd(cmd, []string{"a.out", "x"})
	if err == nil {
		t.Errorf("Command with no exec should not fail: %v", err)
	}
}

func TestUsage(t *testing.T) {

	showCmd := getShowCmd()
	showCmd.ReqString(0, "format", "The format to be used")
	err := testCmd(showCmd, []string{
		"a.out", "users", "showCmd", "--admin", "--id", "3"})
	if err == nil {
		t.Errorf("Parse should fail")
	}

}
func TestNegativeArg(t *testing.T) {
	operCmd := getOperCmd()
	err := testCmd(operCmd, []string{
		"a.out", "oper", "string", "bad", "3", "3", "3.0"})
	if err == nil {
		t.Errorf("Parse should fail")
	}

	operCmd.Float64Operand(6, "Float64")
	err = testCmd(operCmd, []string{
		"a.out", "oper", "string", "true", "3", "3", "3.0", "3.0"})
	if err == nil {
		t.Errorf("Parse should fail")
	}

	operCmd.StringOperand(5, "Float64")
	err = testCmd(operCmd, []string{
		"a.out", "oper", "string", "true", "3", "3", "3.0", "3.0", "3.0", "3.0"})
	if err == nil {
		t.Errorf("Parse should fail")
	}
}
func TestPositiveOperands(t *testing.T) {
	operCmd := getOperCmd()
	err := testCmd(operCmd, []string{
		"a.out", "oper", "string", "true", "3", "3", "3.0"})
	if err != nil {
		t.Errorf("Parse failed: %v", err)
	}
}
func TestParserSlices(t *testing.T) {
	searchCmd := getSearchCmd()
	err := testCmd(searchCmd, []string{
		"a.out", "search", "--include", "/boot", "--include=/home", "-I/root", "-I", "/opt"})
	if err != nil {
		t.Errorf("Parse failed: %v", err)
	}
}
func TestParserPositive(t *testing.T) {
	showCmd := getShowCmd()

	err := testCmd(showCmd, []string{
		"a.out", "users", "showCmd", "--admin", "--id", "3", "clothing"})
	if err != nil {
		t.Errorf("Failed to showCmd: %v", err)
	}
	err = testCmd(showCmd, []string{
		"a.out", "users", "showCmd", "-a", "-i", "3", "clothing"})
	if err != nil {
		t.Errorf("Failed to showCmd: %v", err)
	}
	err = testCmd(showCmd, []string{
		"a.out", "users", "showCmd", "-ai3", "clothing"})
	if err != nil {
		t.Errorf("Failed to showCmd: %v", err)
	}
	err = testCmd(showCmd, []string{
		"a.out", "users", "showCmd", "-ai3", "clothing"})
	if err != nil {
		t.Errorf("Failed to showCmd: %v", err)
	}
	err = testCmd(showCmd, []string{
		"a.out", "users", "showCmd", "-a", "--id=3", "clothing"})
	if err != nil {
		t.Errorf("Failed to showCmd: %v", err)
	}
}

func TestParserNegative(t *testing.T) {

	showCmd := getShowCmd()
	deleteCmd := getDelCmd()

	err := testCmd(showCmd, []string{"a.out", "users", "showCmd", "-t"})
	if err == nil {
		t.Errorf("Should fail with -t option")
	}
	err = testCmd(showCmd, []string{"a.out", "users", "showCmd", "--table"})
	if err == nil {
		t.Errorf("Should fail with -t option")
	}
	err = testCmd(showCmd, []string{"a.out", "users", "showCmd", "-"})
	if err == nil {
		t.Errorf("Should fail with dash")
	}
	err = testCmd(showCmd, []string{"a.out", "users", "bad"})
	if err == nil {
		t.Errorf("Should fail with bad command")
	}
	err = testCmd(showCmd, []string{"a.out"})
	if err == nil {
		t.Errorf("Should fail with no args")
	}
	err = testCmd(showCmd, []string{
		"a.out", "users", "showCmd", "--admin", "--id", "3"})
	if err == nil {
		t.Errorf("Parse should fail with invalid department")
	}
	err = testCmd(deleteCmd, []string{"a.out", "users", "delete", "-i3"})
	if err == nil {
		t.Errorf("Delete must fail")
	}
	err = testCmd(deleteCmd, []string{"a.out", "users", "delete"})
	if err == nil {
		t.Errorf("Delete must fail")
	}
}

func getShowCmd() *Cmd {
	showCmd := &Cmd{
		Prefix:      "users",
		Name:        "showCmd",
		Description: "Show users",
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
			if args.GetOption("i", nil) {
				return fmt.Errorf("Invalid flag value must fail")
			}
			if !args.GetOption("i", &i) {
				return fmt.Errorf("ID flag must succeed")
			}
			return nil
		},
	}
	showCmd.Option('a', "admin", "Include admin users")
	showCmd.ReqInt64('i', "id", "The ID of the user to be delted")
	showCmd.StringOperand(0, "department")

	return showCmd
}
func getDelCmd() *Cmd {
	deleteCmd := &Cmd{
		Prefix:      "users",
		Name:        "delete",
		Description: "Deletes a user",
		Exec: func(args ExecArgs) error {
			return fmt.Errorf("Delete must fail")
		},
	}
	deleteCmd.ReqInt64('i', "id", "The ID of the user to be delted")
	return deleteCmd
}
func getSearchCmd() *Cmd {
	searchCmd := &Cmd{
		Prefix:      "",
		Name:        "search",
		Description: "Search files",
		Exec: func(args ExecArgs) error {
			var boot string
			var all []string
			if !args.GetOption("I", &boot) && boot != "/boot" {
				return fmt.Errorf("Invalid first flag value %v", &boot)
			}
			if !args.GetOption("I", &all) && len(all) != 4 {
				return fmt.Errorf("Invalid flag values %#v", all)
			}
			return nil
		},
	}
	searchCmd.OptString('I', "include", "Directories to include")
	return searchCmd
}
func getOperCmd() *Cmd {
	operCmd := &Cmd{
		Prefix:      "",
		Name:        "oper",
		Description: "Test the operands",
		Exec: func(args ExecArgs) error {
			//Test some basic cases here
			if args.GetOption("", nil) {
				return fmt.Errorf("Should fail with empty")
			}
			if args.GetOption("x", nil) {
				return fmt.Errorf("Should fail with x option")
			}
			return nil
		},
	}
	operCmd.StringOperand(0, "String")
	operCmd.BoolOperand(1, "Bool")
	operCmd.Int64Operand(2, "Int64")
	operCmd.Uint64Operand(3, "Uint64")
	operCmd.Float64Operand(4, "Float64")
	return operCmd
}
func testCmd(cmd *Cmd, args []string) error {
	output := &bytes.Buffer{}
	parser := NewParser(output)
	parser.AddCmd(cmd)
	return parser.Parse(false, args)
}
