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

var rooms = map[string]Room{
	"WHOUS": {
		Descr: RoomDescr{
			Short: "West of House",
			Long:  "You are in an open field west of a big white house, with a boarded\nfront door.",
		},
		Exits: map[string]Exit{
			"NORTH": {Room: "NHOUS"}, "SOUTH": {Room: "SHOUS"}, "WEST": {Room: "FORE1"},
			"EAST": {Room: "#NEXIT", Text: "The door is locked, and there is evidently no key."},
		},
		Objects: []string{"FDOOR", "MAILB"},
		Bits:    RoomBits{Land: true, Light: true, Sacred: true},
	},
	"NHOUS": {
		Descr: RoomDescr{
			Short: "North of House",
			Long:  "You are facing the north side of a white house.  There is no door here,\nand all the windows are barred.",
		},
		Exits: map[string]Exit{
			"WEST": {Room: "WHOUS"}, "EAST": {Room: "EHOUS"}, "NORTH": {Room: "FORE3"},
			"SOUTH": {Room: "#NEXIT", Text: "The windows are all barred."},
		},
		Bits: RoomBits{Land: true, Light: true, Sacred: true},
	},
	"SHOUS": {
		Descr: RoomDescr{
			Short: "South of House",
			Long:  "You are facing the south side of a white house. There is no door here,\nand all the windows are barred.",
		},
		Exits: map[string]Exit{
			"WEST": {Room: "WHOUS"}, "EAST": {Room: "EHOUS"}, "SOUTH": {Room: "FORE2"},
			"NORTH": {Room: "#NEXIT", Text: "The windows are all barred."},
		},
		Bits: RoomBits{Land: true, Light: true, Sacred: true},
	},
	"EHOUS": {
		Descr: RoomDescr{Short: "Behind House"},
		Exits: map[string]Exit{
			"NORTH": {Room: "NHOUS"}, "SOUTH": {Room: "SHOUS"}, "EAST": {Room: "CLEAR"},
			"WEST":  {Door: "WINDO"},
			"ENTER": {Door: "WINDO"},
		},
		Objects: []string{"WINDO"},
		Bits:    RoomBits{Land: true, Light: true, Sacred: true},
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
