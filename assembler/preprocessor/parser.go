package preprocessor

import (
	"fmt"
	"strings"

	lex "github.com/bbuck/go-lexer"
)

const (
	DefineDirective       = "#define"
	IfDefinedDirective    = "#ifdef"
	IfNotDefinedDirective = "#ifndef"
	EndifDirective        = "#endif"
	IncludeDirective      = "#include"
)

type NodeType int

const (
	NilNode NodeType = iota
	TextNode
	IfDefinedDirectiveNode
	IfNotDefinedDirectiveNode
	EndifDirectiveNode
	IncludeDirectiveNode
	DefineDirectiveNode

	StringNode
	ConstantNode
)

type AST struct {
	Children    []*AST
	ValueType   NodeType
	ValueString string
	Parent      *AST
}

type Parser struct {
	Expect []lex.TokenType
	Tokens []*lex.Token
	Cur    int
	AST    *AST
}

func (p *Parser) Current() *lex.Token {
	if p.Cur < len(p.Tokens) {
		return p.Tokens[p.Cur]
	}
	return nil
}
func (p *Parser) Next() {
	p.Cur++
}

func (p *Parser) HasTokens() bool {
	return p.Cur < len(p.Tokens)
}

func (p *Parser) AddChild(ast *AST) {
	ast.Parent = ast
	p.AST.Children = append(p.AST.Children, ast)
}

type ParseFunc func(*Parser) ParseFunc

func Parse(tokens []*lex.Token) *AST {
	p := &Parser{
		Tokens: tokens,
		Cur:    0,
		AST:    &AST{},
	}
	for p.HasTokens() {
		ParseTextOrDirective(p)
	}
	return p.AST
}

func ParseTextOrDirective(p *Parser) {
	currentToken := p.Current()
	if currentToken == nil {
		return
	}
	if currentToken.Type == TextToken {
		if strings.TrimSpace(currentToken.Value) == "" {
			p.Next() // Not saving this one since it has no content.  Skip.
			return
		}
		p.AddChild(&AST{
			ValueType:   TextNode,
			ValueString: currentToken.Value,
		})
		p.Next()
		return
	}
	if currentToken.Value == DefineDirective {
		ParseDefine(p)
		return
	}
	if currentToken.Value == IfDefinedDirective {
		ParseInsideIfDef(p)
		return
	}
	if currentToken.Value == IfNotDefinedDirective {
		ParseInsideIfNotDef(p)
		return
	}

	if currentToken.Value == IncludeDirective {
		ParseInclude(p)
		return
	}
}

func ParseInsideIfDef(p *Parser) {
	root := p.AST
	ifNode := &AST{
		Parent:      root,
		ValueType:   IfDefinedDirectiveNode,
		ValueString: IfDefinedDirective,
	}
	p.Next() // read over the if
	ifNode.Children = append(ifNode.Children, &AST{
		ValueType:   ConstantNode,
		ValueString: p.Current().Value,
	})
	p.AddChild(ifNode)
	p.Next() // read over the constant
	p.AST = ifNode
	for p.Current().Type != DirectiveToken && p.Current().Value != EndifDirective {
		ParseTextOrDirective(p)
	}

	p.AddChild(&AST{
		ValueType:   EndifDirectiveNode,
		ValueString: EndifDirective,
	})

	p.AST = root

	p.Next()

}

func ParseInsideIfNotDef(p *Parser) {
	root := p.AST
	ifNode := &AST{
		Parent:      root,
		ValueType:   IfNotDefinedDirectiveNode,
		ValueString: IfNotDefinedDirective,
	}
	p.Next() // read over the if
	ifNode.Children = append(ifNode.Children, &AST{
		ValueType:   ConstantNode,
		ValueString: p.Current().Value,
	})
	p.AddChild(ifNode)
	p.Next() // read over the constant
	p.AST = ifNode
	for p.Current().Type != DirectiveToken && p.Current().Value != EndifDirective {
		fmt.Printf("%d: %v, %v", p.Cur, p.Current().Type != DirectiveToken, p.Current().Value != EndifDirective)
		ParseTextOrDirective(p)
	}

	p.AddChild(&AST{
		ValueType:   EndifDirectiveNode,
		ValueString: EndifDirective,
	})

	p.AST = root

	p.Next()
}

func ParseDefine(p *Parser) {
	p.Next()
	p.AddChild(&AST{
		ValueType:   DefineDirectiveNode,
		ValueString: DefineDirective,
		Children: []*AST{
			{
				ValueType:   ConstantNode,
				ValueString: p.Current().Value,
			},
		},
	})
	p.Next()
}

func ParseInclude(p *Parser) {
	p.Next()
	p.AddChild(&AST{
		ValueType:   IncludeDirectiveNode,
		ValueString: IncludeDirective,
		Children: []*AST{
			{
				ValueType:   StringNode,
				ValueString: p.Current().Value,
			},
		},
	})
	p.Next()
}
