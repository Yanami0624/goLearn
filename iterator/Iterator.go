package iterator

func Exp_push[T int | float64](n T) func(yield func(T) bool) {
	var a T = 1
	return func(yield func(T) bool) {
		for {
			a *= n
			if !yield(a) {
				return
			}
		}
	}
}

type Exp_pull[T int | float64] struct {
	a T
	n T
}

func NewExp[T int | float64](num T) *Exp_pull[T] {
	exp := new(Exp_pull[T])
	exp.a = 1
	exp.n = num
	return exp
}

func (exp *Exp_pull[T]) Next() T {
	ret := exp.a
	exp.a *= exp.n
	return ret
}

func (exp Exp_pull[T]) Stop(limit T) bool {
	return exp.a > limit
}

func (exp Exp_pull[T]) Val() T {
	return exp.a
}
