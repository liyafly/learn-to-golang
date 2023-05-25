package singleton

import (
	"fmt"
	"sync"
)

type Singleton struct {
	Value int
}

func (s *Singleton) Show() {

	fmt.Println(s.Value)
	fmt.Println("hello world")
}

var (
	once   sync.Once
	single *Singleton
)

func GetSingleInstance() *Singleton {
	once.Do(func() {
		single = &Singleton{Value: 1}
	})
	return single
}
