package tool

type Number interface {
	int | int16 | int32 | int64
}

func Sum[T Number](a, b T) T {
	return a + b
}

func Equal[T Number](a, b T) bool {
	return a == b
}