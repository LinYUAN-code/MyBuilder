package parser

import (
	"os"

	"github.com/LinYUAN-code/MyBuilder/internal/ast"
	. "github.com/LinYUAN-code/MyBuilder/internal/lexer"
)

/*
	首先找到具体的JS文法
	https://hepunx.rl.ac.uk/~adye/jsspec11/llr.htm
	补充自己还想要的语法 , 然后计算first 和 follow 集合 当然刚开始只凭感觉写也是没有问题的
	然后直接递归下降分下法一直写下来就好了
*/



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


// func (parser *Parser) parseArgs() ([]ast.Arg) {
// 	result := make([]ast.Arg,0)
// 	for parser.Lexer.Tocken == LIdentifier {
// 		result = append(result, ast.NewArg(parser.Lexer.Identifier))
// 		parser.Lexer.Next()
// 		if parser.Lexer.Tocken != LComma {
// 			break
// 		}
// 	}
// 	return result
// }

// func (parser *Parser) parserFunction() (ast.SFunction) {
// 	functionStmt := ast.SFunction{}
// 	parser.Lexer.Next()
// 	functionStmt.Name = parser.parseAssignmentExpression()
// 	parser.Lexer.Next() //(
// 	functionStmt.Args = parser.parseArgs()
// 	parser.Lexer.Next() //)
// 	functionStmt.Body = parser.ParseStmts()
// 	parser.Lexer.Next() //}
// 	return functionStmt
// }

// func (parser *Parser) parserFor() (ast.SFor) {
// 	forStmt := ast.SFor{}
// 	return forStmt
// }

// func (parser *Parser) parseIf() (ast.SIf) {
// 	ifStmt := ast.SIf{}
// 	parser.Lexer.Next()
// 	parser.Lexer.Next()
// 	ifStmt.Expr = parser.parseAssignmentExpression()

// 	if parser.Lexer.Tocken == LOpenBrace {
// 		parser.Lexer.Next()
// 		ifStmt.Pass = parser.ParseStmts()
// 		parser.Lexer.Next()
// 	} else {
// 		ifStmt.Pass = parser.ParseStmts()
// 	}

// 	if parser.Lexer.Tocken == LElse {
// 		parser.Lexer.Next()
// 		if parser.Lexer.Tocken == LOpenBrace {
// 			parser.Lexer.Next()
// 			ifStmt.Fail = parser.ParseStmts()
// 			parser.Lexer.Next()
// 		} else {
// 			ifStmt.Fail = parser.ParseStmts()
// 		}
// 	}

// 	return ifStmt
// }

// // 采用递归下降分析法
// func (parser *Parser) ParseStmts() (ast.Stmts) {
// 	program := ast.Stmts{}
// 	program.Stmt = make([]ast.Stmt,0)
// 	parser.Lexer.Next()
// 	for {
// 		switch parser.Lexer.Tocken {
// 		case LENDOFFILE:
// 			return program
// 		case LFuntion:
// 			functionStmt := parser.parserFunction()
// 			program.Stmt = append(program.Stmt, functionStmt)
// 		case LFor:
// 			forStmt := parser.parserFor()
// 			program.Stmt = append(program.Stmt, forStmt)
// 		case LIf:
// 			ifStmt := parser.parseIf()
// 			program.Stmt = append(program.Stmt, ifStmt)

// 		}
// 	}
// }


// func (parser *Parser) ParseExpr() (ast.Expr) {
// 	var expr ast.Expr
// 	switch parser.Lexer.Tocken {
// 	case LIdentifier:
// 		expr = ast.EIdentifier{
// 			Name: parser.Lexer.Identifier,
// 		}

// 	}	

// 	return expr
// }


// func (parser *Parser) parseAssignmentExpression() (ast.Expr) {
// 	expr1 := 
// }

// func (parser *Parser) parseConditionalExpression() (ast.Expr)  {
	
// }

func (parser *Parser) parsePrimaryExpression() (ast.Expr) {
	switch parser.Lexer.Tocken {
	case LOpenBrace:
		
	case LIdentifier:

	case LInteger:

	case LFloat:

	case LStringLiteral:

	case LFalse:
	
	case LTrue:

	case LNull:

	case LThis:
		
	}
}