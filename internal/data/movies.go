package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"` // in minutes
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"` // starts at 1 and increments at every update of the movie
}

// could also do MarshalJSON() ([]byte, error) here
// but might lose control of field order
