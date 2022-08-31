package parser

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"

	"github.com/reedom/copygen/cli/parser/options"
)

const convertOptionSplitAmount = 3

// Keep removes ast.Nodes from an ast.File that will be kept in a generated output file.
func (p *Parser) Keep(astFile *ast.File) error {
	var trash []*ast.Comment

	for i := len(astFile.Decls) - 1; i > -1; i-- {
		switch declaration := astFile.Decls[i].(type) {
		case *ast.GenDecl:

			// keep all declaration objects in the setup file except for the `type Copygen interface`.
			if _, ok := assertCopygenInterface(declaration); ok {

				// remove from the `type Copygen interface` (from the slice).
				astFile.Decls[i] = astFile.Decls[len(astFile.Decls)-1]
				astFile.Decls = astFile.Decls[:len(astFile.Decls)-1]

				// remove the `type Copygen interface` function ast.Comments.
				comments := getNodeComments(declaration)
				if err := p.assignFieldOption(comments); err != nil {
					return fmt.Errorf("%w", err)
				}
				trash = append(trash, comments...)
			}

		case *ast.FuncDecl:
			comments, err := p.assignConvertOptions(declaration)
			if err != nil {
				return fmt.Errorf("%w", err)
			}

			// remove convert option ast.Comments.
			trash = append(trash, comments...)
		}
	}

	// also trash comments that match against the specific patterns.
	for _, fileCommentGroup := range astFile.Comments {
		for _, comment := range fileCommentGroup.List {
			if p.shouldSkip(comment.Text) {
				trash = append(trash, comment)
			}
		}
	}

	// Remove ast.Comments that will be parsed into options from the ast.File.
	astRemoveComments(astFile, trash)

	return nil
}

// assignFieldOption parses a list of ast.Comments into options
// and places them in a map[text]Option.
func (p *Parser) assignFieldOption(comments []*ast.Comment) error {
	if p.Options.CommentOptionMap == nil {
		p.Options.CommentOptionMap = make(map[string]*options.Option, len(comments))
	}

	for _, comment := range comments {
		text := comment.Text

		// do NOT parse comments that have already been parsed.
		if p.Options.CommentOptionMap[text] != nil {
			continue
		}

		splitcomments := strings.Fields(text[2:])
		if len(splitcomments) >= 1 {

			category := splitcomments[0]
			if category == options.CategoryConvert {
				continue
			}

			optiontext := strings.Join(splitcomments[1:], " ")
			option, err := options.NewFieldOption(category, optiontext)
			if err != nil {
				return fmt.Errorf("%w", err)
			}

			p.Options.CommentOptionMap[text] = option
		}
	}

	return nil
}

// assignConvertOptions initializes convert options.
// Used in the context of functions other than the type Copygen interface.
func (p *Parser) assignConvertOptions(x *ast.FuncDecl) ([]*ast.Comment, error) {
	var (
		convertComments []*ast.Comment
		assignErr       error
	)

	ast.Inspect(x, func(node ast.Node) bool {
		commentGroup, ok := node.(*ast.CommentGroup)
		if !ok {
			return true
		}

		for _, comment := range commentGroup.List {
			text := comment.Text
			splitcomments := strings.Fields(text[2:])

			// determine if the comment is a convert option.
			if len(splitcomments) == convertOptionSplitAmount {
				category := splitcomments[0]
				value := strings.Join(splitcomments[1:], " ")
				if category == options.CategoryConvert {
					option, err := options.ParseConvert(value, x.Name.Name)
					if err != nil {
						assignErr = err
						return false
					}

					p.Options.ConvertOptions = append(p.Options.ConvertOptions, option)
					convertComments = append(convertComments, comment)
				}
			}
		}

		return true
	})

	return convertComments, assignErr
}

var reSkip = regexp.MustCompile(`^//go:(generate|build copygen)\b`)

// shouldSkip determines whether the generator should strip the comment from the result code.
func (p *Parser) shouldSkip(comment string) bool {
	return reSkip.MatchString(comment)
}
