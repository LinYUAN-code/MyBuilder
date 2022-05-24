package parser

import (
	"os"

	"github.com/LinYUAN-code/MyBuilder/internal/lexer"
)

/*
	首先找到具体的JS文法
	https://hepunx.rl.ac.uk/~adye/jsspec11/llr.htm
	补充自己还想要的语法 , 然后计算first 和 follow 集合 当然刚开始只凭感觉写也是没有问题的
	然后直接递归下降分下法一直写下来就好了
*/

type Parser struct {
	Lexer *lexer.Lexer
}

func NewParser(path string) (*Parser,error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil,err
	}
	lexer := lexer.NewLexer(string(data))
	return &Parser{
		Lexer: lexer,
	},nil
}

func (parse *Parser) ShowAllTocken() {
	parse.Lexer.Next()
	for parse.Lexer.Tocken != lexer.LENDOFFILE {
		switch parse.Lexer.Tocken {
		case lexer.LEqual:
			println("=")
		case lexer.LIdentifier:
			println("identifier: ",parse.Lexer.Identifier)
		case lexer.LInteger:
			println("Integer: ",parse.Lexer.Integer)
		case lexer.LSemicolon:
			println(";")
		case lexer.LLet:
			println("let")
		}
		parse.Lexer.Next()
	}
}