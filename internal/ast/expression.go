package ast

import "github.com/LinYUAN-code/MyBuilder/internal/util"

// expression
type E interface {
	ToJson() string
}

type Opt uint8
const (
	Dot Opt = iota
	None
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

var optToString = map[Opt]string {
	Dot: ".",
	None: "none",
	Braket: "[]",
	Paren: "()",
	Plus: "+",
	Minus: "-",
	INSTANCEOF: "instanceof",

	LShift: "<<",
	RShift: ">>",

	// unary
	NPlusPlus: "++ in left",  //和lexer里面的LPlusPlus 冲突
	NMinusMinus: "-- in left",
	RPlusPlus: "++",
	RMinusMinus: "--",
	TYPEOF: "typeof",
	New: "new",
	Delete: "delete",
	
	// AssignmentOperator 
	PlusEqual: "+=",
	MinusEqual: "-=",

	Caret: "^", // ^ 
	Bar: "|",
	And: "&",
	BarBar: "||",
	AndAnd: "&&",

	ExclamationEqual: "!=",
	ExclamationEqualEqual: "!==",
	EqualEuqal: "==",
	EqualEuqalEuqal: "===",

	Less: "<",
	LessEqual: "<=",
	Greater: ">",
	GreaterEqual: ">=",

	Multi: "*",
}

type Declear uint8
const (
	Let Declear = iota
	Const 
	Var
)

type Expr struct {
	locPos uint32
	Data E
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
func (e EParenExpression) ToJson() string {
	return "EParenExpression"
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
func (e EIdentifier) ToJson() string {
	return e.Value
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
func (e EIntegerLiteral) ToJson() string {
	return e.Value
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
func (e EFloatingPointLiteral) ToJson() string {
	return e.Value
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
func (e EStringLiteral) ToJson() string {
	return e.Value
}


func NewEFalse() Expr {
	return Expr{
		Data: EFalse{},
	}
}
type EFalse struct {
}
func (e EFalse) ToJson() string {
	return "EFalse"
}

func NewETrue() Expr {
	return Expr{
		Data: ETrue{},
	}
}
type ETrue struct {
}

func (e ETrue) ToJson() string {
	return "ETrue"
}

func NewENull() Expr {
	return Expr{
		Data: ENull{},
	}
}
type ENull struct {
}

func (e ENull) ToJson() string {
	return "ENull"
}

func NewEThis() Expr {
	return Expr{
		Data: EThis{},
	}
}
type EThis struct {
}

func (e EThis) ToJson() string {
	return "EThis"
}

func NewEMember(fir Expr,opt Opt,sec Expr) Expr {
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
	Opt Opt
	Sec Expr
}
func (e EMember) ToJson() string {
	if e.Opt == None {
		return e.Fir.Data.ToJson() 
	} else {
		return e.Fir.Data.ToJson() + optToString[e.Opt] + e.Sec.Data.ToJson()
	}
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

func (e EConstructorCall) ToJson() string {
	return "EConstructorCall"
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

func (e EConstructor) ToJson() string {
	return "EConstructor"
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
func (e EUnary) ToJson() string {
	return "EUnary"
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
func (e EDualCal) ToJson() string {
	return util.StructToJson(
		util.KV("Fir",e.Fir.Data.ToJson(),false),
		util.KV("Opt",optToString[e.Opt],false),
		util.KV("Sec",e.Sec.Data.ToJson(),false),
	)
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

func (e EConditional) ToJson() string {
	return "EConditional"
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
func (e  EExpressionComma) ToJson() string {
	return "EExpressionComma"
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
func (e EDeclearVariable) ToJson() string {
	return "EDeclearVariable"
}


func NewEVariables(variables []Expr) Expr {
	return Expr{
		Data: EVariables{
			Variables: variables,
		},
	}
}
type EVariables struct {
	Variables []Expr
}

func (e EVariables) ToJson() string {
	return "EVariables"
}


func NewEVariable(fir Expr,sec Expr) Expr {
	return Expr{
		Data: EVariable{
			Identifier: fir,
			AssignmentExpression: sec,
		},
	}
}

type EVariable struct {
	Identifier Expr
	AssignmentExpression Expr
}

func (e EVariable) ToJson() string {
	return "EVariable"
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
func (e EConditionalExpression) ToJson() string {
	return "EConditionalExpression"
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
func (e EArgumentList) ToJson() string {
	arr := make([]string,0)
	for _,item := range e.AssignmentExpressions {
		arr = append(arr, item.Data.ToJson())
	}
	return util.JsonArray(arr...)
}


func NewEExpression(exprs []Expr) Expr {
	return Expr{
		Data: EExpression{
			Exprs: exprs,
		},
	}
}
type EExpression struct {
	Exprs []Expr
}
func (e EExpression) ToJson() string {
	arr := make([]string,0)
	for _,item := range e.Exprs {
		arr = append(arr, item.Data.ToJson())
	}
	return util.JsonArray(arr...)
}

func NewEEmpty() Expr {
	return Expr{
		Data: EEmpty{},
	}
}
type EEmpty struct {

}

func (e EEmpty) ToJson() string {
	return "null"
}


func NewEArgs(args []Expr) Expr {
	return Expr{
		Data: EArgs{
			Args: args,
		},
	}
} 
type EArgs struct {
	Args []Expr
}

func (e EArgs) ToJson() string {
	arr := make([]string,0)
	for _,item := range e.Args {
		arr = append(arr, item.Data.ToJson())
	}
	return util.JsonArray(arr...)
}