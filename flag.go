package arg

type flag struct {
	name  string
	short string
	help  string
	req   bool
	typ   flagType
	enum  []string
	count uint
}

func (f *flag) Name() string   { return f.name }
func (f *flag) Short() string  { return f.short }
func (f *flag) Help() string   { return f.help }
func (f *flag) Required() bool { return f.req }
func (f *flag) Type() flagType { return f.typ }
func (f *flag) Count() uint    { return f.count }
