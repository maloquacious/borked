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

// Exit is an exit from a room.
// It defines something like the location the exit leads to.
type Exit struct {
	Door string // id of the door that is the exit?
	Room string // id of the room the exit leads to.
	Text string // text displayed for a #NOEXIT exit?
}

// Objects holds all the objects defined in the game.
type Objects map[string]*Object

// Object defines an object in the game.
type Object struct {
	Id   string
	Bits struct {
		Open bool
	}
}

// Rooms holds all the rooms defined in the game.
type Rooms map[string]*Room

// Room is a location in the game that can be moved into and (hopefully) out of.
type Room struct {
	Id      string
	Descr   RoomDescr
	Exits   map[string]Exit
	Objects []string
	Bits    RoomBits
	Action  func(verb string)
}

// RoomBits determine features of the room.
type RoomBits struct {
	Land   bool
	Light  bool
	Sacred bool
}

// RoomDescr is the long and short text describing the room.
type RoomDescr struct {
	Long  string
	Short string
}

// RoomExits define the exits leaving a room.
type RoomExits struct {
	Direction map[string]string
}
