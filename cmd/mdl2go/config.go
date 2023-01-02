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
	"flag"
	"github.com/peterbourgon/ff/v3"
	"log"
	"os"
)

type config struct {
	debug bool
	dump  struct {
		objects bool
		rooms   bool
		symbols bool
	}
	mdl struct {
		root                   string
		act1, act2, act3, act4 string
		dung                   string
		makstr                 string
		rooms                  string
	}
}

func getConfig() *config {
	cfg := &config{}
	cfg.mdl.root = "d:/zork/zork-1978-04/zork/"
	cfg.mdl.act1 = cfg.mdl.root + "act1.1"
	cfg.mdl.act2 = cfg.mdl.root + "act2.1"
	cfg.mdl.act3 = cfg.mdl.root + "act3.1"
	cfg.mdl.act4 = cfg.mdl.root + "act4.1"
	cfg.mdl.dung = cfg.mdl.root + "dung.1"
	cfg.mdl.makstr = cfg.mdl.root + "~~~gsb/makstr.1"
	cfg.mdl.rooms = cfg.mdl.root + "~~~gsb/rooms.1"

	fs := flag.NewFlagSet("my-program", flag.ContinueOnError)
	fs.String("config", "", "config file (optional)")
	fs.BoolVar(&cfg.debug, "debug", cfg.debug, "log debug information")
	fs.BoolVar(&cfg.dump.objects, "objects", cfg.dump.objects, "dump objects")
	fs.BoolVar(&cfg.dump.symbols, "symbols", cfg.dump.symbols, "dump symbols")
	fs.BoolVar(&cfg.dump.rooms, "rooms", cfg.dump.rooms, "dump rooms")
	if err := ff.Parse(fs, os.Args[1:],
		ff.WithEnvVarPrefix("MDL2GO"),
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ff.PlainParser),
	); err != nil {
		log.Fatal(err)
	}

	return cfg
}
