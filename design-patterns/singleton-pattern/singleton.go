package singleton_pattern

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DbRepository struct {
	DbName string
	ConnectionCount int
}

var instanceLock sync.Mutex
var instance *DbRepository = nil

func NewDbRepositoryWithLock() *DbRepository {
	if instance == nil {
		instanceLock.Lock()
		defer instanceLock.Unlock()
		if instance == nil {
			random := rand.Intn(3)
			time.Sleep(time.Duration(random) * time.Second)
			instance = &DbRepository{"postgres", 0}
			fmt.Println("Instance created")
		} else {
			fmt.Println("Returning existing instance")
		}
	} else {
		fmt.Println("Returning existing instance")
	}

	return instance
}

var once sync.Once
func NewDbRepositoryWithOnce() *DbRepository {
	once.Do(func() {
		var t time.Duration = 10
		time.Sleep(t * time.Second)
		instance = &DbRepository{"postgres", 0}
		fmt.Println("Instance created")
	})

	return instance
}