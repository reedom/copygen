// Package copygen contains the setup information for copygen generated code.
package copygen

import (
	service "github.com/reedom/copygen/examples/main/domain"
	data "github.com/reedom/copygen/examples/main/models"
)

// Copygen defines the functions that will be generated.
type Copygen interface {
	ModelsToDomain(*data.Account, *data.User) *service.Account
}
