package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type MockModel[T any] interface {
	Insert(v *T) error
	Get(id int64) (*T, error)
	Update(v *T) error
	Delete(id int64) error
}

// Exportable Models struct to wrap all the accessible database models
// Optional, but conveniently gives a single container to hold and represent all the DB models
type Models struct {
	Movies MockModel[Movie]
}

// Uses the DB connection pool that was created upon start-up
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}

func NewMockModels() Models {
	return Models{
		Movies: MockMovieModel{},
	}
}
