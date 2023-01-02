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

/*
<DEFINE EAST-HOUSE ()
    <COND (<VERB? "LOOK">
	   <TELL
"You are behind the white house.  In one corner of the house there
is a small window which is " ,LONG-TELL1 <COND (<TRNN <FIND-OBJ "WINDO"> ,OPENBIT>
		  		      "open.")
		 		     ("slightly ajar.")>>)>>

*/

func eastHouse(verb string) {
	switch verb {
	case "LOOK":
		obj, ok := findObj("WINDO")
		if !ok {
			panic("assert(findObj('WINDO') != false)")
		}
		tell("You are behind the white house.  In one corner of the house there\nis a small window which is %s.\n", ifstr(obj.OpenBit(), "open", "slightly ajar."))
	}
}
