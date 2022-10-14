package procedures

type Args struct {
	A, B string
}

type Jobs int
type Result string

func (t *Jobs) Multiply(args Args, result *Result) error {
	return Multiply(args, result)
}

func Multiply(args Args, result *Result) error {
	*result = Result(args.A + "+" + args.B)
	return nil
}
