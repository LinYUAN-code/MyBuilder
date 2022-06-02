package ast

import "github.com/LinYUAN-code/MyBuilder/internal/util"

// Statement
type S interface {
	ToJson() string 
}

type Stmts struct {
	Stmts []Stmt
}


type Stmt struct {
	// 用来表示源码中的位置
	locPos uint32
	Data S
}

func NewSFunction(name Expr,args Expr,body Stmt) Stmt {
	return Stmt{
		Data: SFunction{
			Name: name,
			Args: args,
			Body: body,
		},
	}
}

// 字母S开头的表示是Stmt E开头的表示是Expr
// stmt.Data.(type)到时候可以这样取出类型--运用go的动态性
type SFunction struct {
	Body Stmt
	Name Expr
	Args Expr
}
func (s SFunction) ToJson() string {
	return util.StructToJson(util.KV("Name",s.Name.Data.ToJson(),false),
		util.KV("Args",s.Args.Data.ToJson(),false),
		util.KV("Body",s.Body.Data.ToJson(),true))
}

func NewCompoundStatement(stmts []Stmt) Stmt {
	return Stmt{
		Data: SCompoundStatement{
			Stmts: stmts,
		},
	}
}
type SCompoundStatement struct {
	Stmts []Stmt
}
func (s SCompoundStatement) ToJson() string {
	arr := make([]string,0)
	for _,item := range s.Stmts {
		arr = append(arr, item.Data.ToJson())
	}
	return util.JsonArray(arr...)
}

/*
	for(Stmt1;expr;Stmt2) {
		Stmt3
	}

*/

func NewSFor(expr1 Expr,expr2 Expr,expr3 Expr,stmt Stmt) Stmt {
	return Stmt{
		Data: SFor{
			Expr1: expr1,
			Expr2: expr2,
			Expr3: expr3,
			Stmt: stmt,
		},
	}
}
type SFor struct {
	Expr1 Expr
	Expr2 Expr
	Expr3 Expr
	Stmt Stmt
}
func (s SFor) ToJson() string {
	return util.StructToJson(
		util.KV("Expr1",s.Expr1.Data.ToJson(),false),
		util.KV("Expr2",s.Expr2.Data.ToJson(),false),
		util.KV("Expr3",s.Expr3.Data.ToJson(),false),
		util.KV("Body",s.Stmt.Data.ToJson(),true),
	)
}

func NewSForIn(fir Expr,sec Expr,stmt Stmt) Stmt {
	return Stmt{
		Data: SForIn{
			Fir: fir,
			Sec: sec,
			Stmt: stmt,
		},
	}
}
type SForIn struct {
	Fir Expr
	Sec Expr
	Stmt Stmt
}
func (s SForIn) ToJson() string {
	return "SForIn"
}


func NewSIf(expr Expr,pass Stmt,fail Stmt) Stmt {
	return Stmt{
		Data: SIf{
			Expr: expr,
			Pass: pass,
			Fail: fail,
		},
	}
}

type SIf struct {
	Expr Expr
	Pass Stmt
	Fail Stmt
}
func (s SIf) ToJson() string {
	if IsSEmpty(s.Fail) {
		return util.StructToJson(
			util.KV("Expr",s.Expr.Data.ToJson(),false),
			util.KV("Pass",s.Pass.Data.ToJson(),true),
		)
	} else {
		return util.StructToJson(
			util.KV("Expr",s.Expr.Data.ToJson(),false),
			util.KV("Pass",s.Pass.Data.ToJson(),true),
			util.KV("Fail",s.Fail.Data.ToJson(),true),
		)
	}

}

func NewSWhile(expr Expr,body Stmt) Stmt {
	return Stmt{
		Data: SWhile{
			Expr: expr,
			Body: body,
		},
	}
}

type SWhile struct {
	Expr Expr
	Body Stmt
}
func (s SWhile) ToJson() string {
	return "SWhile"
}


func NewSBreak() Stmt {
	return Stmt{
		Data: SBreak{},
	}
}
type SBreak struct {

}
func (s SBreak) ToJson() string {
	return "SBreak"
}


func NewSContinue() Stmt {
	return Stmt{
		Data: SContinue{},
	}
}
type SContinue struct {

}
func (s SContinue) ToJson() string {
	return "SContinue"
}


/*
	with(Expression)Statement
*/
func NewSWith(fir Expr,sec Stmt) Stmt {
	return Stmt{
		Data: SWith{
			Fir: fir,
			Sec: sec,
		},
	}
}
type SWith struct {
	Fir Expr
	Sec Stmt
}
func (s SWith) ToJson() string {
	return "SWith"
}

// return ExpressionOpt
func NewSReturn(fir Expr) Stmt {
	return Stmt{
		Data: SReturn{
			Fir: fir,
		},
	}
}
type SReturn struct {
	Fir Expr
}
func (s SReturn) ToJson() string {
	return util.StructToJson(
		util.KV("returnValue", s.Fir.Data.ToJson(), false),
	)
}

func NewSVariablesOrExpression(fir Expr) Stmt {
	return Stmt{
		Data: SVariablesOrExpression{
			Fir: fir,
		},
	}
}
type SVariablesOrExpression struct {
	Fir Expr
}
func (s SVariablesOrExpression) ToJson() string {
	return s.Fir.Data.ToJson()
}

func NewSEmpty() Stmt {
	return Stmt{
		Data: SEmpty{},
	}
}
type SEmpty struct {

}

func (s SEmpty) ToJson() string {
	return "SEmpty"
}