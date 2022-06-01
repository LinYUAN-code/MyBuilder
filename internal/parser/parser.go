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
	fmt.Println("parser 发生语法错误")
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
	Assert(parser.Lexer.Tocken == LOpenParen)
	parser.Lexer.Next()
	result := make([]Expr,0)
	for parser.Lexer.Tocken == LIdentifier {
		result = append(result, NewEIdentifier(parser.Lexer.Value))
		parser.Lexer.Next()
		if parser.Lexer.Tocken != LComma {
			break
		}
		parser.Lexer.Next()
	}
	Assert(parser.Lexer.Tocken == LCloseParen)
	parser.Lexer.Next()
	return result
}

// 采用递归下降分析法
func (parser *Parser) ParseStmts() (Stmts) {
	program := Stmts{}
	program.Stmt = make([]Stmt,0)
	parser.Lexer.Next()
	for parser.Lexer.Tocken != LENDOFFILE{
		program.Stmt = append(program.Stmt, parser.ParseElement())
	}
	return program
}

func (parser *Parser) ParseElement() (Stmt) {
	switch parser.Lexer.Tocken {
	case LFuntion:
		parser.Lexer.Next()
		Assert(parser.Lexer.Tocken==LIdentifier)
		name := NewEIdentifier(parser.Lexer.Value)
		parser.Lexer.Next()
		args := parser.parseArgs()
		body := parser.ParseCompoundStatement()
		return NewSFunction(name,args,body)
	default:
		return parser.ParseStatement()
	}
}

func (parser *Parser) ParseCompoundStatement() Stmt {
	Assert(parser.Lexer.Tocken==LOpenBrace)
	parser.Lexer.Next()
	if parser.Lexer.Tocken == LCloseBrace {
		parser.Lexer.Next()
		return NewSEmpty()
	}
	body := parser.ParseStatements()
	Assert(parser.Lexer.Tocken==LCloseBrace)
	parser.Lexer.Next()
	return NewCompoundStatement(body)
}



func (parser *Parser) ParseStatements() []Stmt {
	result := make([]Stmt,0)
	for parser.Lexer.Tocken != LCloseBrace {
		// 去除多余分号
		result = append(result, parser.ParseStatement())
		parser.DelSemicolon()
	}
	return result
}

func (parser *Parser) DelSemicolon() {
	for parser.Lexer.Tocken == LSemicolon {
		parser.Lexer.Next()
	}
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
			parser.Lexer.Next()
			var fir,sec,thir Expr
			var stmt Stmt
			if parser.Lexer.Tocken == LComma {
				fir = NewEEmpty()
			} else {
				fir = parser.ParseVariablesOrExpression()
			}
			parser.Lexer.Next()
			if parser.Lexer.Tocken == LIn {
				parser.Lexer.Next()
				sec = parser.ParseExpression()
				Assert(parser.Lexer.Tocken == LCloseParen)
				parser.Lexer.Next()
				stmt = parser.ParseStatement()
				return NewSForIn(fir,sec,stmt)
			} else {
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
				Assert(parser.Lexer.Tocken == LCloseParen)
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
			Assert(parser.Lexer.Tocken == LSemicolon)
			parser.Lexer.Next()
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
		parser.Lexer.Next()
		opt = Let
	case LVar:
		parser.Lexer.Next()
		opt = Var
	case LConst:
		parser.Lexer.Next()
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
		results = append(results, parser.ParseVariable())
		if parser.Lexer.Tocken != LComma {
			break
		}
		parser.Lexer.Next()
	}
	return NewEVariables(results)
}

func (parser *Parser) ParseVariable() Expr {
	Assert(parser.Lexer.Tocken == LIdentifier)
	identifier := NewEIdentifier(parser.Lexer.Value)
	parser.Lexer.Next()
	var assignmentExpression Expr
	if parser.Lexer.Tocken == LEqual {
		parser.Lexer.Next()
		assignmentExpression = parser.ParseAssignmentExpression()
	} else {
		assignmentExpression = NewEEmpty()
	}
	return NewEVariable(identifier,assignmentExpression)
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
		parser.Lexer.Next()
		sec := parser.ParseOrExpression()
		return NewEDualCal(fir,BarBar,sec)
	}
	return fir
}

func (parser *Parser) ParseAndExpression() Expr {
	fir := parser.ParseBitwiseOrExpression()
	if parser.Lexer.Tocken == LAndAnd {
		parser.Lexer.Next()
		sec := parser.ParseAndExpression()
		return NewEDualCal(fir,AndAnd,sec)
	}
	return fir
}

func (parser *Parser) ParseBitwiseOrExpression() Expr {
	fir := parser.ParseBitwiseXorExpression()
	if parser.Lexer.Tocken == LBar {
		parser.Lexer.Next()
		sec := parser.ParseBitwiseOrExpression()
		return NewEDualCal(fir,Bar,sec)
	}
	return fir
}

