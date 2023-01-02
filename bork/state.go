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

type State struct {
	turn     int
	isDone   bool
	location int
}

func New() *State {
	return &State{}
}

func (s *State) Eval(words ...string) []string {
	if s.isDone || len(words) == 0 {
		return nil
	} else if words[0] == "/quit" {
		s.isDone = true
	}
	return []string{"ok"}
}

func (s *State) IsDone() bool {
	return s.isDone
}
