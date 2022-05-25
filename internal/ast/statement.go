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

// 字母S开头的表示是Stmt E开头的表示是Expr
// stmt.Data.(type)到时候可以这样取出类型--运用go的动态性
type SFunction struct {
	Body Stmts
	Name Expr
	Args []Expr
}


/*
	for(Stmt1;expr;Stmt2) {
		Stmt3
	}

*/
type SFor struct {
	Stmt1 Stmts
	Expr Expr
	Stmt2 Stmts
	Stmt3 Stmts
}

type SIf struct {
	Expr Expr
	Pass Stmts
	Fail Stmts
}
