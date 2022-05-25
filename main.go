package main

import (
	"os"

	"github.com/LinYUAN-code/MyBuilder/internal/parser"
)


func main() {
	println("LRJ builder *★,°*:.☆(￣▽￣)/$:*.°★* 。")
	parser,err := parser.NewParser("./test/fibo.js")
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}
	parser.Lexer.SayHello()
	// program := parser.ParseProgram()
	println(len(program.Stmt))
}