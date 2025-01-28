package repository

import "errors"

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// album 저장소 인터페이스
type Repository interface {
	GetAll() ([]Album, error)
	GetByID(id string) (*Album, error)
	Save(album Album) error
}

// In-Memory 저장소 구현
type MemoryRepository struct {
	albums []Album
}

func NewMemoryRepository() *MemoryRepository {
	var albums = []Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	return &MemoryRepository{albums: albums}
}

func (r *MemoryRepository) GetAll() ([]Album, error) {
	return r.albums, nil
}

func (r *MemoryRepository) GetByID(id string) (*Album, error) {
	for _, a := range r.albums {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("album not found")
}

func (r *MemoryRepository) Save(album Album) error {
	r.albums = append(r.albums, album)
	return nil
}
