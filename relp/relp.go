package relp

import (
	"bufio"
	"compiler/constants"
	"compiler/evaluator"
	"compiler/lexer"
	"compiler/parser"
	"fmt"
	"io"
)

func StartRELP(input io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(input)
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
		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
