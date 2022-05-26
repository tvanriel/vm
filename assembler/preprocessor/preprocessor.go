package preprocessor

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Scan a file for directives
func LexAndParse(path string) (*AST, error) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	tokens := Tokenize(string(fileContent))
	ast := Parse(tokens)

	err = processDependencies(ast, nil)
	if err != nil {
		return nil, err
	}

	return ast, nil
}

func processDependencies(ast *AST, n *int) error {
	var err error
	if ast.ValueType == IncludeDirectiveNode {
		child, err := LexAndParse(ast.Children[0].ValueString)
		if err != nil {
			return err
		}
		ast.Parent.Children = child.Children
		ast.Children = child.Children
		ast.ValueString = child.ValueString
		ast.ValueType = child.ValueType

	}

	for i, childAst := range ast.Children {
		err = processDependencies(childAst, &i)
		if err != nil {
			return err
		}
	}
	return nil
}

func Process(path string) (string, error) {
	ast, err := LexAndParse(path)
	if err != nil {
		return "", err
	}

	sb := &strings.Builder{}
	stringifyLeaf(ast, &[]string{}, sb)
	return sb.String(), err

}

func stringifyLeaf(ast *AST, constants *[]string, sb *strings.Builder) {
	if ast.ValueType == NilNode {
		for _, child := range ast.Children {
			stringifyLeaf(child, constants, sb)
		}
		return
	}
	if ast.ValueType == TextNode {
		sb.WriteString(ast.ValueString)
		return
	}
	if ast.ValueType == DefineDirectiveNode {
		*constants = append(*constants, ast.Children[0].ValueString)
		return
	}
	if ast.ValueType == IfDefinedDirectiveNode {
		if constIsDefined(constants, ast.Children[0].ValueString) {
			for _, child := range ast.Children[1:] {
				stringifyLeaf(child, constants, sb)
			}
			return
		}
	}

	if ast.ValueType == IfNotDefinedDirectiveNode {
		if !constIsDefined(constants, ast.Children[0].ValueString) {
			for _, child := range ast.Children[1:] {
				stringifyLeaf(child, constants, sb)
			}
			return
		}
	}
	fmt.Println("How could this happened?!")
}

func constIsDefined(constants *[]string, name string) bool {
	for _, v := range *constants {
		if v == name {
			return true
		}
	}
	return false
}