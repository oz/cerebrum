package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Program struct {
	code    []byte
	tape    []byte
	codePos int
	tapePos int
}

func (p *Program) growTape() {
	if p.tapePos >= len(p.tape) {
		p.tape = append(p.tape, 0)
	}
}

func (p *Program) Run(skip bool) bool {
	for p.tapePos >= 0 && p.codePos < len(p.code) {
		p.growTape()

		// Jump operators
		if p.code[p.codePos] == '[' {
			p.codePos++
			oldPos := p.codePos
			for p.Run(p.tape[p.tapePos] == 0) {
				p.codePos = oldPos
			}
		} else if p.code[p.codePos] == ']' {
			return p.tape[p.tapePos] != 0
		} else if !skip {
			// Simple operators
			switch p.code[p.codePos] {
			case '+':
				p.tape[p.tapePos]++
			case '-':
				p.tape[p.tapePos]--
			case '>':
				p.tapePos += 1
			case '<':
				p.tapePos -= 1
			case '.':
				fmt.Print(string(p.tape[p.tapePos]))
			case ',':
				b := make([]byte, 1)
				os.Stdin.Read(b)
				p.tape[p.tapePos] = b[0]
			default:
				// Ignore unknown chars
			}
		}
		p.codePos++
	}
	return skip
}

func main() {
	if len(os.Args) == 1 {
		buf := make([]byte, 1000)
		os.Stdin.Read(buf)
		p := &Program{code: buf}
		p.Run(false)
	}

	for _, name := range os.Args[1:] {
		buf, err := ioutil.ReadFile(name)
		if err != nil {
			panic(err)
		}
		p := &Program{code: buf}
		p.Run(false)
	}
}
