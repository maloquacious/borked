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

package main

import (
	"fmt"
	"github.com/maloquacious/borked/mdl"
	"log"
	"sort"
)

func main() {
	cfg := getConfig()
	for _, name := range []string{cfg.mdl.dung} {
		if err := run(cfg, name); err != nil {
			log.Fatal(err)
		}
	}
}

func run(cfg *config, name string) error {
	if cfg.debug {
		fmt.Printf("run: filename %s\n", name)
	}

	symbols := make(map[string]int)
	parser, err := mdl.Parser(name)
	if err != nil {
		return err
	}
	for {
		cell, err := parser.Read()
		if err != nil {
			return err
		}
		if cell == nil {
			break
		}
		if cfg.dump.symbols {
			for _, symbol := range cell.Symbols() {
				symbols[symbol] = symbols[symbol] + 1
			}
		}
		if cfg.dump.objects {
			if cell.Kind == "list" && cell.List[0].Kind == "symbol" && cell.List[0].Symbol == "OBJECT" {
				fmt.Printf("object: %3d: %s\n", len(cell.List), cell.String())
			}
		}
		if cfg.dump.rooms {
			if cell.Kind == "list" && cell.List[0].Kind == "symbol" && cell.List[0].Symbol == "ROOM" {
				fmt.Printf("room: %3d: %s\n", len(cell.List), cell.String())
			}
		}
	}
	if cfg.dump.symbols {
		var list []string
		for k := range symbols {
			list = append(list, k)
		}
		sort.Strings(list)
		for k, l := range list {
			fmt.Printf("%5d %-20s %8d\n", k, l, symbols[l])
		}
	}
	return nil
}
