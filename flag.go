package arg

type flagType int

const (
	unknownFlag flagType = iota
	stringFlag
	boolFlag
	int64Flag
	uint64Flag
	float64Flag
	enumFlag
	argsFlag
)

type flag struct {
	name  string
	short string
	help  string
	req   bool
	typ   flagType
	enum  []string
	names []string
}
