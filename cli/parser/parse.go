// Package parser parses a setup file's functions, types, and fields using an Abstract Syntax Tree.
package parser

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"path/filepath"

	"github.com/switchupcb/copygen/cli/models"
)

// Parser represents a parser that parses Abstract Syntax Tree data into models.
type Parser struct {
	// The fileset of the parser.
	Fileset *token.FileSet

	// The setup file as an Abstract Syntax Tree.
	SetupFile *ast.File

	// The option-comments parsed in the OptionMap.
	Comments []*ast.Comment

	// The ast.Node of the `type Copygen Interface`.
	Copygen *ast.InterfaceType

	// The imports discovered in the set up file (map[packagevar]importpath).
	// In the context of the parser, packagevar refers to the the variable used
	// to reference the package (alias) rather the package's actual name.
	Imports map[string]string

	// The parser options contain options located in the entire setup file.
	Options OptionMap

	// The setup filepath.
	Setpath string
}

// Parse parses a generator's setup file.
func Parse(gen *models.Generator) error {
	// determine the actual filepath of the setup.go file.
	absfilepath, err := filepath.Abs(filepath.Join(filepath.Dir(gen.Loadpath), gen.Setpath))
	if err != nil {
		return err
	}

	p := Parser{Setpath: absfilepath}
	p.Fileset = token.NewFileSet()
	p.SetupFile, err = parser.ParseFile(p.Fileset, absfilepath, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("An error occurred parsing the specified .go setup file: %v.\n%v", gen.Setpath, err)
	}

	// Traverse the Abstract Syntax Tree.
	p.Options = make(OptionMap)
	p.parseImports()
	err = p.Traverse(gen)
	if err != nil {
		return err
	}

	// write the Keep
	buf := new(bytes.Buffer)
	buf.WriteString("// Code generated by github.com/switchupcb/copygen\n// DO NOT EDIT.\n\n")
	err = printer.Fprint(buf, p.Fileset, p.SetupFile)
	if err != nil {
		return fmt.Errorf("An error occurred writing the code that will be kept after generation.\n%v", err)
	}
	gen.Keep = buf.Bytes()
	return nil
}

// parseImports parses the AST for imports in the setup file.
func (p *Parser) parseImports() {
	if p.Imports == nil {
		p.Imports = make(map[string]string) // map[packagevar]importpath
	}

	for _, imprt := range p.SetupFile.Imports {
		if imprt.Name != nil { // aliased package (i.e c "strconv")
			p.Imports[imprt.Name.Name] = imprt.Path.Value
		} else {
			base := filepath.Base(imprt.Path.Value)
			// [:removes the last `"` from the package name]
			p.Imports[base[:len(base)-1]] = imprt.Path.Value
		}
	}
}
