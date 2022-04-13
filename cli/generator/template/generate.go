// DO NOT CHANGE PACKAGE

// Package template provides a template used by copygen to generate custom code.
package template

import (
	"github.com/switchupcb/copygen/cli/models"
)

// Generate generates code.
// GENERATOR FUNCTION.
// EDITABLE.
// DO NOT REMOVE.
func Generate(gen *models.Generator) (string, error) {
	content := string(gen.Keep) + "\n"

	for i := range gen.Functions {
		content += Function(&gen.Functions[i]) + "\n"
	}

	return content, nil
}

// Function provides generated code for a function.
func Function(function *models.Function) string {
	// comment
	fn := generateComment(function) + "\n"

	// signature
	fn += generateSignature(function) + "\n"

	// body
	fn += generateBody(function)

	// return
	fn += generateReturn(function)

	// end of function
	if fn[len(fn)-1:] != "\n" {
		fn += "\n"
	}
	fn += "}"

	return fn
}

// generateComment generates a function comment.
func generateComment(function *models.Function) string {
	var toComment string

	for _, toType := range function.To {
		toComment += toType.Field.FullNameWithoutContainer("") + ", "
	}

	if toComment != "" {
		// remove last ", "
		toComment = toComment[:len(toComment)-2]
	}

	var fromComment string
	for _, fromType := range function.From {
		fromComment += fromType.Field.FullNameWithoutContainer("") + ", "
	}

	if fromComment != "" {
		// remove last ", "
		fromComment = fromComment[:len(fromComment)-2]
	}

	return "// " + function.Name + " copies a " + fromComment + " to a " + toComment + "."
}

// generateSignature generates a function's signature.
func generateSignature(function *models.Function) string {
	return "func " + function.Name + "(" + generateParameters(function) + ") {"
}

// generateParameters generates the parameters of a function.
func generateParameters(function *models.Function) string {
	var parameters string

	// Generate To-Type parameters
	for _, toType := range function.To {
		parameters += toType.Field.VariableName + " "
		parameters += toType.ParameterName() + ", "
	}

	// Generate From-Type parameters
	for _, fromType := range function.From {
		parameters += fromType.Field.VariableName + " "
		parameters += fromType.ParameterName() + ", "
	}

	if parameters == "" {
		return parameters
	}

	// remove last ", "
	return parameters[:len(parameters)-2]
}

// generateBody generates the body of a function.
func generateBody(function *models.Function) string {
	var body string

	// Assign fields to ToType(s).
	for _, toType := range function.To {
		body += "// " + toType.Field.FullNameWithoutContainer("") + " fields\n"

		for _, toField := range toType.Field.Fields {
			body += toField.FullVariableName("") + " = "
			fromField := toField.From

			switch {
			case fromField.Options.Convert != "":
				body += fromField.Options.Convert + "(" + fromField.FullVariableName("") + ")\n"
			case toField.Definition != fromField.Definition:
				// match alias types to respective basic types.
				if toField.Package != "" {
					body += toField.Package + "."
				}
				body += toField.Definition + "(" + fromField.FullVariableName("") + ")" + "\n"
			default:
				body += fromField.FullVariableName("") + "\n"
			}
		}
	}

	return body
}

// generateReturn generates a return statement for the function.
func generateReturn(function *models.Function) string {
	return ""
}
