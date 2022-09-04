// Package models contains data storage models (i.e database).
package copygen

// DomainModel represents a domain model.
type DomainModel struct {
	id int
}

func (d DomainModel) ID() int {
	return d.id
}

type DatabaseModel struct {
	ID int
}
