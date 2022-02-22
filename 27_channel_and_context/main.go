package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	storage map[string]int
	mu      sync.RWMutex
}

func (c *Cache) Increase(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] += value
}

func (c *Cache) Set(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] = value
}

func (c *Cache) Get(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.storage[key]
}

func (c *Cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.storage, key)
}

const (
	k1   = "key1"
	step = 7
)

func main() {
	cache := Cache{storage: make(map[string]int)}

	semaphore := make(chan struct{}, 4)
	defer close(semaphore)
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*100)

	go func() {
		select {
		case semaphore <- struct{}{}:
			defer func() {
				<-semaphore
			}()
			cache.Increase(k1, step)
		case <-ctx.Done():
			fmt.Println("context deadline exceeded")
			return
		default:
			time.Sleep(time.Millisecond * 20)
		}
	}()

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			select {
			case semaphore <- struct{}{}:
				defer func() {
					<-semaphore
				}()
				cache.Set(k1, step*i)
			case <-ctx.Done():
				fmt.Println("context deadline exceeded")
				return
			default:
				time.Sleep(time.Millisecond * 20)
			}
		}()
	}

	for len(semaphore) > 0 {
		time.Sleep(time.Millisecond * 10)
	}
	fmt.Println(cache.Get(k1))
}
