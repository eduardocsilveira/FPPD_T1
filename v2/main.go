package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	N           = 5
	ciclos      = 20
	tempoPensar = 50 * time.Millisecond
	tempoComer  = 50 * time.Millisecond
)

var garfos [N]sync.Mutex
var mesa sync.Mutex

var mu sync.Mutex
var somaEspera, maiorEspera time.Duration
var comendo, maxComendo int
var ciclosFeitos [N]int

func comecouComer(espera time.Duration) {
	mu.Lock()
	somaEspera += espera
	if espera > maiorEspera {
		maiorEspera = espera
	}
	comendo++
	if comendo > maxComendo {
		maxComendo = comendo
	}
	mu.Unlock()
}

func terminouComer(id int) {
	mu.Lock()
	comendo--
	ciclosFeitos[id]++
	mu.Unlock()
}

func filosofo(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	esq := id
	dir := (id + 1) % N

	for c := 0; c < ciclos; c++ {
		time.Sleep(tempoPensar)

		pedido := time.Now()
		mesa.Lock()
		garfos[esq].Lock()
		garfos[dir].Lock()
		comecouComer(time.Since(pedido))

		time.Sleep(tempoComer)
		terminouComer(id)

		garfos[dir].Unlock()
		garfos[esq].Unlock()
		mesa.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	inicio := time.Now()
	for id := 0; id < N; id++ {
		wg.Add(1)
		go filosofo(id, &wg)
	}
	wg.Wait()

	fmt.Println("V2 - trava da mesa")
	fmt.Println("tempo total:", time.Since(inicio))
	fmt.Println("espera media:", somaEspera/(N*ciclos))
	fmt.Println("maior espera:", maiorEspera)
	fmt.Println("max comendo juntos:", maxComendo)
	fmt.Println("ciclos por filosofo:", ciclosFeitos)
}
