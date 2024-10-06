package dberrors

import "fmt"

type NotFoundError struct {
	Entity string
	ID     uint
}

func (n *NotFoundError) Error() string {
	return fmt.Sprintf("Unable to fin %s with id %d", n.Entity, n.ID)
}
