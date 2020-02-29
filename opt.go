package arg

type optionType int

const (
	unknownOption optionType = iota
	stringOption
	boolOption
	int64Option
	uint64Option
	float64Option
	enumOption
	argsOption
)

type option struct {
	name       string
	short      string
	help       string
	required   bool
	optionType optionType
	enum       []string
	args       []string
}
