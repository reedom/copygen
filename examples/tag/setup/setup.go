// Package copygen contains the setup information for copygen generated code.
package copygen

import (
	"github.com/reedom/copygen/examples/tag/domain"
	"github.com/reedom/copygen/examples/tag/models"
)

// Copygen defines the functions that will be generated.
type Copygen interface {
	// tag .* api
	ModelsToDomain(*models.Account, *models.User) *domain.Account
}
