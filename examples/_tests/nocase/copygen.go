// Code generated by github.com/reedom/copygen
// DO NOT EDIT.

// Generator should strip "//go:build copygen" line.

// Package copygen contains the setup information for copygen generated code.
package copygen

// Generator should strip "//go:generate line.

// ModelsToDomain copies a *DatabaseModel to a *DomainModel.
func ModelsToDomain(tD *DomainModel, fD *DatabaseModel) {
	// *DomainModel fields
	tD.id = fD.ID
}
