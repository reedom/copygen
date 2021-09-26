// Code generated by github.com/switchupcb/copygen
// DO NOT EDIT.
package copygen

import (
	"github.com/switchupcb/copygen/examples/automatch/domain"
	"github.com/switchupcb/copygen/examples/automatch/models"
)

// ModelsToDomain copies a Account, T, User to a Account.
func ModelsToDomain(tA domain.Account, fA models.Account, fT domain.T, fU models.User) {
	// Account fields
	tA.ID = fA.ID
	tA.Name = fA.Name
	tA.Email = fA.Email

}
