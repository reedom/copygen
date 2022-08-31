// Package copygen contains the setup information for copygen generated code.
package copygen

import (
	"github.com/reedom/copygen/examples/_tests/cyclic/domain"
	"github.com/reedom/copygen/examples/_tests/cyclic/duplicate"
	"github.com/reedom/copygen/examples/_tests/cyclic/models"
)

// Copygen defines the functions that will be generated.
type Copygen interface {
	ModelsToDomain(*models.Account, *models.User) *domain.Account
	DuplicateCyclic(*duplicate.Account) *domain.Account
	SuperCyclic(domain.CyclicInterface) *domain.CyclicInterface
}
