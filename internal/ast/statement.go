package ast





// Statement
type S interface {

}

type Stmts struct {
	Stmt []Stmt
}


type Stmt struct {
	// 用来表示源码中的位置
	locPos uint32
	Data S
}

func NewSFunction(name Expr,args []Expr,body Stmt) Stmt {
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
	Args []Expr
}

func NewCompoundStatement(stmts []Stmt) Stmt {
	return Stmt{
		Data: CompoundStatement{
			Stmts: stmts,
		},
	}
}
type CompoundStatement struct {
	Stmts []Stmt
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


func NewSIf(expr Expr,pass Stmt,fail *Stmt) Stmt {
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
	Fail *Stmt
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


func NewSBreak() Stmt {
	return Stmt{
		Data: SBreak{},
	}
}
type SBreak struct {

}


func NewSContinue() Stmt {
	return Stmt{
		Data: SContinue{},
	}
}
type SContinue struct {

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

func NewSEmpty() Stmt {
	return Stmt{
		Data: SEmpty{},
	}
}
type SEmpty struct {

}