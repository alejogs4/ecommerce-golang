package valueobject

// ValueObject Abstract interface to be considered a value object, defining the rule of equality
type ValueObject interface {
	Equals(another ValueObject) bool
}
