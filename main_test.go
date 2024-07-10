package checker

type baseCase struct {
	name  string
	arg   any
	want  bool
	panic bool
}
