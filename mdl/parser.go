// borked - a broken clone of a great game
// Copyright (c) 2023 Michael D Henderson
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package mdl

import (
	"fmt"
	"os"
	"strconv"
)

type PARSER struct {
	b *BUFFER
}

func Parser(name string) (*PARSER, error) {
	b, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("parser: %w", err)
	}
	return &PARSER{b: &BUFFER{buffer: b, line: 1, col: 1}}, nil
}

func (p *PARSER) Read() (*CELL, error) {
	var stack []*CELL

	for token := p.b.next(); token != nil; token = p.b.next() {
		switch token.Kind {
		case "number":
			cell := &CELL{Line: token.Line, Kind: token.Kind, Number: token.Number}
			if len(stack) == 0 {
				return cell, nil
			}
			stack[len(stack)-1].List = append(stack[len(stack)-1].List, cell)
		case "symbol":
			cell := &CELL{Line: token.Line, Kind: token.Kind, Symbol: string(token.Value)}
			if len(stack) == 0 {
				return cell, nil
			}
			stack[len(stack)-1].List = append(stack[len(stack)-1].List, cell)
		case "text":
			cell := &CELL{Line: token.Line, Kind: token.Kind, Text: string(token.Value)}
			if len(stack) == 0 {
				return cell, nil
			}
			stack[len(stack)-1].List = append(stack[len(stack)-1].List, cell)
		case "olist":
			stack = append(stack, &CELL{Line: token.Line, Kind: "list"})
		case "clist":
			if len(stack) == 0 {
				return nil, fmt.Errorf("unexpected clist")
			}
			cell := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				return cell, nil
			}
			stack[len(stack)-1].List = append(stack[len(stack)-1].List, cell)
		}
	}

	if len(stack) != 0 {
		for n, c := range stack {
			fmt.Printf("%4d: %5d: %s\n", n, c.Line, c.Kind)
		}
		return nil, fmt.Errorf("unexpected end-of-input")
	}

	return nil, nil
}

func Parse(name string) error {
	b, err := os.ReadFile(name)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}
	buffer := &BUFFER{buffer: b, line: 1, col: 1}
	for !buffer.iseof() {
		token := buffer.next()
		if token == nil {
			fmt.Printf("token: %6d: end-of-text\n", buffer.line)
			return nil
		}
		if token.Error != nil {
			fmt.Printf("%+v\n", token.Error)
			return token.Error
		}
		//switch token.Kind {
		//case "number":
		//	fmt.Printf("%s: %6d: %12d\n", token.Kind, token.Line, token.Number)
		//case "pname":
		//	fmt.Printf("%s: %6d: %s\n", token.Kind, token.Line, string(token.Value))
		//case "qtext":
		//	fmt.Printf("%s: %6d: %s\n", token.Kind, token.Line, string(token.Value))
		//default:
		//	fmt.Printf("%s: %6d: %s\n", token.Kind, token.Line, string(token.Value))
		//}
	}
	return nil
}

const (
	ETX byte = 3  // end-of-text
	FF  byte = 12 // form-feed
)

// TOKEN is
type TOKEN struct {
	Line   int
	Kind   string
	Value  []byte
	Number int
	Error  error
}

type BUFFER struct {
	line, col, pos int
	buffer         []byte
}

