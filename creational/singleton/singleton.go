package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct{}

var singleInstance *single

func Run() {
	// Limpa qualquer instância anterior para garantir um teste limpo
	singleInstance = nil

	var wg sync.WaitGroup
	fmt.Println("Starting singleton test with 10 concurrent requests...")

	for i := range 10 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			instance := getInstance()
			fmt.Printf("Goroutine %d got instance: %p\n", id, instance)
		}(i) // Passa o índice para a goroutine
	}

	wg.Wait()
	fmt.Println("All goroutines completed.")
}

func getInstance() *single {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
