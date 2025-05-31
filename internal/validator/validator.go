package validator

import (
	"regexp"
	"slices"
)

type Validators struct {
	Errors map[string]string
}
