package options

import (
	"fmt"
	"strings"
)

const (
	CategoryNoCase = "nocase"

	// FormatNoCase represents an end-user facing format for nocase options.
	// <option> refers to the "nocase" option.
	FormatNoCase = "<option>"
)

// ParseNoCase parses a noCase option.
func ParseNoCase(option string) (*Option, error) {
	if strings.TrimSpace(option) != "" {
		return nil, fmt.Errorf("there is a misconfigured %s option: %q.\nIs it in format %s?", CategoryNoCase, option, FormatNoCase)
	}

	return &Option{
		Category: CategoryNoCase,
	}, nil
}
