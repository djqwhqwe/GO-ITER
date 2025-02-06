package main

type IntegerTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func iter[T IntegerTypes](container []T, f func(args ...any)) {
	for i := 0; i < len(container); i++ {
		f(i, container[i])
	}
}

func main() {
	a := []int{1, 2, 45}
	iter(a, func(args ...any) {
		if args[1] == 45 {
			println("ok")
		}
	})
}
