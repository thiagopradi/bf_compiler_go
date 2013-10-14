package compiler_test

import (
  "github.com/tchandy/brainfuck_compiler_go/compiler"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("compiler", func() {

  Describe("RunProgram", func() {
    Context("for a valid program", func() {
      It("should return the memory correctly", func() {
        memory := compiler.RunProgram("++++++")

        Expect(6).To(Equal(memory[0]))
        Expect(0).To(Equal(memory[1]))
      })

      It("should return the memory correctly with loop", func() {
        memory := compiler.RunProgram("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.")

        Expect(0).To(Equal(memory[0]))
        Expect(87).To(Equal(memory[1]))
      })
    })
  })

  Describe("Parse", func() {
    Context("for a valid program", func() {
      It("should parse and return the correct map when there's only one loop in the program", func() {
        correct_map := make(map[int]int)
        correct_map[3] = 7
        map_var := compiler.Parse("+++[+++]")
        Expect(map_var).To(Equal(correct_map))
      })

      It("should parse and return the correct map when there's more than one loop in the program", func() {
        correct_map := make(map[int]int)
        correct_map[3] = 9
        correct_map[5] = 7
        map_var := compiler.Parse("+++[+[+]+]")
        Expect(map_var).To(Equal(correct_map))
      })

      It("should parse and return the correct map for hello world program", func() {
        correct_map := make(map[int]int)
        correct_map[10] = 41
        map_var := compiler.Parse("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.")
        Expect(map_var).To(Equal(correct_map))
      })

      PIt("should raise a error when the parentesis ins't balanced", func() {
        /* _, err := compiler.Parse("+++[+[++]]]") */
        /* Expect(map_var).To(Equal(correct_map)) */
      })
    })
  })

  Describe("CharIsAPrimitive", func() {
    It("should return false or true when the char is a primitive or not", func() {
      Expect(compiler.CharIsAPrimitive("+")).To(Equal(true))
      Expect(compiler.CharIsAPrimitive("-")).To(Equal(true))
      Expect(compiler.CharIsAPrimitive(">")).To(Equal(true))
      Expect(compiler.CharIsAPrimitive("<")).To(Equal(true))
      Expect(compiler.CharIsAPrimitive(".")).To(Equal(true))
      Expect(compiler.CharIsAPrimitive(",")).To(Equal(true))
      Expect(compiler.CharIsAPrimitive("[")).To(Equal(false))
      Expect(compiler.CharIsAPrimitive("]")).To(Equal(false))
    })
  })   
})
