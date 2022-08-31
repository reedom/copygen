// Package copygen contains the setup information for copygen generated code.
package copygen

import (
	"github.com/reedom/copygen/examples/_tests/duplicate/domain"
	"github.com/reedom/copygen/examples/_tests/duplicate/models"
)

// Copygen defines the functions that will be generated.
type Copygen interface {
	ModelsToDomain(models.Account, models.User) *domain.Account
}
