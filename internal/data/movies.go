package data

import(
	"time"
	"greenlight.damkols.net/internal/validator"
)

type Movie struct {
	ID int64 `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title string `json:"title"`
	Year int32 `json:"year,omitzero"`
	Runtime Runtime `json:"runtime,omitzero,string"`
	Genres []string `json:"genres,omitempty"`
	Version int32 `json:"version"`
}

func ValidateMovie(v "*validator.Validator", movie *Movie) {
	
}