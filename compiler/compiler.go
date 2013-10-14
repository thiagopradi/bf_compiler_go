package compiler;

import (
  "fmt"
)

func RunProgram(code string) []int {
  var memory = make([]int, 9999)
  var position = 0
  var loop_tokens = Parse(code)
  var stack = make([]int, 0, 1000)
  var pc = 0

  for ; pc < len(code);  {
    char := string(code[pc])

    if CharIsAPrimitive(char) {
      switch char  {
        case "+": 
        memory[position] += 1

        case "-": 
        memory[position] -= 1

        case ">": 
        position += 1

        case "<": 
        position -= 1

        case ".": 
        fmt.Print(string(memory[position]))

        case ",": 
        var i int
        fmt.Scan(&i)
        memory[position] = i
      }
    } else if char == "[" {
      if memory[position] > 0 {
        stack = append(stack, pc)
      } else {
        pc = loop_tokens[pc]
      }
    } else if char == "]" {
      pc, stack = stack[len(stack)-1], stack[:len(stack)-1]
      pc = pc - 1
    }

    pc += 1
  }

  return memory
}

func Parse(code string) map[int]int {
  var opening []int
  loop_map := make(map[int]int)
  var begin int

  for i := 0; i < len(code); i++ {
    char := string(code[i])

    if char == "[" {
      opening = append(opening, i)
    } else if char == "]" {
      begin, opening = opening[len(opening)-1], opening[:len(opening)-1]
      loop_map[begin] = i
    }
  }

  return loop_map
}

func CharIsAPrimitive(a string) bool {
  var primitives = []string{"+", "-", "<", ">", ".", ","}

  for _, b := range primitives {
    if b == a {
      return true
    }
  }

  return false
}
