package echoargs

func NewArgs(first string, rest ...string) Args {
	vals := make([]string, 1+len(rest))
	vals[0] = first
	for i := 1; i < len(vals); i++ {
		vals[i] = rest[i]
	}

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
