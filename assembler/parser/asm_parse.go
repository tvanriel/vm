package parser

import (
	lex "github.com/bbuck/go-lexer"
	parse "github.com/tvanriel/go-parser"
	"github.com/tvanriel/vm/assembler/lexer"
)

const (
	NilNode parse.NodeType = iota

	ConstantNode
	ExpressionNode
	WordNode

	HexadecimalNumberNode
	OctalNumberNode
	BinaryNumberNode
	DecimalNumberNode

	PlusNode
	MinusNode
	AsteriskNode
	SlashNode
	GreaterThanNode
	LessThanNode

	LabelNode
	InstructionNode
	ArgumentNode
)

func Parse(tokens []*lex.Token) (*parse.AST, error) {
	p := &parse.Parser{
		Tokens: tokens,
		Cur:    0,
		AST:    &parse.AST{},
	}
	for p.Continue() {
		ParseLabelOrConstant(p)
	}
	if p.Err() != nil {
		return nil, p.Err()
	}
	return p.AST, nil
}

func ParseLabelOrConstant(p *parse.Parser) {
	if p.Current().Type == lexer.ConstantToken {
		ParseConstant(p)
		return
	}

	if p.Current().Type == lexer.LabelToken {
		ParseLabel(p)
		return
	}
}

func ParseConstant(p *parse.Parser) {
	root := p.AST
	constant := &parse.AST{
		ValueType:   ConstantNode,
		ValueString: p.Current().Value,
	}
	p.Next() // Read over the constant
	if p.Current().Type != lexer.EqualsSign {
		p.SetError("parse error: expected =, got: " + p.Current().Value)
	}
	p.Next() // Read over the =
	p.AddChild(constant)
	p.AST = constant

	ParseExpression(p)

	p.AST = root
}

func ParseExpression(p *parse.Parser) {

	if p.Current().Type == lexer.WordToken {
		p.AddChild(&parse.AST{
			ValueType:   WordNode,
			ValueString: p.Current().Value,
		})
		p.Next()
		return
	}

	if IsNumber(p.Current().Type) {
		p.AddChild(&parse.AST{
			ValueType:   tokennumtonodenum[p.Current().Type],
			ValueString: p.Current().Value,
		})
		p.Next()
		return
	}

	if p.Current().Type == lexer.ConstantToken {
		p.AddChild(&parse.AST{
			ValueType:   ConstantNode,
			ValueString: p.Current().Value,
		})
		p.Next()
		return
	}
	if p.Current().Type == lexer.ParenthesisOpen {
		root := p.AST
		expression := &parse.AST{
			ValueType:   ExpressionNode,
			ValueString: "Expr:",
		}
		p.AddChild(expression)
		p.AST = expression
		p.Next() // Read over (
		ParseExpression(p)
		ParseOperator(p)
		ParseExpression(p)
		if p.Current().Type != lexer.ParenthesisClose {
			p.SetError("parse error: expected ), got: " + p.Current().Value)
		}
		p.Next() // Read over )
		p.AST = root
		return
	}
	p.SetError("parse error: expected number constant (, got: " + p.Current().Value)
}

func ParseOperator(p *parse.Parser) {
	if p.Current().Type == lexer.PlusSign {
		p.AddChild(&parse.AST{
			ValueString: "+",
			ValueType:   PlusNode,
		})
		p.Next()
		return
	}
	if p.Current().Type == lexer.MinusSign {
		p.AddChild(&parse.AST{
			ValueString: "-",
			ValueType:   MinusNode,
		})
		p.Next()
		return
	}
	if p.Current().Type == lexer.AsteriskSign {
		p.AddChild(&parse.AST{
			ValueString: "*",
			ValueType:   AsteriskNode,
		})
		p.Next()
		return
	}
	if p.Current().Type == lexer.SlashSign {
		p.AddChild(&parse.AST{
			ValueString: "/",
			ValueType:   SlashNode,
		})
		p.Next()
		return
	}

	if p.Current().Type == lexer.OpLessThan {
		p.AddChild(&parse.AST{
			ValueString: "<",
			ValueType:   LessThanNode,
		})
		p.Next()
		return
	}

	if p.Current().Type == lexer.OpGreaterThan {
		p.AddChild(&parse.AST{
			ValueString: ">",
			ValueType:   GreaterThanNode,
		})
		p.Next()
		return
	}

	p.SetError("parse error: expected + - / * < >, got: " + p.Current().Value)
}

func ParseLabel(p *parse.Parser) {
	root := p.AST
	label := &parse.AST{
		ValueType:   LabelNode,
		ValueString: p.Current().Value,
	}
	p.Next() // Read over the label
	if p.Current().Type != lexer.ColonSign {
		p.SetError("parse error: expected :, got " + p.Current().Value)
		return
	}
	p.Next() // Read over :
	SkipComments(p)
	p.AddChild(label)
	p.AST = label

	for p.Current().Type == lexer.InstructionToken {
		ParseInstruction(p)
		if p.Current() == nil {
			break
		}
	}
	p.AST = root
}

func SkipComments(p *parse.Parser) {
	for p.Current().Type == lexer.CommentToken {
		p.Next()
	}
}

func ParseInstruction(p *parse.Parser) {
	root := p.AST

	instruction := &parse.AST{
		ValueType:   InstructionNode,
		ValueString: p.Current().Value,
	}
	p.Next() // Read over the instruction

	p.AddChild(instruction)
	p.AST = instruction

	if p.Current() == nil {
		p.AST = root
		return
	}
	for p.Current().Type == lexer.ConstantToken || p.Current().Type == lexer.ParenthesisOpen || IsNumber(p.Current().Type) || p.Current().Type == lexer.WordToken {
		ParseExpression(p)
		if p.Current() == nil {
			break
		}
		if p.Current().Type == lexer.CommaSign {
			p.Next()
		}
		SkipComments(p)
	}

	p.AST = root
}
