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