func (parser *Parser) ParseBitwiseXorExpression() Expr {
	fir := parser.ParseBitwiseAndExpression()
	if parser.Lexer.Tocken == LCaret {
		parser.Lexer.Next()

		sec := parser.ParseBitwiseAndExpression()
		return NewEDualCal(fir,Caret,sec)
	}
	return fir
}

func (parser *Parser) ParseBitwiseAndExpression() Expr {
	fir := parser.ParseEqualityExpression()
	if parser.Lexer.Tocken == LAnd {
		parser.Lexer.Next()

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
		parser.Lexer.Next()
		sec := parser.ParseEqualityExpression()
		return NewEDualCal(fir,EqualEuqal,sec)
	case LEqualEqualEqual:
		parser.Lexer.Next()
		sec := parser.ParseEqualityExpression()
		return NewEDualCal(fir,EqualEuqalEuqal,sec)
	case LExclamationEqual:
		parser.Lexer.Next()
		sec := parser.ParseEqualityExpression()
		return NewEDualCal(fir,ExclamationEqual,sec)
	case LExclamationEqualEqual:
		parser.Lexer.Next()
		sec := parser.ParseEqualityExpression()
		return NewEDualCal(fir,ExclamationEqualEqual,sec)
	}
	return fir
}

func (parser *Parser) ParseRelationalExpression() Expr {
	fir := parser.ParseShiftExpression()
	switch parser.Lexer.Tocken {
	case LLess:
		parser.Lexer.Next()

		sec := parser.ParseRelationalExpression()
		return NewEDualCal(fir,Less,sec)
	case LLessEqual:
		parser.Lexer.Next()

		sec := parser.ParseRelationalExpression()
		return NewEDualCal(fir,LessEqual,sec)
	case LGreater:
		parser.Lexer.Next()

		sec := parser.ParseRelationalExpression()
		return NewEDualCal(fir,Greater,sec)
	case LGreaterEqual:
		parser.Lexer.Next()

		sec := parser.ParseRelationalExpression()
		return NewEDualCal(fir,GreaterEqual,sec)
	}
	return fir
}

func (parser *Parser) ParseShiftExpression() Expr {
	fir := parser.ParseAdditiveExpression()
	switch parser.Lexer.Tocken {
	case LLshift:
		parser.Lexer.Next()

		sec := parser.ParseShiftExpression()
		return NewEDualCal(fir,LShift,sec)
	case LRshift:
		parser.Lexer.Next()

		sec := parser.ParseShiftExpression()
		return NewEDualCal(fir,RShift,sec)
	}
	return fir
}

func (parser *Parser) ParseAdditiveExpression() Expr {
	fir := parser.ParseMultiplicativeExpression()
	switch parser.Lexer.Tocken {
	case LPlus:
		parser.Lexer.Next()
		sec := parser.ParseAdditiveExpression()
		return NewEDualCal(fir,Plus,sec)
	case LMinus:
		parser.Lexer.Next()
		sec := parser.ParseAdditiveExpression()
		return NewEDualCal(fir,Minus,sec)
	}
	return fir
}

func (parser *Parser) ParseMultiplicativeExpression() Expr {
	fir := parser.ParseUnaryExpression()
	if parser.Lexer.Tocken == LMulti {
		parser.Lexer.Next()

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
		parser.Lexer.Next()

			return NewEUnary(false,RPlusPlus,expr)
		case LMinusMinus:
		parser.Lexer.Next()

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
	parser.Lexer.Next()
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
		parser.Lexer.Next()
		value1 := Dot
		opt = &value1
		sec := parser.ParseMemberExpression()
		return NewEMember(fir,opt,&sec)
	case LOpenBraket:
		parser.Lexer.Next()
		value := Braket
		opt = &value
		Assert(parser.Lexer.Tocken == LCloseBraket)
		sec := parser.ParseExpression()
		return NewEMember(fir,opt,&sec)
	case LOpenParen:
		parser.Lexer.Next()
		value := Paren
		opt = &value
		sec := parser.ParseArgumentList()
		Assert(parser.Lexer.Tocken == LCloseParen)
		parser.Lexer.Next()
		return NewEMember(fir,opt,&sec)
	}
	return NewEMember(fir,opt,se)
}

func (parser *Parser) ParseArgumentList() Expr {
	results := make([]Expr,0)
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
		parser.Lexer.Next()
		return NewEIdentifier(parser.Lexer.Value)
	case LInteger:
		parser.Lexer.Next()
		return NewEIntegerLiteral(parser.Lexer.Value)
	case LFloat:
		parser.Lexer.Next()
		return NewEFloatingPointLiteral(parser.Lexer.Value)
	case LStringLiteral:
		parser.Lexer.Next()
		return NewEStringLiteral(parser.Lexer.Value)
	case LFalse:
		parser.Lexer.Next()
		return NewEFalse()
	case LTrue:
		parser.Lexer.Next()
		return NewETrue()
	case LNull:
		parser.Lexer.Next()
		return NewENull()
	case LThis:
		parser.Lexer.Next()
		return NewEThis()
	}
	return Expr{}
}