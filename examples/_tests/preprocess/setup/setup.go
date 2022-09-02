// Package copygen contains the setup information for copygen generated code.
package copygen

import (
	"github.com/switchupcb/copygen/examples/automatch/domain"
	"github.com/switchupcb/copygen/examples/automatch/models"
)

// Copygen defines the functions that will be generated.
type Copygen interface {
	// preprocess _process
	Func1(*models.Account) *domain.Account

	// preprocess _processErr
	Func2(*models.Account, *models.User) (*domain.Account, error)

	// postprocess _process
	Func3(*models.Account) *domain.Account

	// postprocess _processErr
	Func4(*models.Account, *models.User) (*domain.Account, error)
}

func _process(*domain.Account, *models.Account) {
}

func _processErr(*domain.Account, *models.Account, *models.User) error {
	return nil
}
