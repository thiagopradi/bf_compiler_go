package main;

import (
  "flag"
  "github.com/tchandy/brainfuck_compiler_go/compiler"
)

func main() {
  var code string 

  flag.StringVar(&code, "c", "", "The code to run")
  flag.Parse()

  compiler.RunProgram(code)
}
