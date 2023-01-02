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

import "fmt"

// CELL is either an atom or a list
type CELL struct {
	Line   int
	Kind   string // list, number, symbol, or text
	List   []*CELL
	Number int
	Symbol string
	Text   string
}

/*
	read token
	if end-of-input
		if stack is empty
			return end-of-input
		error - unexpected end-of-input
	if atom
		if stack is empty
			return atom
		append to TOS
		continue
	if open-list
		push list
		continue
	if close-list
		pop stack
		if stack is empty
			return list
		append to TOS
		continue
	error - should be atom or list
*/

func (c *CELL) String() string {
	switch c.Kind {
	case "list":
		s := "("
		for _, cell := range c.List {
			s += cell.String()
		}
		return s + ")"
	case "number":
		return fmt.Sprintf(" %d ", c.Number)
	case "symbol":
		return fmt.Sprintf(" %s ", c.Symbol)
	case "text":
		return fmt.Sprintf(" %s ", c.Text)
	}
	return ""
}

/*
        o  b  o
    r 	          z
 f   M  A  G  I  C   z
 c    W  E   L  L    y
    o             n
        m  p  a
*/

func (c *CELL) Symbols() (symbols []string) {
	switch c.Kind {
	case "list":
		for _, cell := range c.List {
			for _, symbol := range cell.Symbols() {
				symbols = append(symbols, symbol)
			}
		}
	case "symbol":
		symbols = append(symbols, c.Symbol)
	}
	return symbols
}
