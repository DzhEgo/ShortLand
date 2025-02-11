package typeStorage

import (
	"ShortLand/internal/common/model"
	"fmt"
	"sync"
	"time"
)

type InMemory struct {
	mu    sync.RWMutex
	store map[string]*model.LinkTable
	quit  chan struct{}
}

func NewInMemory() *InMemory {
	return &InMemory{
		store: make(map[string]*model.LinkTable),
		quit:  make(chan struct{}),
	}
}

func (i *InMemory) SaveLink(data model.LinkTable) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.store[data.ShortLink] = &data
	return nil
}

func (i *InMemory) GetLink(shortLink string) (*model.LinkTable, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	data, exists := i.store[shortLink]
	if !exists {
		return nil, fmt.Errorf("failed to find link")
	}

	return data, nil
}

func (i *InMemory) DeleteLink(shortLink string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	delete(i.store, shortLink)
	return nil
}

func (i *InMemory) Cleanup(t time.Duration) {
	ticker := time.NewTicker(t)

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				i.cleanExpired()
			case <-i.quit:
				return
			}
		}
	}()
}

func (i *InMemory) Close() {
	close(i.quit)
}

func (i *InMemory) cleanExpired() {
	i.mu.Lock()
	defer i.mu.Unlock()
	
	for k, v := range i.store {
		if time.Now().Unix() > v.ExpireAt {
			delete(i.store, k)
		}
	}
}
