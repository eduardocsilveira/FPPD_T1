package main

import (
	"fmt"
	"time"
)

const (
	N           = 5
	ciclos      = 20
	tempoPensar = 50 * time.Millisecond
	tempoComer  = 50 * time.Millisecond
)

func main() {
	var ciclosFeitos [N]int

	inicio := time.Now()
	for id := 0; id < N; id++ {
		for c := 0; c < ciclos; c++ {
			time.Sleep(tempoPensar)
			time.Sleep(tempoComer)
			ciclosFeitos[id]++
		}
	}

	fmt.Println("V0 - sequencial")
	fmt.Println("tempo total:", time.Since(inicio))
	fmt.Println("ciclos por filosofo:", ciclosFeitos)
}
