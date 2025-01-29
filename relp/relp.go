package relp

import (
	"bufio"
	"compiler/compiler"
	"compiler/constants"
	"compiler/evaluator"
	"compiler/lexer"
	"compiler/object"
	"compiler/parser"
	"compiler/virtualmachine"
	"fmt"
	"io"
)

func StartRELP(input io.Reader, out io.Writer, compilationMode bool) {
	scanner := bufio.NewScanner(input)
	env := object.NewEnviornment()
	for {
		fmt.Printf(constants.PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		if compilationMode {
			comp := compiler.New()
			err := comp.Compile(program)
			if err != nil {
				fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
				continue
			}
			machine := virtualmachine.New(comp.ByteCode())
			err = machine.Run()
			if err != nil {
				fmt.Fprintf(out, "Woops! Bytecode Execution failed:\n %s\n", err)
				continue
			}
			stackTop := machine.StackTop()
			io.WriteString(out, stackTop.Inspect())
			io.WriteString(out, "\n")
		} else {
			evaluated := evaluator.Eval(program, env)
			if evaluated != nil {
				io.WriteString(out, evaluated.Inspect())
				io.WriteString(out, "\n")
			}
		}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
