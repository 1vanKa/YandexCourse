package main

import (
	"fmt"
	"time"
	"yandexcourse/hashbyte"
	"yandexcourse/randbyte"
)

func main() {

	// создаём генератор случайных чисел
	generator := randbyte.New(time.Now().UnixNano()) // в качестве затравки передаём ему текущее время, и при каждом запуске оно будет разным.

	buf := make([]byte, 16)

	for i := 0; i < 5; i++ {
		n, _ := generator.Read(buf) // единственный доступный метод, но он нам и нужен.
		fmt.Printf("Generate bytes: %v size(%d)\n", buf, n)
	}

	hasher := hashbyte.New(0)
	hasher.Write(buf)
	fmt.Printf("Hash: %v \n", hasher.Hash())
}
