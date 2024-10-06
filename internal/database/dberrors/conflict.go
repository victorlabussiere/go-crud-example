package dberrors

type ConflictError struct{}

func (e *ConflictError) Error() string {
	return "Attempt to create a record with an existing key"
}
