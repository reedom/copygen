// Code generated by github.com/switchupcb/copygen
// DO NOT EDIT.

// Package copygen contains the setup information for copygen generated code.
package copygen

import (
	"github.com/switchupcb/copygen/examples/automatch/domain"
	"github.com/switchupcb/copygen/examples/automatch/models"
)

// ModelsToDomain copies a models.Account, models.User to a domain.Account.
func ModelsToDomain(tA *domain.Account, fA *models.Account, fU *models.User) {
	// domain.Account fields
	tA.ID = fA.ID
	tA.Name = fA.Name
	tA.Email = fA.Email
	tA.User.UserID = fU.UserID
	tA.User.Username = fU.Username
}
