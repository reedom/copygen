// Code generated by github.com/reedom/copygen
// DO NOT EDIT.

// Package copygen contains the setup information for copygen generated code.
package copygen

// ModelsToDomain copies a *DatabaseModel to a *DomainModel.
func ModelsToDomain(tD *DomainModel, fD *DatabaseModel) {
	// *DomainModel fields
	tD.id = fD.ID
}
