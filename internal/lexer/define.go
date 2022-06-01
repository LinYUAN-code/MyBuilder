package lexer

type T uint32
// 记录词法符号
const (
	LENDOFFILE T  = iota

	LIdentifier 
	
	LInteger
	LFloat

	LStringLiteral

	LTrue
	LFalse
	LNull
	LThis

	LSemicolon

	LDot
	LComma

	LEqual
	LEqualEqual
	LEqualEqualEqual


	LOpenBraket
	LCloseBraket

	LLet
	LVar
	LConst

	LFor
	LIf
	LElse
	LWhile
	LBar
	LLess
	LLessEqual


	LGreater
	LGreaterEqual

	LBarBar
	LBarEqual
	LReturn
	LPlus
	LMinus
	LMinusMinus
	LPlusPlus
	LFuntion
	LOpenParen 
	LCloseParen
	LOpenBrace
	LCloseBrace

	LDelete
	LNew
	LTypeof

	LQuestion

	// AssignmentOperator 
	LPlusEqual
	LMinusEqual

	LIn
	
	LAnd
	LAndAnd

	LCaret

	LExclamation //!
	LExclamationEqual //!=
	LExclamationEqualEqual //!==

	LLshift
	LRshift

	LMulti //*

	LBreak

	LContinue

	LWith
)



type Lexer struct {

	// 当前读取的Tocken
	Tocken T

	// 用来记录比如整数,浮点数，变量名的具体字符串
	Value string

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

var keywords = map[string] T {
	"let": LLet,
	"var": LVar,
	"const": LConst,
	"for": LFor,
	"if":  LIf,
	"function": LFuntion,
	"return": LReturn,
	"else": LElse,
	"false": LFalse,
	"true": LTrue,
	"this": LThis,
	"delete": LDelete,
	"new": LNew,
	"typeof": LTypeof,
	"while": LWhile,
	"in": LIn,
	"break": LBreak,
	"continue": LContinue,
	"with": LWith
}

var TToString = map[T] string {
	LENDOFFILE: "LENDOFFILE",

	LIdentifier: "LIdentifier",
	
	LInteger: "LInteger",

	LSemicolon: "LSemicolon",

	LDot: "LDot",
	LComma: "LComma",

	LEqual: "LEqual",
	LEqualEqual: "LEqualEqual",
	LEqualEqualEqual: "LEqualEqualEqual",


	LOpenBraket: "LOpenBraket",
	LCloseBraket: "LCloseBraket",

	LLet: "LLet",
	LFor: "LFor",
	LIf: "LIf",
	LElse: "LElse",
	LBar: "LBar",
	LLess: "LLess",
	LGreater: "LGreater",
	LBarBar: "LBarBar",
	LBarEqual: "LBarEqual",
	LReturn: "LReturn",
	LPlus: "LPlus",
	LMinus: "LMinus",
	LMinusMinus: "LMinusMinus",
	LPlusPlus: "LPlusPlus",
	LFuntion: "LFuntion",
	LOpenParen: "LOpenParen",
	LCloseParen: "LCloseParen",
	LOpenBrace: "LOpenBrace",
	LCloseBrace: "LCloseBrace",
}