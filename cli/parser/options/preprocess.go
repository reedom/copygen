package options

import (
	"fmt"
	"strings"
)

const (
	CategoryPreProcess = "preprocess"

	// FormatPreProcess represents an end-user facing format for preProcess options.
	// <option> refers to the "preprocess" option.
	FormatPreProcess = "<option>:<whitespaces><ident>"
)

// ParsePreProcess parses a preprocess option.
func ParsePreProcess(option string) (*Option, error) {
	splitoption := strings.Fields(option)
	if len(splitoption) != 1 {
		return nil, fmt.Errorf("there is a misconfigured %s option: %q.\nIs it in format %s?", CategoryPreProcess, option, FormatPreProcess)
	}

	return &Option{
		Category: CategoryPreProcess,
		Value:    splitoption[0],
	}, nil
}
