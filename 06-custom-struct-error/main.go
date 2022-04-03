package main

import "fmt"

type ParserError struct {
	parser string
	file   string
	line   int
	pos    int
}

func (p *ParserError) Error() string {
	return fmt.Sprintf("parser:%s, file:%s, line:%d ,pos:%d", p.parser, p.file, p.line, p.pos)
}

func tryParse(expr string) error {
	return &ParserError{
		parser: "expr",
		file:   "expr.awk",
		line:   90,
		pos:    100,
	}
}

func main() {
	err := tryParse("{ print $1}")
	if err != nil {
		fmt.Println(err)
		pe, ok := err.(*ParserError)
		if ok {
			fmt.Printf("line:[%d], pos:[%d]\n", pe.line, pe.pos)
		}
	}
}
