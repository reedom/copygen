// Package extract uses the `yaegi extract` tool in order to generate the reflect.Value symbols of internal types.
package extract

import "reflect"

// Symbols are extracted from the internal types (compiled at runtime).
var Symbols = make(map[string]map[string]reflect.Value)

//go:generate yaegi extract github.com/reedom/copygen/cli/models
