package lexer

import (
	"unicode/utf8"
)

/*
	使用"unicode/utf8"读取utf8码点 然后直接进行状态机模型找到具体的Tocken就好了
*/


func NewLexer(content string) (*Lexer) {
	lexer :=  &Lexer{
		content: content,
		current: 0,
	}
	lexer.ReadUtf8()
	return lexer
}


func (lexer Lexer) SayHello() {
	println("你好,我是一个小Lexer")
}

func (lexer *Lexer) ReadUtf8() {
	codePoint, width := utf8.DecodeRuneInString(lexer.content[lexer.current:])

	// 读取完毕
	if width == 0 {
		codePoint = -1
	}

	lexer.CodePoint = codePoint
	lexer.end = lexer.current
	lexer.current = lexer.current + uint32(width)
}


func IsIdentifierContinue(codePoint rune) bool {
	switch codePoint {
		case '_', '$', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
			'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
			'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
			'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
			'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			return true
	}

	// All ASCII identifier start code points are listed above
	if codePoint < 0x7F {
		return false
	}

	return false
}

func IsIdentifierStart(codePoint rune) bool {
	switch codePoint {
	case '_', '$',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
		return true
	}

	// All ASCII identifier start code points are listed above
	if codePoint < 0x7F {
		return false
	}

	return false
}



func toKeyWords(key string) T {
	tocken, ok := keywords[key]
	if !ok {
		return LIdentifier
	} else {
		return tocken
	}
}

func (lexer Lexer) Raw() string {
	return lexer.content[lexer.start:lexer.end]
}

func (lexer *Lexer) parseNumber() {
	// 简单的支持一下十进制整数
	if lexer.CodePoint == '.' {
		lexer.ReadUtf8()
		if lexer.CodePoint < '0' || lexer.CodePoint > '9' {
			lexer.Tocken = LDot
			return 
		}
	}
	// base := 10
	// integer := 0
	// for lexer.CodePoint >= '0' && lexer.CodePoint <= '9' {
	// 	lexer.ReadUtf8()
	// 	integer = integer*base + (int(lexer.CodePoint) - '0')
	// }
	for lexer.CodePoint >= '0' && lexer.CodePoint <= '9' {
		lexer.ReadUtf8()
	}
	lexer.Tocken = LInteger
	lexer.Value = lexer.content[lexer.start:lexer.end]
}

func (lexer *Lexer) newLine() {

}

// 获取下一个Tocken
func (lexer *Lexer) Next() {
	for {
		lexer.start = lexer.end
		switch lexer.CodePoint {
		
		case -1:
			lexer.Tocken = LENDOFFILE
		case '\r', '\n': //换行
			lexer.newLine()
			lexer.ReadUtf8()
			continue 
		case ' ', '\t': //空白符
			lexer.ReadUtf8()
			continue
		case ';':
			lexer.Tocken = LSemicolon
			lexer.ReadUtf8()
		case '=':
			lexer.Tocken = LEqual
			lexer.ReadUtf8()
			if lexer.CodePoint == '=' {
				lexer.ReadUtf8()
				if lexer.CodePoint == '=' {
					lexer.Tocken = LEqualEqualEqual 
					lexer.ReadUtf8()
				} else {
					lexer.Tocken = LEqualEqual
				}
			}
		case '|':
			lexer.ReadUtf8()
			switch lexer.CodePoint {
			case '=':
				lexer.Tocken = LBarEqual
				lexer.ReadUtf8()
			case '|':
				lexer.Tocken = LBarBar
				lexer.ReadUtf8()
			default:
				lexer.Tocken = LBar
			}
		case '&':
			lexer.ReadUtf8()
			switch lexer.CodePoint {
			case '&':
				lexer.Tocken = LAndAnd
				lexer.ReadUtf8()
			default:
				lexer.Tocken = LAnd
			}
		case '^':
			lexer.Tocken = LCaret
			lexer.ReadUtf8()
		case '!':
			lexer.ReadUtf8()
			if lexer.CodePoint == '=' {
				lexer.ReadUtf8()
				if lexer.CodePoint == '=' {
					lexer.Tocken = LExclamationEqualEqual
				} else {
					lexer.Tocken = LExclamationEqual
				}
			}
			lexer.Tocken = LExclamation
		case '(':
			lexer.Tocken = LOpenParen
			lexer.ReadUtf8()
		case ')':
			lexer.Tocken = LCloseParen
			lexer.ReadUtf8()
		case '{':
			lexer.Tocken = LOpenBrace
			lexer.ReadUtf8()
		case '}':
			lexer.Tocken = LCloseBrace
			lexer.ReadUtf8()
		case '[':
			lexer.Tocken = LOpenBraket
			lexer.ReadUtf8()
		case ']':
			lexer.Tocken = LCloseBraket
			lexer.ReadUtf8()
		case '+':
			lexer.ReadUtf8()
			switch lexer.CodePoint {
			case '+':
				lexer.Tocken = LPlusPlus
				lexer.ReadUtf8()
			case '=':
				lexer.Tocken = LPlusEqual
				lexer.ReadUtf8()
			default:
				lexer.Tocken = LPlus
			}
		case '-':
			lexer.ReadUtf8()
			switch lexer.CodePoint {
			case '-':
				lexer.Tocken = LMinusMinus
				lexer.ReadUtf8()
			case '=':
				lexer.Tocken = LMinusEqual
				lexer.ReadUtf8()
			default:
				lexer.Tocken = LMinus
			}
		case '*':
			lexer.ReadUtf8()
			switch lexer.CodePoint {
			default:
				lexer.Tocken = LMulti
			}
		case '<':
			lexer.ReadUtf8()
			switch lexer.CodePoint {
			case '=':
				lexer.ReadUtf8()
				lexer.Tocken = LLessEqual
			case '<':
				lexer.ReadUtf8()
				lexer.Tocken = LLshift
			}
			lexer.Tocken = LLess
		case '>':
			lexer.ReadUtf8()
			switch lexer.CodePoint {
			case '=':
				lexer.ReadUtf8()
				lexer.Tocken = LGreaterEqual
			case '>':
				lexer.ReadUtf8()
				lexer.Tocken = LRshift
			}
			if lexer.CodePoint == '=' {

			}
			lexer.Tocken = LGreater
		case ',':
			lexer.Tocken = LComma
			lexer.ReadUtf8()
		case '.','0','1', '2', '3', '4', '5', '6', '7', '8', '9': //简单支持十进制整数
			lexer.parseNumber()
		case '?':
			lexer.Tocken = LQuestion
			lexer.ReadUtf8()
		case '_', '$',  //识别标识符
			'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
			'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
			'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
			'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			lexer.ReadUtf8()
			for IsIdentifierContinue(lexer.CodePoint) {
				lexer.ReadUtf8()
			}
			contents := lexer.Raw()
			lexer.Tocken = toKeyWords(contents)
			if lexer.Tocken == LIdentifier {
				lexer.Value = contents
			}
		case '"','\'':
		}
		return 
	}
}