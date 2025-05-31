package validator

import (
	"regexp"
	"slices"
)

type Validators struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}