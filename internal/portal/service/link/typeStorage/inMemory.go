package typeStorage

import (
	"ShortLand/internal/common/model"
	"fmt"
	"sync"
)

type InMemory struct {
	mu    sync.RWMutex
	store map[string]*model.LinkTable
}

func NewInMemory() *InMemory {
	return &InMemory{
		store: make(map[string]*model.LinkTable),
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
