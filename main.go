package main;

import (
  "flag"
  "io/ioutil"
  "github.com/tchandy/brainfuck_compiler_go/compiler"
)

func main() {
  var code string 
  var filepath string

  flag.StringVar(&code, "c", "", "The code to run")
  flag.StringVar(&filepath, "f", "", "The File to run the code")
  flag.Parse()

  if code != "" {
    compiler.RunProgram(code)
  } else if filepath != "" {
    content, _ := ioutil.ReadFile(filepath)
    compiler.RunProgram(string(content))
  }
}
