// Package copygen contains the setup information for copygen generated code.
package copygen

import (
	"fmt"
	c "strconv"

	"github.com/reedom/copygen/examples/error/domain"
	"github.com/reedom/copygen/examples/error/models"
)

// Copygen defines the functions that will be generated.
type Copygen interface {
	// To create a function that returns error, place "error" at the end of return values.
	ModelsToDomain(*models.Account, *models.User) (*domain.Account, error)
}

/* The paring converter function can return an error value. */
// convert .* models.User.UserID
// Itoa converts an integer to an ascii value.
func Itoa(i int) (string, error) {
	if i < 1 {
		return "", fmt.Errorf("invalid id")
	}
	return c.Itoa(i), nil
}
