package kvsvc

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type KvStoreService struct {
	storeMap map[string]string
	filter   map[string]func(key string)
	mu       sync.Mutex
}

func NewKvStoreService() *KvStoreService {
	return &KvStoreService{
		storeMap: make(map[string]string),
		filter:   make(map[string]func(key string)),
	}
}

func (p *KvStoreService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if v, ok := p.storeMap[key]; ok {
		*value = v
		return nil
	}
	return fmt.Errorf("not found")
}

func (p *KvStoreService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	key, value := kv[0], kv[1]
	if oldValue := p.storeMap[key]; oldValue != value {
		for _, fn := range p.filter {
			fn(key)
		}
	}
	p.storeMap[key] = value
	return nil
}

func (p *KvStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10)
	p.mu.Lock()
	p.filter[id] = func(key string) { ch <- key }
	p.mu.Unlock()
	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("time out")
	case key := <-ch:
		*keyChanged = key
		return nil
	}
}
