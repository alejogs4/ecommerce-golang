package valueobject

import "strings"

// MaybeEmpty is a string which maybe could be empty
type MaybeEmpty string

// NewMaybeEmpty returns a new MaybeEmpty object sanitizing the underlying string
func NewMaybeEmpty(value string) MaybeEmpty {
	sanitizedValue := strings.TrimSpace(value)
	return MaybeEmpty(sanitizedValue)
}

// IsEmpty returns true is underlying string is indeed empty
func (ne MaybeEmpty) IsEmpty() bool {
	return len(string(ne)) == 0
}

// Equals returns if both values are equals
func (ne MaybeEmpty) Equals(other MaybeEmpty) bool {
	return string(ne) == string(other)
}

func (ne MaybeEmpty) String() string {
	return string(ne)
}
