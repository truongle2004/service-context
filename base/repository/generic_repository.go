package repository

import (
	"errors"
	"sync"
)

type GenericRepository[T any] struct {
	store map[string]*T
	mu    sync.RWMutex    // use a read-write mutex for better performance
	getID func(*T) string // function to extract ID from entity
}

func NewGenericRepository[T any](getID func(*T) string) *GenericRepository[T] {
	return &GenericRepository[T]{
		store: make(map[string]*T),
		getID: getID,
	}
}

func (r *GenericRepository[T]) Create(entity *T) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := r.getID(entity)
	if _, exists := r.store[id]; exists {
		return errors.New("Data already exists")
	}
	r.store[id] = entity
	return nil
}

func (r *GenericRepository[T]) GetByID(id string) (*T, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	entity, ok := r.store[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return entity, nil
}

func (r *GenericRepository[T]) Update(entity *T) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := r.getID(entity)
	r.store[id] = entity
	return nil
}

func (r *GenericRepository[T]) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.store, id)
	return nil
}

func (r *GenericRepository[T]) List() ([]*T, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*T
	for _, v := range r.store {
		result = append(result, v)
	}
	return result, nil
}
