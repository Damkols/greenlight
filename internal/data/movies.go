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
		v.Check(input.Title != "", "title", "must be provided")
	v.Check(len(input.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(input.Year != 0, "year", "must be provided")
	v.Check(input.Year >= 1888, "year", "must be greater than 1888")
	v.Check(input.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(input.Runtime != 0, "runtime", "must be provided")
	v.Check(input.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(input.Genres != nil, "genres", "must be provided")
	v.Check(len(input.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(input.Genres) <= 5, "genres", "must not contain more than 5 genres")

	v.Check(validator.Unique(input.Genres), "genres", "must not contain duplicate values")

}