package gopad

// FromPtr transform input from pointer.
func FromPtr[T any](v *T) T {
	return *v
}

// ToPtr transform input to a pointer.
func ToPtr[T any](v T) *T {
	return &v
}
