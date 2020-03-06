package arg

import (
	"fmt"
)

// A user represent a user
type User struct {
	ID      int64
	Name    string
	Enabled bool
	Salary  float64
}

var users = []User{
	{100, "joe", true, 50000},
	{101, "ben", true, 37000},
	{101, "bob", true, 60000},
	{101, "tod", true, 58000},
}

func ShowUser(args ExecArgs) error {
	var id int64

	if !args.GetOption("-id", &id) {
		return fmt.Errorf("Expected uint64 -id flag")
	}

	for i := range users {
		if users[i].ID == id {
			fmt.Printf("%+v\n", users[i])
			break
		}
	}
	return nil
}
func DisableUser(args ExecArgs) error {
	var id int64

	if !args.GetOption("-id", &id) {
		return fmt.Errorf("Expected uint64 -id flag")
	}
	for i := range users {
		if users[i].ID == id {
			users[i].Enabled = false
			fmt.Printf("Disabled user with ID %v\n", id)
			break
		}
	}
	return nil
}

// ExampleUsage Shows how this package can be used
func Example_usage() {

	showCmd := &Cmd{
		Prefix: "users",
		Name:   "show",
		Help:   "Displays all the users",
		Exec:   ShowUser,
	}
	showCmd.ReqInt64('i', "id", "The user ID to be shown")

	disableCmd := &Cmd{
		Prefix: "users",
		Name:   "disable",
		Help:   "Disables a user by the given ID",
		Exec:   DisableUser,
	}
	disableCmd.ReqInt64('i', "id", "The user ID to be disabled")

	//The parser can have a nil output writer.
	//We could use os.Stdout, but that will
	//mess with the testable output.
	parser := NewParser(nil)
	parser.AddCmd(showCmd)
	parser.AddCmd(disableCmd)

	showArgs := []string{"a.out", "users", "show", "-id", "100"}
	disableArgs := []string{"a.out", "users", "disable", "-id", "100"}
	badArgs := []string{"a.out", "unknown"}

	//Should not fail since the arguments are fine
	parser.Parse(false, showArgs)
	parser.Parse(false, disableArgs)

	//Should fail with unknown argument
	err := parser.Parse(false, badArgs)
	if err == ErrInvalidArgs {
		fmt.Printf("%v\n", err)
	}

	/*
		// Output: {ID:100 Name:joe Enabled:true Salary:50000}
		// Disabled user with ID 100
		// Invalid Arguments
	*/
}
