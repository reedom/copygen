// DO NOT CHANGE PACKAGE

// Package templates provides a template used by copygen to generate custom code.
package templates

import (
	"github.com/switchupcb/copygen/cli/models"
)

// Header provides the header of the generated code.
// GENERATOR FUNCTION
// EDITABLE.
// DO NOT REMOVE.
func Header(gen models.Generator) string {
	return DefaultHeader(gen)
}

// DefaultHeader provides the header of the generated file using the default method.
func DefaultHeader(gen models.Generator) string {
	var header string

	// package
	header += "// Code generated by github.com/switchupcb/copygen\n"
	header += "// DO NOT EDIT.\n"
	header += "\n"
	header += "package " + gen.Package + "\n"

	// imports
	header += "import (\n"
	for _, iprt := range gen.Imports {
		header += "\"" + iprt + "\"\n"
	}
	header += ")"
	return header
}
