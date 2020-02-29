package arg

type flag struct {
	name  string
	short string
	help  string
	req   bool
	typ   flagType
	enum  []string
	names []string
}
