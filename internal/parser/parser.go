package parser

import (
	"fmt"
	"os"

	. "github.com/LinYUAN-code/MyBuilder/internal/ast"
	. "github.com/LinYUAN-code/MyBuilder/internal/lexer"
)

/*
	首先找到具体的JS文法
	https://hepunx.rl.ac.uk/~adye/jsspec11/llr.htm
	补充自己还想要的语法 , 然后计算first 和 follow 集合 当然刚开始只凭感觉写也是没有问题的
	然后直接递归下降分下法一直写下来就好了
*/

func Assert(condition bool) {
	fmt.Errorf("parser 发生语法错误")
}



func NewParser(path string) (*Parser,error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil,err
	}
	lexer := NewLexer(string(data))
	return &Parser{
		Lexer: lexer,
	},nil
}

func (parse *Parser) ShowAllTocken() {
	parse.Lexer.Next()
	for parse.Lexer.Tocken != LENDOFFILE {
		// switch parse.Lexer.Tocken {
		// case lexer.LEqual:
		// 	println("=")
		// case lexer.LIdentifier:
		// 	println("identifier: ",parse.Lexer.Identifier)
		// case lexer.LInteger:
		// 	println("Integer: ",parse.Lexer.Integer)
		// case lexer.LSemicolon:
		// 	println(";")
		// case lexer.LLet:
		// 	println("let")
		// }
		println(TToString[parse.Lexer.Tocken])
		parse.Lexer.Next()
	}
}


func (parser *Parser) parseArgs() ([]Expr) {
	result := make([]Expr,0)
	for parser.Lexer.Tocken == LIdentifier {
		result = append(result, NewEIdentifier(parser.Lexer.Value))
		parser.Lexer.Next()
		if parser.Lexer.Tocken != LComma {
			break
		}
		parser.Lexer.Next()
	}
	return result
}

// 采用递归下降分析法
func (parser *Parser) ParseStmts() (Stmts) {
	program := Stmts{}
	program.Stmt = make([]Stmt,0)
	for {
		program.Stmt = append(program.Stmt, parser.ParseElement())
	}
}

func (parser *Parser) ParseElement() (Stmt) {
	switch parser.Lexer.Tocken {
	case LFuntion:
		parser.Lexer.Next()
		Assert(parser.Lexer.Tocken==LIdentifier)
		name := NewEIdentifier(parser.Lexer.Value)
		parser.Lexer.Next()
		args := parser.parseArgs()
		parser.Lexer.Next()
		body := parser.ParseCompoundStatement()
		return NewSFunction(name,args,body)
	default:
		return parser.ParseStatement()
	}
}

func (parser *Parser) ParseCompoundStatement() Stmt {
	Assert(parser.Lexer.Tocken==LOpenBraket)
	body := parser.ParseStatements()
	Assert(parser.Lexer.Tocken==LCloseBraket)
	return NewCompoundStatement(body)
}



func (parser *Parser) ParseStatements() []Stmt {
	result := make([]Stmt,0)
	Assert(parser.Lexer.Tocken==LOpenBraket)
	parser.Lexer.Next()
	for parser.Lexer.Tocken != LCloseBraket {
		result = append(result, parser.ParseStatement())
		parser.Lexer.Next()
	}
	return result
}

