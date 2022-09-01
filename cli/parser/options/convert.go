package options

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"

	"github.com/reedom/copygen/cli/models"
)

const (
	CategoryConvert = "convert"

	// FormatConvert represents an end-user facing format for convert options.
	// <option> refers to the "convert" option.
	FormatConvert = "<option>:<whitespaces><regex><whitespaces><regex>"
)

var errorMarker = regexp.MustCompile("error")

// ParseConvert parses a convert option.
func ParseConvert(option, value string, funcType *ast.FuncType) (*Option, error) {
	splitoption := strings.Fields(option)
	if len(splitoption) == 0 {
		return nil, fmt.Errorf("there is an unspecified %s option at an unknown line", CategoryConvert)
	} else if len(splitoption) != 2 {
		return nil, fmt.Errorf("there is a misconfigured %s option: %q.\nIs it in format %s?", CategoryConvert, option, FormatConvert)
	}

	funcRe, err := regexp.Compile("^" + splitoption[0] + "$")
	if err != nil {
		return nil, fmt.Errorf("an error occurred compiling the regex for the first field in the %s option: %q\n%w", CategoryConvert, option, err)
	}

	fieldRe, err := regexp.Compile("^" + splitoption[1] + "$")
	if err != nil {
		return nil, fmt.Errorf("an error occurred compiling the regex for the second field in the %s option: %q\n%w", CategoryConvert, option, err)
	}

	var errRe *regexp.Regexp
	if returnsError(funcType) {
		errRe = errorMarker
	}

	return &Option{
		Category: CategoryConvert,
		Regex:    map[int]*regexp.Regexp{0: funcRe, 1: fieldRe, 2: errRe},
		Value:    value,
	}, nil
}

// SetConvert sets a field's convert option.
func SetConvert(field *models.Field, option Option) {
	// A convert option can only be set to a field once.
	if !field.Options.Convert.IsEmpty() {
		return
	}

	if option.Regex[1] != nil && option.Regex[1].MatchString(field.FullNameWithoutPointer("")) {
		if value, ok := option.Value.(string); ok {
			field.Options.Convert.Ident = value
			field.Options.Convert.Error = option.Regex[2] == errorMarker
		}
	}
}

func returnsError(funcType *ast.FuncType) bool {
	results := funcType.Results
	if 0 < results.NumFields() {
		if ident, ok := results.List[results.NumFields()-1].Type.(*ast.Ident); ok {
			return ident.Name == "error"
		}
	}
	return false
}
