package singleton

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

type single struct{}

var singleInstance *single

func Run() {
	for range 10 {
		go getInstance()
	}
	time.Sleep(1 * time.Second)
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
