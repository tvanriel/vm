package preprocessor

import (
	"strings"

	lex "github.com/bbuck/go-lexer"
	parse "github.com/tvanriel/go-parser"
)

const (
	DefineDirective       = "#define"
	IfDefinedDirective    = "#ifdef"
	IfNotDefinedDirective = "#ifndef"
	EndifDirective        = "#endif"
	IncludeDirective      = "#include"
)

const (
	NilNode parse.NodeType = iota
	TextNode
	IfDefinedDirectiveNode
	IfNotDefinedDirectiveNode
	EndifDirectiveNode
	IncludeDirectiveNode
	DefineDirectiveNode

	StringNode
	ConstantNode
)

func Parse(tokens []*lex.Token) *parse.AST {
	p := &parse.Parser{
		Tokens: tokens,
		Cur:    0,
		AST:    &parse.AST{},
	}
	for p.HasTokens() {
		ParseTextOrDirective(p)
	}
	return p.AST
}

func ParseTextOrDirective(p *parse.Parser) {
	currentToken := p.Current()
	if currentToken == nil {
		return
	}
	if currentToken.Type == TextToken {
		if strings.TrimSpace(currentToken.Value) == "" {
			p.Next() // Not saving this one since it has no content.  Skip.
			return
		}
		p.AddChild(&parse.AST{
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

func ParseInsideIfDef(p *parse.Parser) {
	root := p.AST
	ifNode := &parse.AST{
		Parent:      root,
		ValueType:   IfDefinedDirectiveNode,
		ValueString: IfDefinedDirective,
	}
	p.Next() // read over the if
	ifNode.Children = append(ifNode.Children, &parse.AST{
		ValueType:   ConstantNode,
		ValueString: p.Current().Value,
	})
	p.AddChild(ifNode)
	p.Next() // read over the constant
	p.AST = ifNode
	for p.Current().Type != DirectiveToken && p.Current().Value != EndifDirective {
		ParseTextOrDirective(p)
	}

	p.AddChild(&parse.AST{
		ValueType:   EndifDirectiveNode,
		ValueString: EndifDirective,
	})

	p.AST = root

	p.Next()

}

func ParseInsideIfNotDef(p *parse.Parser) {
	root := p.AST
	ifNode := &parse.AST{
		Parent:      root,
		ValueType:   IfNotDefinedDirectiveNode,
		ValueString: IfNotDefinedDirective,
	}
	p.Next() // read over the if
	ifNode.Children = append(ifNode.Children, &parse.AST{
		ValueType:   ConstantNode,
		ValueString: p.Current().Value,
	})
	p.AddChild(ifNode)
	p.Next() // read over the constant
	p.AST = ifNode
	for p.Current().Type != DirectiveToken && p.Current().Value != EndifDirective {
		ParseTextOrDirective(p)
	}

	p.AddChild(&parse.AST{
		ValueType:   EndifDirectiveNode,
		ValueString: EndifDirective,
	})

	p.AST = root

	p.Next()
}

func ParseDefine(p *parse.Parser) {
	p.Next()
	p.AddChild(&parse.AST{
		ValueType:   DefineDirectiveNode,
		ValueString: DefineDirective,
		Children: []*parse.AST{
			{
				ValueType:   ConstantNode,
				ValueString: p.Current().Value,
			},
		},
	})
	p.Next()
}

func ParseInclude(p *parse.Parser) {
	p.Next()
	p.AddChild(&parse.AST{
		ValueType:   IncludeDirectiveNode,
		ValueString: IncludeDirective,
		Children: []*parse.AST{
			{
				ValueType:   StringNode,
				ValueString: p.Current().Value,
			},
		},
	})
	p.Next()
}
