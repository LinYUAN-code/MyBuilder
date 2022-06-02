package ast

// 定义抽象语法树结构
/*
	这里主要分成三大类
	program
		AST的根节点
	statement
		语句:函数定义，赋值语句，if语句，for语句....
	expression
		表达式:箭头函数，各种值，函数调用....

	关于语句和表达式的区别？
		可以简单理解为 语句 = 关键字 + 表达式
*/


func GenerateJson(program Stmts) string {
	ans := "{\n"
	for _,stmt := range program.Stmts {
		ans += generateJson(stmt) + "\n"
	}
	ans += "}"
	return ans
}


func generateJson(stmt Stmt) string {
	return  stmt.Data.ToJson()
}

func IsSEmpty(stmt Stmt) bool {
	_,ok := stmt.Data.(SEmpty)
	return ok
}

func IsEEmpty(expr Expr) bool {
	_,ok := expr.Data.(EEmpty)
	return ok
}