func (parser *Parser) ParseStatement() Stmt {
	for {
		switch parser.Lexer.Tocken {
		case LSemicolon:
			parser.Lexer.Next()
			continue
		case LIf: 	
			parser.Lexer.Next()
			expr := parser.ParseCondition()
			pass := parser.ParseStatement()
			var fail *Stmt = nil
			if parser.Lexer.Tocken==LElse {
				value := parser.ParseStatement()
				fail = &value
			}
			return NewSIf(expr,pass,fail)
		case LWhile:
			parser.Lexer.Next()
			expr := parser.ParseCondition()
			stmt := parser.ParseStatement()
			return NewSWhile(expr,stmt)
		case LFor: //还不支持for of
			parser.Lexer.Next()
			Assert(parser.Lexer.Tocken==LOpenParen)
			var fir,sec,thir Expr
			var stmt Stmt
			if parser.Lexer.Tocken == LComma {
				fir = NewEEmpty()
			} else {
				fir = parser.ParseVariablesOrExpression()
			}
			if parser.Lexer.Tocken == LIn {
				parser.Lexer.Next()
				sec = parser.ParseExpression()
				Assert(parser.Lexer.Tocken == LCloseParen)
				parser.Lexer.Next()
				stmt = parser.ParseStatement()
				return NewSForIn(fir,sec,stmt)
			} else {
				parser.Lexer.Next()
				if parser.Lexer.Tocken == LComma {
					sec = NewEEmpty()
				} else {
					sec = parser.ParseExpression()
				}
				parser.Lexer.Next()
				if parser.Lexer.Tocken == LCloseParen {
					thir = NewEEmpty()
				} else {
					thir = parser.ParseExpression()
				}
				parser.Lexer.Next()
				stmt = parser.ParseStatement()
				return NewSFor(fir,sec,thir,stmt)
			}
		case LBreak:
			return NewSBreak()
		case LContinue:
			return NewSContinue()
		case LWith:
			parser.Lexer.Next()
			Assert(parser.Lexer.Tocken == LOpenParen)
			parser.Lexer.Next()
			fir := parser.ParseExpression()
			Assert(parser.Lexer.Tocken == LCloseParen)
			parser.Lexer.Next()
			stmt := parser.ParseStatement()
			return NewSWith(fir,stmt)
		case LReturn:
			parser.Lexer.Next()
			var fir Expr
			if parser.Lexer.Tocken == LSemicolon {
				fir = NewEEmpty()
			} else {
				fir = parser.ParseExpression()
			}
			return NewSReturn(fir)
		case LOpenBrace:
			return parser.ParseCompoundStatement()
		default:
			fir := parser.ParseVariablesOrExpression()
			return NewSVariablesOrExpression(fir)
		}

	}

}

func (parser *Parser) ParseVariablesOrExpression() Expr {
	var opt Declear
	switch parser.Lexer.Tocken {
	case LLet:
		opt = Let
	case LVar:
		opt = Var
	case LConst:
		opt = Const
	default:
		return parser.ParseExpression()
	}
	variables := parser.ParseVariables()
	return NewEDeclearVariable(opt,variables)
}

func (parser *Parser) ParseVariables() Expr {	
	results := make([]Expr,0)
	for {
		Assert(parser.Lexer.Tocken == LIdentifier)
		results = append(results, NewEIdentifier(parser.Lexer.Value))
		if parser.Lexer.Tocken != LComma {
			break
		}
		parser.Lexer.Next()
	}	
	assignmentExpression := parser.ParseAssignmentExpression()
	return NewEVariables(results,assignmentExpression)
}

func (parser *Parser) ParseAssignmentExpression() Expr {
	fir := parser.ParseConditionalExpression()
	// AssignmentOperator 
	switch parser.Lexer.Tocken {
	case LMinusEqual:
		parser.Lexer.Next()
		sec := parser.ParseAssignmentExpression()
		return NewEDualCal(fir,MinusEqual,sec)
	case LPlusEqual:
		parser.Lexer.Next()
		sec := parser.ParseAssignmentExpression()
		return NewEDualCal(fir,PlusEqual,sec)
	}
	return fir
}

func (parser *Parser) ParseConditionalExpression() Expr {
	fir := parser.ParseOrExpression()
	if parser.Lexer.Tocken != LQuestion {
		return fir
	}
	sec := parser.ParseAssignmentExpression()
	thir := parser.ParseAssignmentExpression()
	return NewEConditionalExpression(fir,sec,thir)
}

func (parser *Parser) ParseOrExpression() Expr {
	fir := parser.ParseAndExpression()
	if parser.Lexer.Tocken == LBarBar {
		sec := parser.ParseOrExpression()
		return NewEDualCal(fir,BarBar,sec)
	}
	return fir
}

func (parser *Parser) ParseAndExpression() Expr {
	fir := parser.ParseBitwiseOrExpression()
	if parser.Lexer.Tocken == LAndAnd {
		sec := parser.ParseAndExpression()
		return NewEDualCal(fir,AndAnd,sec)
	}
	return fir
}

