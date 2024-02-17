package relp

import (
	"bufio"
	"compiler/constants"
	"compiler/lexer"
	"compiler/token"
	"fmt"
	"io"
)

func StartRELP(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)
	for {
		fmt.Printf(constants.PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
