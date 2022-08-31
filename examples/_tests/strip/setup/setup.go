// Generator should strip either "//go:build ignore" or "go:build exclude" line.
//go:build ignore

// Package copygen contains the setup information for copygen generated code.
package copygen

import (
	service "github.com/reedom/copygen/examples/main/domain"
	data "github.com/reedom/copygen/examples/main/models"
)

// Generator should strip "//go:generate line.
//go:generate go run github.com/switchupcb/copygen -yml setup.yml

// Copygen defines the functions that will be generated.
type Copygen interface {
	ModelsToDomain(*data.Account) *service.Account
}