func (parser *Parser) ParseBitwiseOrExpression() Expr {
	fir := parser.ParseBitwiseXorExpression()
	if parser.Lexer.Tocken == LBar {
		sec := parser.ParseBitwiseOrExpression()
		return NewEDualCal(fir,Bar,sec)
	}
	return fir
}

func (parser *Parser) ParseBitwiseXorExpression() Expr {
	fir := parser.ParseBitwiseAndExpression()
	if parser.Lexer.Tocken == LCaret {
		sec := parser.ParseBitwiseAndExpression()
		return NewEDualCal(fir,Caret,sec)
	}
	return fir
}

func (parser *Parser) ParseBitwiseAndExpression() Expr {
	fir := parser.ParseEqualityExpression()
	if parser.Lexer.Tocken == LAnd {
		sec := parser.ParseBitwiseAndExpression()
		return NewEDualCal(fir,And,sec)
	}
	return fir
}


// == === != !==
func (parser *Parser) ParseEqualityExpression() Expr {
	fir := parser.ParseRelationalExpression()
	switch parser.Lexer.Tocken {
	case LEqualEqual:
		sec := parser.ParseEqualityExpression()
		return NewEDualCal(fir,EqualEuqal,sec)
	case LEqualEqualEqual:
		sec := parser.ParseEqualityExpression()
		return NewEDualCal(fir,EqualEuqalEuqal,sec)
	case LExclamationEqual:
		sec := parser.ParseEqualityExpression()
		return NewEDualCal(fir,ExclamationEqual,sec)
	case LExclamationEqualEqual:
		sec := parser.ParseEqualityExpression()
		return NewEDualCal(fir,ExclamationEqualEqual,sec)
	}
	return fir
}

func (parser *Parser) ParseRelationalExpression() Expr {
	fir := parser.ParseShiftExpression()
	switch parser.Lexer.Tocken {
	case LLess:
		sec := parser.ParseRelationalExpression()
		return NewEDualCal(fir,Less,sec)
	case LLessEqual:
		sec := parser.ParseRelationalExpression()
		return NewEDualCal(fir,LessEqual,sec)
	case LGreater:
		sec := parser.ParseRelationalExpression()
		return NewEDualCal(fir,Greater,sec)
	case LGreaterEqual:
		sec := parser.ParseRelationalExpression()
		return NewEDualCal(fir,GreaterEqual,sec)
	}
	return fir
}

func (parser *Parser) ParseShiftExpression() Expr {
	fir := parser.ParseAdditiveExpression()
	switch parser.Lexer.Tocken {
	case LLshift:
		sec := parser.ParseShiftExpression()
		return NewEDualCal(fir,LShift,sec)
	case LRshift:
		sec := parser.ParseShiftExpression()
		return NewEDualCal(fir,RShift,sec)
	}
	return fir
}

func (parser *Parser) ParseAdditiveExpression() Expr {
	fir := parser.ParseMultiplicativeExpression()
	switch parser.Lexer.Tocken {
	case LPlus:
		sec := parser.ParseAdditiveExpression()
		return NewEDualCal(fir,Plus,sec)
	case LMinus:
		sec := parser.ParseAdditiveExpression()
		return NewEDualCal(fir,Minus,sec)
	}
	return fir
}

func (parser *Parser) ParseMultiplicativeExpression() Expr {
	fir := parser.ParseUnaryExpression()
	if parser.Lexer.Tocken == LMulti {
		sec := parser.ParseMultiplicativeExpression()
		return NewEDualCal(fir,Multi,sec)
	}
	return fir
}

func (parser *Parser) ParseUnaryExpression() Expr {
	switch parser.Lexer.Tocken {
	case LNew:
		parser.Lexer.Next()
		expr := parser.ParseConstructor()
		return NewEUnary(false,New,expr)
	case LDelete:
		parser.Lexer.Next()
		expr := parser.ParseMemberExpression()
		return NewEUnary(false,Delete,expr)
	case LPlusPlus:
		parser.Lexer.Next()
		expr := parser.ParseMemberExpression()
		return NewEUnary(false,NPlusPlus,expr)
	case LMinusMinus:
		parser.Lexer.Next()
		expr := parser.ParseMemberExpression()
		return NewEUnary(false,NMinusMinus,expr)
	case LMinus:
		parser.Lexer.Next()
		expr := parser.ParseUnaryExpression()
		unaryExpr := expr.Data.(EUnary)
		unaryExpr.Neg = true
		return Expr{
			Data: unaryExpr,
		}
	default:
		expr := parser.ParseMemberExpression()
		switch parser.Lexer.Tocken {
		case LPlusPlus:
			return NewEUnary(false,RPlusPlus,expr)
		case LMinusMinus:
			return NewEUnary(false,RMinusMinus,expr)
		}
		return expr
	}	
}

