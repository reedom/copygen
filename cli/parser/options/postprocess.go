package options

import (
	"fmt"
	"strings"
)

const (
	CategoryPostProcess = "postprocess"

	// FormatPostProcess represents an end-user facing format for postProcess options.
	// <option> refers to the "postprocess" option.
	FormatPostProcess = "<option>:<whitespaces><ident>"
)

// ParsePostProcess parses a postprocess option.
func ParsePostProcess(option string) (*Option, error) {
	splitoption := strings.Fields(option)
	if len(splitoption) != 1 {
		return nil, fmt.Errorf("there is a misconfigured %s option: %q.\nIs it in format %s?", CategoryPostProcess, option, FormatPostProcess)
	}

	return &Option{
		Category: CategoryPostProcess,
		Value:    splitoption[0],
	}, nil
}
