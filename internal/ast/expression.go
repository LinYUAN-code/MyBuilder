package ast




// expression
type E interface {

}

type Opt uint8
const (
	Dot Opt = iota
	Braket
	Paren
	Plus
	Minus
	INSTANCEOF

	LShift
	RShift

	// unary
	LPlusPlus
	LMinusMinus
	RPlusPlus
	RMinusMinus
	TYPEOF
	NEW
	DELETE
)


type Expr struct {
	locPos uint32
	Data S
}

func NewEParenExpression(expr Expr) (EParenExpression) {
	return EParenExpression{
		Expr: expr,
	}
}
type EParenExpression struct {
	Expr Expr
}


func NewEIdentifier(value string) (EIdentifier) {
	return EIdentifier{
		Value: value,
	}
}
type EIdentifier struct {
	Value string
}

func NewEIntegerLiteral(value string) (EIntegerLiteral) {
	return EIntegerLiteral{
		Value: value,
	}
}
type EIntegerLiteral struct {
	Value string
}

type EFloatingPointLiteral struct {
	Value string
}

type EStringLiteral struct {
	Value string
}

type EFalse struct {
}

type ETrue struct {
}


type ENull struct {
}

type EThis struct {
}



type EMember struct {
	Fir Expr
	Opt Opt
	Sec Expr
}



type EConstructorCall struct {
	Identifier Expr
	Opt Opt
	Sec Expr
}


type EConstructor struct {
	ConstructorCall Expr
	This bool
}

// 一元操作
type EUnary struct {
	Neg bool //前面是否有负号比如 -++i
	Opt Opt
	Expr Expr
}


type EMultiplicative struct {
	Fir Expr
	Sec Expr  //这里用EMultiplicative会造成循环定义
}

/*
	注意优先级
	*
	+
	-
	>>
	<<
	instanceof
	==,===
	& ^ | && ||

*/
type EDualCalAnd struct {
	Fir Expr
	Opt Opt
	Sec Expr
}


type EConditional struct {
	Fir Expr
	Sec Expr
	Thir Expr
}

type EExpressionComma struct {
	Exprs []Expr
}



