package lexer

import (
	"unicode/utf8"
)

/*
	使用"unicode/utf8"读取utf8码点 然后直接进行状态机模型找到具体的Tocken就好了
*/
type T uint32
// 记录词法符号
const (
	LENDOFFILE T  = iota

	LIdentifier 
	
	LInteger

	LSemicolon

	LEqual

	LLet
	LFor
)


type Lexer struct {

	// 当前读取的Tocken
	Tocken T

	// 如果T 是 LIdentifier 那么这里记录的就是变量名
	Identifier string

	// 如果Number整数 负责记录整数大小
	Integer int64 

	// 记录utf8码点
	CodePoint rune

	// 输入的文件文本
	content string

	// 下一个开始的读入位置
	current uint32

	// 当前codePoint 的开始和结束位置
	start uint32
	end uint32

}

func NewLexer(content string) (*Lexer) {
	return &Lexer{
		content: content,
		current: 0,
	}
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

var keywords = map[string] T {
	"let": LLet,
	"for": LFor,
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
	base := 10
	integer := 0
	for lexer.CodePoint >= '0' && lexer.CodePoint <= '9' {
		lexer.ReadUtf8()
		integer = integer*base + (int(lexer.CodePoint) - '0')
	}
	lexer.Tocken = LInteger
	lexer.Integer = int64(integer)
}

func (lexer *Lexer) newLine() {

}

// 获取下一个Tocken
func (lexer *Lexer) Next() {
	lexer.ReadUtf8()
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
		case '=':
			lexer.Tocken = LEqual
		case '1', '2', '3', '4', '5', '6', '7', '8', '9': //简单支持十进制整数
			lexer.parseNumber()
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
				lexer.Identifier = contents
			}
		}
		return 
	}
}