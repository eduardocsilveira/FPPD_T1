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

func filosofo(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	esq := id
	dir := (id + 1) % N

	for c := 0; c < ciclos; c++ {
		time.Sleep(tempoPensar)

		garfos[esq].Lock()
		fmt.Printf("F%d pegou o garfo %d\n", id, esq)
		time.Sleep(5 * time.Millisecond)

		fmt.Printf("F%d esperando o garfo %d\n", id, dir)
		garfos[dir].Lock()

		time.Sleep(tempoComer)

		garfos[dir].Unlock()
		garfos[esq].Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	for id := 0; id < N; id++ {
		wg.Add(1)
		go filosofo(id, &wg)
	}
	wg.Wait()
	fmt.Println("terminou sem deadlock")
}
