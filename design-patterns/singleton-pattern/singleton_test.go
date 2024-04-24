package singleton_pattern

import (
	"fmt"
	"sync"
	"testing"
)

func TestWithLock(t *testing.T) {
	var wg sync.WaitGroup
	k := 50
	wg.Add(k)
	for i := 0; i < k; i++ {
		go func() {
			_ = NewDbRepositoryWithLock()
			wg.Done()
		}()
	}

	wg.Wait()
}

func TestWithOnce(t *testing.T) {
	var wg sync.WaitGroup
	k := 50
	wg.Add(k)
	for i := 0; i < k; i++ {
		go func() {
			db := NewDbRepositoryWithOnce()
			fmt.Println(db)
			wg.Done()
		}()
	}

	wg.Wait()
}