func (b *BUFFER) next() *TOKEN {
	// consume spaces and comments
	for !b.iseof() {
		if isspace(b.chpeek()) { // white space
			b.chnext()
			continue
		} else if b.chpeek() == ';' { // comment block
			b.chnext()
			for isspace(b.chpeek()) {
				b.chnext()
			}
			if b.chpeek() != '"' {
				return &TOKEN{Line: b.line, Error: fmt.Errorf("assert(comment must be quoted)")}
			}
			b.chnext()
			if t := b.quotedText(); t.Error != nil {
				return t
			}
			continue
		}
		break
	}
	if b.iseof() {
		return nil
	}
	line, _, ch := b.line, b.pos, b.chnext()
	switch ch {
	case '(':
		return &TOKEN{Line: line, Value: []byte{ch}, Kind: "olist"}
	case ')':
		return &TOKEN{Line: line, Value: []byte{ch}, Kind: "clist"}
	case '<':
		return &TOKEN{Line: line, Value: []byte{ch}, Kind: "olist"}
	case '>':
		return &TOKEN{Line: line, Value: []byte{ch}, Kind: "clist"}
	case '[':
		return &TOKEN{Line: line, Value: []byte{ch}, Kind: "olist"}
	case ']':
		return &TOKEN{Line: line, Value: []byte{ch}, Kind: "clist"}
	case '.':
		return &TOKEN{Line: line, Value: []byte{ch}}
	case '%':
		return &TOKEN{Line: line, Value: []byte{ch}}
	case '!':
		return &TOKEN{Line: line, Value: []byte{ch}}
	case '\'':
		return &TOKEN{Line: line, Value: []byte{ch}}
	case '"':
		return b.quotedText()
	}
	t := b.symbol(ch)
	if i, err := strconv.Atoi(string(t.Value)); err == nil {
		t.Number, t.Kind = i, "number"
	}
	return t
}

func (b *BUFFER) chnext() byte {
	ch := b.chpeek()
	if ch != ETX {
		if ch == '\n' {
			b.line, b.col = b.line+1, 0
		}
		b.pos, b.col = b.pos+1, b.col+1
	}
	return ch
}

func (b *BUFFER) chpeek() byte {
	if b.iseof() {
		return ETX
	}
	return b.buffer[b.pos]
}

func (b *BUFFER) eatln() *TOKEN {
	line, start := b.line, b.pos-1
	for !b.iseof() && b.chpeek() != '\n' {
		b.chnext()
	}
	if b.chpeek() == '\n' {
		b.chnext()
	}
	return &TOKEN{Line: line, Value: b.buffer[start:b.pos]}
}

func (b *BUFFER) iseof() bool {
	return b == nil || !(b.pos < len(b.buffer)) || b.buffer[b.pos] == ETX
}

func (b *BUFFER) quotedText() *TOKEN {
	line, start := b.line, b.pos-1
	for !b.iseof() && b.chpeek() != '"' {
		if b.chpeek() == '\\' { // escaped character
			b.chnext()
		}
		b.chnext()
	}
	if b.chpeek() != '"' { // missing the quote
		return &TOKEN{Line: line, Kind: "text", Error: fmt.Errorf("unterminated quoted text")}
	}
	// consume that closing quote
	b.chnext()
	return &TOKEN{Line: line, Value: b.buffer[start:b.pos], Kind: "text"}
}

// symbol reads the next token as a symbol
func (b *BUFFER) symbol(ch byte) *TOKEN {
	line, start := b.line, b.pos-1
	//fmt.Printf("%d: %d: %c\n", b.line, b.col, ch)
	if ch == '\\' { // symbol starts with an escaped character!
		b.chnext()
	}
	for !isdelim(b.chpeek()) {
		//fmt.Printf("%d: %d: %c\n", b.line, b.col, b.chpeek())
		if b.chpeek() == '\\' { // escaped character
			b.chnext()
		}
		b.chnext()
	}
	return &TOKEN{Line: line, Value: b.buffer[start:b.pos], Kind: "symbol"}
}

// isdelim returns true only if the byte is a delimiter.
func isdelim(ch byte) bool {
	switch ch {
	case '(', ')':
		return true
	case '<', '>':
		return true
	case '[', ']':
		return true
	case ',':
		return true
	case '#':
		return true
	case ';':
		return true
	case '%':
		return true
	case '!':
		return true
	case '\'':
		return true
	case ETX:
		return true
	}
	return isspace(ch)
}

// isspace returns true if the current byte is white space.
func isspace(b byte) bool {
	return b == ' ' || b == '\n' || b == '\r' || b == '\t' || b == ETX || b == FF
}
