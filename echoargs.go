package echoargs

type Echoer interface {
	Echo(Args) error
}

func NewArgs(vals ...string) Args {
	return Args{
		vals: vals,
	}
}

type Args struct {
	vals []string
}

func (a Args) Values() []string {
	return a.vals
}
