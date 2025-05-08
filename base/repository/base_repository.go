package repository

type BaseRepository[T any] interface {
	Create(entity *T) error
	GetByID(id string) (*T, error)
	Update(entity *T) error
	Delete(id string) error
	List() ([]*T, error)
}
