package algorithm

type Comp interface {
	Greater(other Comp) bool
	Equal(other Comp) bool
}

type MyInt int

func (n MyInt) Greater(o Comp) bool {
	val := o.(MyInt)
	return n > val
}

func (n MyInt) Equal(o Comp) bool {
	val := o.(MyInt)
	return n == val
}
