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

package bork

type ROOMS map[string]*ROOM

type ROOM struct {
	Descr   ROOMDESCR
	Exits   map[string]EXIT
	Objects []string
	Bits    ROOMBITS
	Action  func(verb string)
}
type ROOMBITS struct {
	Land   bool
	Light  bool
	Sacred bool
}
type ROOMDESCR struct {
	Long  string
	Short string
}
type ROOMEXITS struct {
	Direction map[string]string
}

type EXIT struct {
	Door string
	Room string
	Text string
}

var rooms = map[string]ROOM{
	"WHOUS": {
		Descr: ROOMDESCR{
			Short: "West of House",
			Long:  "You are in an open field west of a big white house, with a boarded\nfront door.",
		},
		Exits: map[string]EXIT{
			"NORTH": {Room: "NHOUS"}, "SOUTH": {Room: "SHOUS"}, "WEST": {Room: "FORE1"},
			"EAST": {Room: "#NEXIT", Text: "The door is locked, and there is evidently no key."},
		},
		Objects: []string{"FDOOR", "MAILB"},
		Bits:    ROOMBITS{Land: true, Light: true, Sacred: true},
	},
	"NHOUS": {
		Descr: ROOMDESCR{
			Short: "North of House",
			Long:  "You are facing the north side of a white house.  There is no door here,\nand all the windows are barred.",
		},
		Exits: map[string]EXIT{
			"WEST": {Room: "WHOUS"}, "EAST": {Room: "EHOUS"}, "NORTH": {Room: "FORE3"},
			"SOUTH": {Room: "#NEXIT", Text: "The windows are all barred."},
		},
		Bits: ROOMBITS{Land: true, Light: true, Sacred: true},
	},
	"SHOUS": {
		Descr: ROOMDESCR{
			Short: "South of House",
			Long:  "You are facing the south side of a white house. There is no door here,\nand all the windows are barred.",
		},
		Exits: map[string]EXIT{
			"WEST": {Room: "WHOUS"}, "EAST": {Room: "EHOUS"}, "SOUTH": {Room: "FORE2"},
			"NORTH": {Room: "#NEXIT", Text: "The windows are all barred."},
		},
		Bits: ROOMBITS{Land: true, Light: true, Sacred: true},
	},
	"EHOUS": {
		Descr: ROOMDESCR{Short: "Behind House"},
		Exits: map[string]EXIT{
			"NORTH": {Room: "NHOUS"}, "SOUTH": {Room: "SHOUS"}, "EAST": {Room: "CLEAR"},
			"WEST":  {Door: "WINDO"},
			"ENTER": {Door: "WINDO"},
		},
		Objects: []string{"WINDO"},
		Bits:    ROOMBITS{Land: true, Light: true, Sacred: true},
		Action:  eastHouse,
	},
}

type DOOR struct {
	Id   string
	From string
	To   string
}

var (
	KITCHEN_WINDOW = DOOR{Id: "WINDO", From: "KITCH", To: "EHOUS"}
)
