
// Package slot is generated by gogll. Do not edit. 
//
// Generated by gogll. Do not edit.
//
//  Copyright 2019 Marius Ackerman
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
package slot

import(
	"bytes"
	"fmt"
)

type Label int

type Slot struct {
	NT      string
	Alt     int
	Pos     int
	Symbols []string
	Label 	Label
}

type Index struct {
	NT      string
	Alt     int
	Pos     int
}

func GetAlternates(nt string) []Label {
	alts, exist := alternates[nt]
	if !exist {
		panic(fmt.Sprintf("Invalid NT %s", nt))
	}
	return alts
}

func GetLabel(nt string, alt, pos int) Label {
	l, exist := slotIndex[Index{nt,alt,pos}]
	if exist {
		return l
	}
	panic(fmt.Sprintf("Error: no slot label for NT=%s, alt=%d, pos=%d", nt, alt, pos))
}

func (l Label) EoR() bool {
	return l.Slot().EoR()
}

func (l Label) Head() string {
	return l.Slot().NT
}

func (l Label) Index() Index {
	s := l.Slot()
	return Index{s.NT, s.Alt, s.Pos}
}

func (l Label) Alternate() int {
	return l.Slot().Alt
}

func (l Label) Pos() int {
	return l.Slot().Pos
}

func (l Label) Slot() *Slot {
	s, exist := slots[l]
	if !exist {
		panic(fmt.Sprintf("Invalid slot label %d", l))
	}
	return s
}

func (l Label) String() string {
	return l.Slot().String()
}

func (l Label) Symbols() []string {
	return l.Slot().Symbols
}

func (s *Slot) EoR() bool {
	return s.Pos >= len(s.Symbols)
}

func (s *Slot) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s : ", s.NT)
	for i, sym := range s.Symbols {
		if i == s.Pos {
			fmt.Fprintf(buf, "∙")
		}
		fmt.Fprintf(buf, "%s ", sym)
	}
	if s.Pos >= len(s.Symbols) {
		fmt.Fprintf(buf, "∙")
	}
	return buf.String()
}

const(
	A0R0 Label = iota
	A0R1
	B0R0
	B0R1
	S0R0
	S0R1
	S0R2
	S1R0
	S1R1
	S1R2
	S2R0
)

var slots = map[Label]*Slot{ 
	A0R0:&Slot{"A", 0, 0, []string{ `a` }, A0R0 },
	A0R1:&Slot{"A", 0, 1, []string{ `a` }, A0R1 },
	B0R0:&Slot{"B", 0, 0, []string{ `letter` }, B0R0 },
	B0R1:&Slot{"B", 0, 1, []string{ `letter` }, B0R1 },
	S0R0:&Slot{"S", 0, 0, []string{ `A`,`S` }, S0R0 },
	S0R1:&Slot{"S", 0, 1, []string{ `A`,`S` }, S0R1 },
	S0R2:&Slot{"S", 0, 2, []string{ `A`,`S` }, S0R2 },
	S1R0:&Slot{"S", 1, 0, []string{ `B`,`S` }, S1R0 },
	S1R1:&Slot{"S", 1, 1, []string{ `B`,`S` }, S1R1 },
	S1R2:&Slot{"S", 1, 2, []string{ `B`,`S` }, S1R2 },
	S2R0:&Slot{"S", 2, 0, []string{  }, S2R0 },
}

var slotIndex = map[Index]Label { 
	Index{ "A",0,0 }: A0R0,
	Index{ "A",0,1 }: A0R1,
	Index{ "B",0,0 }: B0R0,
	Index{ "B",0,1 }: B0R1,
	Index{ "S",0,0 }: S0R0,
	Index{ "S",0,1 }: S0R1,
	Index{ "S",0,2 }: S0R2,
	Index{ "S",1,0 }: S1R0,
	Index{ "S",1,1 }: S1R1,
	Index{ "S",1,2 }: S1R2,
	Index{ "S",2,0 }: S2R0,
}

var alternates = map[string][]Label{ 
	"S":[]Label{ S0R0,S1R0,S2R0 },
	"A":[]Label{ A0R0 },
	"B":[]Label{ B0R0 },
}

