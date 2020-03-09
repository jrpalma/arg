package arg

// DataType represent a data type
type dataType int

const (
	typeNone dataType = iota
	typeEnum
	typeBool
	typeInt64
	typeUint64
	typeString
	typeFloat64
)
