// Package copygen contains the setup information for copygen generated code.
package copygen

// Copygen defines the functions that will be generated.
type Copygen interface {
	// nocase
	ModelsToDomain(*DatabaseModel) *DomainModel
}
