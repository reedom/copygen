// Package copygen contains the setup information for copygen generated code.
package copygen

import (
	c "strconv"

	"github.com/reedom/copygen/examples/error/domain"
	"github.com/reedom/copygen/examples/error/models"
)

// Copygen defines the functions that will be generated.
type Copygen interface {
	// custom see table in the README for options
	ModelsToDomain(*models.Account, *models.User) (*domain.Account, error)
}

/* Define the function and field this converter is applied to using regex. */
// convert .* models.User.UserID error
// Itoa converts an integer to an ascii value.
func Itoa(i int) (string, error) {
	return c.Itoa(i), nil
}