func (parser *Parser) ParseConstructor() Expr {
	if parser.Lexer.Tocken == LThis {
		parser.Lexer.Next()
		Assert(parser.Lexer.Tocken == LDot)
		expr := parser.ParseConstructorCall()
		return NewEConstructor(expr,true)
	}
	expr := parser.ParseConstructorCall()
	return NewEConstructor(expr,false)
}

func (parser *Parser) ParseConstructorCall() Expr {
	Assert(parser.Lexer.Tocken == LIdentifier)
	identifier := NewEIdentifier(parser.Lexer.Value)
	parser.Lexer.Next()
	switch parser.Lexer.Tocken {
	case LOpenParen:
		se := parser.ParseArgumentList()
		return NewEConstructorCall(identifier,Paren,se)
	case LDot:
		se := parser.ParseConstructorCall()
		return NewEConstructorCall(identifier,Dot,se)
	}
	return identifier
}

func (parser *Parser) ParseArguments() []Expr {
	results := make([]Expr,0)
	for parser.Lexer.Tocken!=LCloseParen {
		results = append(results, parser.ParseAssignmentExpression())
		if parser.Lexer.Tocken != LComma {
			break
		}
	}
	return results
}

func (parser *Parser) ParseCondition() Expr {
	Assert(parser.Lexer.Tocken==LOpenParen)
	expr := parser.ParseExpression()
	Assert(parser.Lexer.Tocken==LCloseParen)
	parser.Lexer.Next()
	return expr
}


// 不检测空
func (parser *Parser) ParseExpression() Expr {
	results := make([]Expr,0)
	for {
		results = append(results, parser.ParseAssignmentExpression())
		if parser.Lexer.Tocken != LComma {
			break
		}
		parser.Lexer.Next()
	}
	return NewEExpression(results)
}


func (parser *Parser) ParseMemberExpression() (Expr) {
	fir := parser.parsePrimaryExpression()
	var opt *Opt = nil
	var se  *Expr
	switch parser.Lexer.Tocken {
	case LDot:
		value1 := Dot
		opt = &value1
		sec := parser.ParseMemberExpression()
		return NewEMember(fir,opt,&sec)
	case LOpenBraket:
		value := Braket
		opt = &value
		sec := parser.ParseExpression()
		return NewEMember(fir,opt,&sec)
	case LOpenParen:
		value := Paren
		opt = &value
		sec := parser.ParseArgumentList()
		parser.Lexer.Next()
		return NewEMember(fir,opt,&sec)
	}
	return NewEMember(fir,opt,se)
}

func (parser *Parser) ParseArgumentList() Expr {
	results := make([]Expr,0)
	parser.Lexer.Next()
	for parser.Lexer.Tocken != LCloseParen {
		results = append(results, parser.ParseAssignmentExpression())
		if parser.Lexer.Tocken != LComma {
			break
		}
		parser.Lexer.Next()
	}
	return NewEArgumentList(results)
}

func (parser *Parser) parsePrimaryExpression() (Expr) {
	switch parser.Lexer.Tocken {
	case LOpenBrace:
		parser.Lexer.Next()
		expr := parser.ParseExpression()
		parser.Lexer.Next()
		return expr
	case LIdentifier:
		return NewEIdentifier(parser.Lexer.Value)
	case LInteger:
		return NewEIntegerLiteral(parser.Lexer.Value)
	case LFloat:
		return NewEFloatingPointLiteral(parser.Lexer.Value)
	case LStringLiteral:
		return NewEStringLiteral(parser.Lexer.Value)
	case LFalse:
		return NewEFalse()
	case LTrue:
		return NewETrue()
	case LNull:
		return NewENull()
	case LThis:
		return NewEThis()
	}
	return Expr{}
}