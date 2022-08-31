// Package copygen contains the setup information for copygen generated code.
package copygen

import (
	"github.com/reedom/copygen/examples/automatch/domain"
	"github.com/reedom/copygen/examples/automatch/models"
)

// Copygen defines the functions that will be generated.
type Copygen interface {
	// map models.Account.ID domain.Account.ID
	// map models.Account.Name domain.Account.Name
	// automatch domain.Account.*
	ModelsToDomain(*models.Account, *models.User) *domain.Account
}
