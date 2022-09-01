// Code generated by github.com/reedom/copygen
// DO NOT EDIT.

// Package copygen contains the setup information for copygen generated code.
package copygen

import (
	"fmt"
	c "strconv"

	"github.com/reedom/copygen/examples/error/domain"
	"github.com/reedom/copygen/examples/error/models"
)

/* The paring converter function can return an error value. */
// Itoa converts an integer to an ascii value.
func Itoa(i int) (string, error) {
	if i < 1 {
		return "", fmt.Errorf("invalid id")
	}
	return c.Itoa(i), nil
}

// ModelsToDomain copies a *models.Account, *models.User to a *domain.Account.
func ModelsToDomain(tA *domain.Account, fA *models.Account, fU *models.User) (err error) {
	// *domain.Account fields
	tA.ID = fA.ID
	tA.UserID, err = Itoa(fU.UserID)
	if err != nil {
		return
	}
	tA.Name = fA.Name

	return
}
