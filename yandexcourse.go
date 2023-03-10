package main

import (
	"fmt"
)

type input struct {
	a       int
	b       int
	counter int
}

// 1) добавьте функцию f
// 2) добавьте функцию test(a, b, counter int) (err error),
// которая преобразует панику в ошибку
// ...

func f(a, b, i int) {
	if i == 0 {
		panic("counter equals 0")
	}
	f(b, a/b, i-1)
}

func test(a, b, counter int) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("%v", p)
		}
	}()
	f(a, b, counter)
	return nil
}

func main() {
	for _, pars := range []input{
		{10, 5, 3},
		{100, 7, 10},
		{1, 1, 1000},
	} {
		fmt.Printf("(%d, %d, %d) => %v\r\n", pars.a, pars.b, pars.counter,
			test(pars.a, pars.b, pars.counter))
	}
}
