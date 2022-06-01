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
	NPlusPlus  //和lexer里面的LPlusPlus 冲突
	NMinusMinus
	RPlusPlus
	RMinusMinus
	TYPEOF
	New
	Delete
	
	// AssignmentOperator 
	PlusEqual
	MinusEqual

	Caret // ^ 
	Bar
	And
	BarBar
	AndAnd

	ExclamationEqual
	ExclamationEqualEqual
	EqualEuqal
	EqualEuqalEuqal

	Less
	LessEqual
	Greater
	GreaterEqual

	Multi
)

type Declear uint8
const (
	Let Declear = iota
	Const 
	Var
)

type Expr struct {
	locPos uint32
	Data S
}

func NewEParenExpression(expr Expr) (Expr) {
	return Expr{
		Data: EParenExpression{
			Expr: expr,
		},
	}
}
type EParenExpression struct {
	Expr Expr
}


func NewEIdentifier(value string) (Expr) {
	return Expr{
		Data: EIdentifier{
			Value: value,
		},
	}
}
type EIdentifier struct {
	Value string
}

func NewEIntegerLiteral(value string) (Expr) {
	return Expr{
		Data: EIntegerLiteral{
			Value: value,
		},
	}
}


type EIntegerLiteral struct {
	Value string
}

func NewEFloatingPointLiteral(value string) (Expr) {
	return Expr{
		Data: EFloatingPointLiteral{
			Value: value,
		},
	}
}

type EFloatingPointLiteral struct {
	Value string
}
func NewEStringLiteral(value string) Expr {
	return Expr{
		Data: EStringLiteral{
			Value: value,
		},
	}
}
type EStringLiteral struct {
	Value string
}

func NewEFalse() Expr {
	return Expr{
		Data: EFalse{},
	}
}
type EFalse struct {
}

func NewETrue() Expr {
	return Expr{
		Data: ETrue{},
	}
}
type ETrue struct {
}

func NewENull() Expr {
	return Expr{
		Data: ENull{},
	}
}
type ENull struct {
}

func NewEThis() Expr {
	return Expr{
		Data: EThis{},
	}
}
type EThis struct {
}

func NewEMember(fir Expr,opt *Opt,sec *Expr) Expr {
	return Expr{
		Data: EMember{
			Fir: fir,
			Opt: opt,
			Sec: sec,
		},
	}
}
type EMember struct {
	Fir Expr
	Opt *Opt
	Sec *Expr
}


func NewEConstructorCall(identifier Expr,opt Opt,sec Expr) Expr {
	return Expr{
		Data: EConstructorCall{
			Identifier: identifier,
			Opt: opt,
			Sec: sec,
		},
	}
}
type EConstructorCall struct {
	Identifier Expr
	Opt Opt
	Sec Expr
}


func NewEConstructor(constructorCall Expr,this bool) Expr {
	return Expr{
		Data: EConstructor{
			ConstructorCall: constructorCall,
			This: this,
		},
	}
} 
type EConstructor struct {
	ConstructorCall Expr
	This bool
}

func NewEUnary(neg bool,opt Opt,expr Expr) Expr {
	return Expr{
		Data: EUnary{
			Neg: neg,
			Opt: opt,
			Expr: expr,
		},
	}
}

// 一元操作
type EUnary struct {
	Neg bool //前面是否有负号比如 -++i
	Opt Opt
	Expr Expr
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
func NewEDualCal(fir Expr,opt Opt,sec Expr) Expr {
	return Expr{
		Data: EDualCal{
			Fir: fir,
			Opt: opt,
			Sec: sec,
		},
	}
}
type EDualCal struct {
	Fir Expr
	Opt Opt
	Sec Expr
}


func NewEConditional(fir Expr,sec Expr,thir Expr) Expr {
	return Expr{
		locPos: 1,
		Data: EConditional{
			Fir: fir,
			Sec: sec,
			Thir: thir,
		},
	}
} 
type EConditional struct {
	Fir Expr
	Sec Expr
	Thir Expr
}

func NewEExpressionComma(exprs []Expr) Expr {
	return Expr{
		locPos: 1,
		Data: EExpressionComma{
			Exprs: exprs,
		},
	}
}
type EExpressionComma struct {
	Exprs []Expr
}

func NewEDeclearVariable(opt Declear,variables Expr) Expr {
	return Expr{
		Data: EDeclearVariable{
			Opt: opt,
			Variables: variables,
		},
	}
}

type EDeclearVariable struct {
	Opt Declear
	Variables Expr
}

func NewEVariables(identifiers []Expr,assignmentExpression Expr) Expr {
	return Expr{
		Data: EVariables{
			Identifiers: identifiers,
			AssignmentExpression: assignmentExpression,		
		},
	}
}
type EVariables struct {
	Identifiers []Expr
	AssignmentExpression Expr
}


func NewEConditionalExpression(expr1 Expr,expr2 Expr,expr3 Expr) Expr {
	return Expr{
		Data: EConditionalExpression{
			Expr1: expr1,
			Expr2: expr2,
			Expr3: expr3,
		},
	}
}
// expr1 ? expr2 : expr3
type EConditionalExpression struct {
	Expr1 Expr
	Expr2 Expr
	Expr3 Expr
}


func NewEArgumentList(assignmentExpressions []Expr) Expr {
	return Expr{
		Data: EArgumentList{
			AssignmentExpressions: assignmentExpressions,		
		},
	}
}
type EArgumentList struct {
	AssignmentExpressions []Expr
}

func NewEExpression(exprs []Expr) Expr {
	return Expr{
		Data: EExpressionComma{
			Exprs: exprs,
		},
	}
}
type EExpression struct {
	Exprs []Expr
}


func NewEEmpty() Expr {
	return Expr{
		Data: EEmpty{},
	}
}
type EEmpty struct {

}
