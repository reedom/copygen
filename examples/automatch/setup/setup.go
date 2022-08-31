// Package copygen contains the setup information for copygen generated code.
package copygen

import (
	"github.com/reedom/copygen/examples/automatch/domain"
	"github.com/reedom/copygen/examples/automatch/models"
)

// Copygen defines the functions that will be generated.
type Copygen interface {
	// depth domain.Account 2
	// depth models.User 1
	ModelsToDomain(*models.Account, *models.User) *domain.Account
}
