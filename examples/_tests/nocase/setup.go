// Generator should strip "//go:build copygen" line.
//go:build copygen

// Package copygen contains the setup information for copygen generated code.
package copygen

// Generator should strip "//go:generate line.
//go:generate go run github.com/reedom/copygen -yml setup.yml

// Copygen defines the functions that will be generated.
type Copygen interface {
	// nocase
	ModelsToDomain(*DatabaseModel) *DomainModel
}
