// Package copygen contains the setup information for copygen generated code.
package copygen

import (
	"github.com/reedom/copygen/examples/basic/domain"
	"github.com/reedom/copygen/examples/basic/models"
)

// Copygen defines the functions that will be generated.
type Copygen interface {
	Basic(A *models.Account, UserID string) *domain.Account
}
