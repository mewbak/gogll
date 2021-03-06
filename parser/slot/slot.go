
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
	Alternate0R0 Label = iota
	Alternate0R1
	Alternate1R0
	Alternate1R1
	Alternate1R2
	Alternate1R3
	Alternate1R4
	Alternate1R5
	Alternates0R0
	Alternates0R1
	Alternates1R0
	Alternates1R1
	Alternates1R2
	Alternates1R3
	Alternates1R4
	Alternates1R5
	CharLiteral0R0
	CharLiteral0R1
	CharLiteral0R2
	CharLiteral0R3
	CharLiteral0R4
	CharLiteral1R0
	CharLiteral1R1
	CharLiteral1R2
	CharLiteral1R3
	EscapedChar0R0
	EscapedChar0R1
	EscapedChar1R0
	EscapedChar1R1
	EscapedChar2R0
	EscapedChar2R1
	EscapedChar3R0
	EscapedChar3R1
	EscapedChar4R0
	EscapedChar4R1
	EscapedChar5R0
	EscapedChar5R1
	GoGLL0R0
	GoGLL0R1
	GoGLL0R2
	GoGLL0R3
	GoGLL0R4
	GoGLL0R5
	Head0R0
	Head0R1
	Head0R2
	Head1R0
	Head1R1
	NTChar0R0
	NTChar0R1
	NTChar1R0
	NTChar1R1
	NTChar2R0
	NTChar2R1
	NTChars0R0
	NTChars0R1
	NTChars1R0
	NTChars1R1
	NTChars1R2
	NTID0R0
	NTID0R1
	NTID1R0
	NTID1R1
	NTID1R2
	NonTerminal0R0
	NonTerminal0R1
	Package0R0
	Package0R1
	Package0R2
	Package0R3
	Package0R4
	Package0R5
	Package0R6
	Package0R7
	Package0R8
	Package0R9
	Rule0R0
	Rule0R1
	Rule0R2
	Rule0R3
	Rule0R4
	Rule0R5
	Rule0R6
	Rule0R7
	Rules0R0
	Rules0R1
	Rules1R0
	Rules1R1
	Rules1R2
	Rules1R3
	Sep0R0
	Sep0R1
	Sep1R0
	Sep1R1
	Sep1R2
	SepChar0R0
	SepChar0R1
	SepChar1R0
	SepChar1R1
	SepE0R0
	SepE0R1
	SepE1R0
	String0R0
	String0R1
	String0R2
	String0R3
	StringChars0R0
	StringChars0R1
	StringChars0R2
	StringChars1R0
	StringChars1R1
	StringChars1R2
	StringChars1R3
	StringChars2R0
	Symbol0R0
	Symbol0R1
	Symbol1R0
	Symbol1R1
	Symbols0R0
	Symbols0R1
	Symbols0R2
	Symbols0R3
	Symbols1R0
	Symbols1R1
	Terminal0R0
	Terminal0R1
	Terminal0R2
	Terminal0R3
	Terminal1R0
	Terminal1R1
	Terminal1R2
	Terminal1R3
	Terminal1R4
	Terminal1R5
	Terminal1R6
	Terminal1R7
	Terminal2R0
	Terminal2R1
	Terminal2R2
	Terminal2R3
	Terminal2R4
	Terminal2R5
	Terminal2R6
	Terminal3R0
	Terminal3R1
	Terminal3R2
	Terminal3R3
	Terminal3R4
	Terminal3R5
	Terminal3R6
	Terminal4R0
	Terminal4R1
	Terminal4R2
	Terminal4R3
	Terminal4R4
	Terminal4R5
	Terminal5R0
	Terminal5R1
	Terminal5R2
	Terminal5R3
	Terminal5R4
	Terminal5R5
	Terminal5R6
	Terminal6R0
	Terminal6R1
	Terminal6R2
	Terminal6R3
	Terminal6R4
	Terminal6R5
	Terminal6R6
	Terminal6R7
	Terminal7R0
	Terminal7R1
	Terminal7R2
	Terminal7R3
	Terminal7R4
	Terminal7R5
	Terminal8R0
	Terminal8R1
	Terminal9R0
	Terminal9R1
)

var slots = map[Label]*Slot{ 
	Alternate0R0:&Slot{"Alternate", 0, 0, []string{ `Symbols` }, Alternate0R0 },
	Alternate0R1:&Slot{"Alternate", 0, 1, []string{ `Symbols` }, Alternate0R1 },
	Alternate1R0:&Slot{"Alternate", 1, 0, []string{ `e`,`m`,`p`,`t`,`y` }, Alternate1R0 },
	Alternate1R1:&Slot{"Alternate", 1, 1, []string{ `e`,`m`,`p`,`t`,`y` }, Alternate1R1 },
	Alternate1R2:&Slot{"Alternate", 1, 2, []string{ `e`,`m`,`p`,`t`,`y` }, Alternate1R2 },
	Alternate1R3:&Slot{"Alternate", 1, 3, []string{ `e`,`m`,`p`,`t`,`y` }, Alternate1R3 },
	Alternate1R4:&Slot{"Alternate", 1, 4, []string{ `e`,`m`,`p`,`t`,`y` }, Alternate1R4 },
	Alternate1R5:&Slot{"Alternate", 1, 5, []string{ `e`,`m`,`p`,`t`,`y` }, Alternate1R5 },
	Alternates0R0:&Slot{"Alternates", 0, 0, []string{ `Alternate` }, Alternates0R0 },
	Alternates0R1:&Slot{"Alternates", 0, 1, []string{ `Alternate` }, Alternates0R1 },
	Alternates1R0:&Slot{"Alternates", 1, 0, []string{ `Alternate`,`SepE`,`|`,`SepE`,`Alternates` }, Alternates1R0 },
	Alternates1R1:&Slot{"Alternates", 1, 1, []string{ `Alternate`,`SepE`,`|`,`SepE`,`Alternates` }, Alternates1R1 },
	Alternates1R2:&Slot{"Alternates", 1, 2, []string{ `Alternate`,`SepE`,`|`,`SepE`,`Alternates` }, Alternates1R2 },
	Alternates1R3:&Slot{"Alternates", 1, 3, []string{ `Alternate`,`SepE`,`|`,`SepE`,`Alternates` }, Alternates1R3 },
	Alternates1R4:&Slot{"Alternates", 1, 4, []string{ `Alternate`,`SepE`,`|`,`SepE`,`Alternates` }, Alternates1R4 },
	Alternates1R5:&Slot{"Alternates", 1, 5, []string{ `Alternate`,`SepE`,`|`,`SepE`,`Alternates` }, Alternates1R5 },
	CharLiteral0R0:&Slot{"CharLiteral", 0, 0, []string{ `\'`,`\\`,`anyof('\"\\nrt)`,`\'` }, CharLiteral0R0 },
	CharLiteral0R1:&Slot{"CharLiteral", 0, 1, []string{ `\'`,`\\`,`anyof('\"\\nrt)`,`\'` }, CharLiteral0R1 },
	CharLiteral0R2:&Slot{"CharLiteral", 0, 2, []string{ `\'`,`\\`,`anyof('\"\\nrt)`,`\'` }, CharLiteral0R2 },
	CharLiteral0R3:&Slot{"CharLiteral", 0, 3, []string{ `\'`,`\\`,`anyof('\"\\nrt)`,`\'` }, CharLiteral0R3 },
	CharLiteral0R4:&Slot{"CharLiteral", 0, 4, []string{ `\'`,`\\`,`anyof('\"\\nrt)`,`\'` }, CharLiteral0R4 },
	CharLiteral1R0:&Slot{"CharLiteral", 1, 0, []string{ `\'`,`any`,`\'` }, CharLiteral1R0 },
	CharLiteral1R1:&Slot{"CharLiteral", 1, 1, []string{ `\'`,`any`,`\'` }, CharLiteral1R1 },
	CharLiteral1R2:&Slot{"CharLiteral", 1, 2, []string{ `\'`,`any`,`\'` }, CharLiteral1R2 },
	CharLiteral1R3:&Slot{"CharLiteral", 1, 3, []string{ `\'`,`any`,`\'` }, CharLiteral1R3 },
	EscapedChar0R0:&Slot{"EscapedChar", 0, 0, []string{ `"` }, EscapedChar0R0 },
	EscapedChar0R1:&Slot{"EscapedChar", 0, 1, []string{ `"` }, EscapedChar0R1 },
	EscapedChar1R0:&Slot{"EscapedChar", 1, 0, []string{ `n` }, EscapedChar1R0 },
	EscapedChar1R1:&Slot{"EscapedChar", 1, 1, []string{ `n` }, EscapedChar1R1 },
	EscapedChar2R0:&Slot{"EscapedChar", 2, 0, []string{ `r` }, EscapedChar2R0 },
	EscapedChar2R1:&Slot{"EscapedChar", 2, 1, []string{ `r` }, EscapedChar2R1 },
	EscapedChar3R0:&Slot{"EscapedChar", 3, 0, []string{ `t` }, EscapedChar3R0 },
	EscapedChar3R1:&Slot{"EscapedChar", 3, 1, []string{ `t` }, EscapedChar3R1 },
	EscapedChar4R0:&Slot{"EscapedChar", 4, 0, []string{ `\\` }, EscapedChar4R0 },
	EscapedChar4R1:&Slot{"EscapedChar", 4, 1, []string{ `\\` }, EscapedChar4R1 },
	EscapedChar5R0:&Slot{"EscapedChar", 5, 0, []string{ `\'` }, EscapedChar5R0 },
	EscapedChar5R1:&Slot{"EscapedChar", 5, 1, []string{ `\'` }, EscapedChar5R1 },
	GoGLL0R0:&Slot{"GoGLL", 0, 0, []string{ `SepE`,`Package`,`Sep`,`Rules`,`SepE` }, GoGLL0R0 },
	GoGLL0R1:&Slot{"GoGLL", 0, 1, []string{ `SepE`,`Package`,`Sep`,`Rules`,`SepE` }, GoGLL0R1 },
	GoGLL0R2:&Slot{"GoGLL", 0, 2, []string{ `SepE`,`Package`,`Sep`,`Rules`,`SepE` }, GoGLL0R2 },
	GoGLL0R3:&Slot{"GoGLL", 0, 3, []string{ `SepE`,`Package`,`Sep`,`Rules`,`SepE` }, GoGLL0R3 },
	GoGLL0R4:&Slot{"GoGLL", 0, 4, []string{ `SepE`,`Package`,`Sep`,`Rules`,`SepE` }, GoGLL0R4 },
	GoGLL0R5:&Slot{"GoGLL", 0, 5, []string{ `SepE`,`Package`,`Sep`,`Rules`,`SepE` }, GoGLL0R5 },
	Head0R0:&Slot{"Head", 0, 0, []string{ `*`,`NTID` }, Head0R0 },
	Head0R1:&Slot{"Head", 0, 1, []string{ `*`,`NTID` }, Head0R1 },
	Head0R2:&Slot{"Head", 0, 2, []string{ `*`,`NTID` }, Head0R2 },
	Head1R0:&Slot{"Head", 1, 0, []string{ `NTID` }, Head1R0 },
	Head1R1:&Slot{"Head", 1, 1, []string{ `NTID` }, Head1R1 },
	NTChar0R0:&Slot{"NTChar", 0, 0, []string{ `letter` }, NTChar0R0 },
	NTChar0R1:&Slot{"NTChar", 0, 1, []string{ `letter` }, NTChar0R1 },
	NTChar1R0:&Slot{"NTChar", 1, 0, []string{ `number` }, NTChar1R0 },
	NTChar1R1:&Slot{"NTChar", 1, 1, []string{ `number` }, NTChar1R1 },
	NTChar2R0:&Slot{"NTChar", 2, 0, []string{ `anyof(!#$%&*+-=@^_)` }, NTChar2R0 },
	NTChar2R1:&Slot{"NTChar", 2, 1, []string{ `anyof(!#$%&*+-=@^_)` }, NTChar2R1 },
	NTChars0R0:&Slot{"NTChars", 0, 0, []string{ `NTChar` }, NTChars0R0 },
	NTChars0R1:&Slot{"NTChars", 0, 1, []string{ `NTChar` }, NTChars0R1 },
	NTChars1R0:&Slot{"NTChars", 1, 0, []string{ `NTChar`,`NTChars` }, NTChars1R0 },
	NTChars1R1:&Slot{"NTChars", 1, 1, []string{ `NTChar`,`NTChars` }, NTChars1R1 },
	NTChars1R2:&Slot{"NTChars", 1, 2, []string{ `NTChar`,`NTChars` }, NTChars1R2 },
	NTID0R0:&Slot{"NTID", 0, 0, []string{ `letter` }, NTID0R0 },
	NTID0R1:&Slot{"NTID", 0, 1, []string{ `letter` }, NTID0R1 },
	NTID1R0:&Slot{"NTID", 1, 0, []string{ `letter`,`NTChars` }, NTID1R0 },
	NTID1R1:&Slot{"NTID", 1, 1, []string{ `letter`,`NTChars` }, NTID1R1 },
	NTID1R2:&Slot{"NTID", 1, 2, []string{ `letter`,`NTChars` }, NTID1R2 },
	NonTerminal0R0:&Slot{"NonTerminal", 0, 0, []string{ `NTID` }, NonTerminal0R0 },
	NonTerminal0R1:&Slot{"NonTerminal", 0, 1, []string{ `NTID` }, NonTerminal0R1 },
	Package0R0:&Slot{"Package", 0, 0, []string{ `p`,`a`,`c`,`k`,`a`,`g`,`e`,`Sep`,`String` }, Package0R0 },
	Package0R1:&Slot{"Package", 0, 1, []string{ `p`,`a`,`c`,`k`,`a`,`g`,`e`,`Sep`,`String` }, Package0R1 },
	Package0R2:&Slot{"Package", 0, 2, []string{ `p`,`a`,`c`,`k`,`a`,`g`,`e`,`Sep`,`String` }, Package0R2 },
	Package0R3:&Slot{"Package", 0, 3, []string{ `p`,`a`,`c`,`k`,`a`,`g`,`e`,`Sep`,`String` }, Package0R3 },
	Package0R4:&Slot{"Package", 0, 4, []string{ `p`,`a`,`c`,`k`,`a`,`g`,`e`,`Sep`,`String` }, Package0R4 },
	Package0R5:&Slot{"Package", 0, 5, []string{ `p`,`a`,`c`,`k`,`a`,`g`,`e`,`Sep`,`String` }, Package0R5 },
	Package0R6:&Slot{"Package", 0, 6, []string{ `p`,`a`,`c`,`k`,`a`,`g`,`e`,`Sep`,`String` }, Package0R6 },
	Package0R7:&Slot{"Package", 0, 7, []string{ `p`,`a`,`c`,`k`,`a`,`g`,`e`,`Sep`,`String` }, Package0R7 },
	Package0R8:&Slot{"Package", 0, 8, []string{ `p`,`a`,`c`,`k`,`a`,`g`,`e`,`Sep`,`String` }, Package0R8 },
	Package0R9:&Slot{"Package", 0, 9, []string{ `p`,`a`,`c`,`k`,`a`,`g`,`e`,`Sep`,`String` }, Package0R9 },
	Rule0R0:&Slot{"Rule", 0, 0, []string{ `Head`,`SepE`,`:`,`SepE`,`Alternates`,`SepE`,`;` }, Rule0R0 },
	Rule0R1:&Slot{"Rule", 0, 1, []string{ `Head`,`SepE`,`:`,`SepE`,`Alternates`,`SepE`,`;` }, Rule0R1 },
	Rule0R2:&Slot{"Rule", 0, 2, []string{ `Head`,`SepE`,`:`,`SepE`,`Alternates`,`SepE`,`;` }, Rule0R2 },
	Rule0R3:&Slot{"Rule", 0, 3, []string{ `Head`,`SepE`,`:`,`SepE`,`Alternates`,`SepE`,`;` }, Rule0R3 },
	Rule0R4:&Slot{"Rule", 0, 4, []string{ `Head`,`SepE`,`:`,`SepE`,`Alternates`,`SepE`,`;` }, Rule0R4 },
	Rule0R5:&Slot{"Rule", 0, 5, []string{ `Head`,`SepE`,`:`,`SepE`,`Alternates`,`SepE`,`;` }, Rule0R5 },
	Rule0R6:&Slot{"Rule", 0, 6, []string{ `Head`,`SepE`,`:`,`SepE`,`Alternates`,`SepE`,`;` }, Rule0R6 },
	Rule0R7:&Slot{"Rule", 0, 7, []string{ `Head`,`SepE`,`:`,`SepE`,`Alternates`,`SepE`,`;` }, Rule0R7 },
	Rules0R0:&Slot{"Rules", 0, 0, []string{ `Rule` }, Rules0R0 },
	Rules0R1:&Slot{"Rules", 0, 1, []string{ `Rule` }, Rules0R1 },
	Rules1R0:&Slot{"Rules", 1, 0, []string{ `Rule`,`SepE`,`Rules` }, Rules1R0 },
	Rules1R1:&Slot{"Rules", 1, 1, []string{ `Rule`,`SepE`,`Rules` }, Rules1R1 },
	Rules1R2:&Slot{"Rules", 1, 2, []string{ `Rule`,`SepE`,`Rules` }, Rules1R2 },
	Rules1R3:&Slot{"Rules", 1, 3, []string{ `Rule`,`SepE`,`Rules` }, Rules1R3 },
	Sep0R0:&Slot{"Sep", 0, 0, []string{ `SepChar` }, Sep0R0 },
	Sep0R1:&Slot{"Sep", 0, 1, []string{ `SepChar` }, Sep0R1 },
	Sep1R0:&Slot{"Sep", 1, 0, []string{ `SepChar`,`Sep` }, Sep1R0 },
	Sep1R1:&Slot{"Sep", 1, 1, []string{ `SepChar`,`Sep` }, Sep1R1 },
	Sep1R2:&Slot{"Sep", 1, 2, []string{ `SepChar`,`Sep` }, Sep1R2 },
	SepChar0R0:&Slot{"SepChar", 0, 0, []string{ `space` }, SepChar0R0 },
	SepChar0R1:&Slot{"SepChar", 0, 1, []string{ `space` }, SepChar0R1 },
	SepChar1R0:&Slot{"SepChar", 1, 0, []string{ `anyof(\n\r\t)` }, SepChar1R0 },
	SepChar1R1:&Slot{"SepChar", 1, 1, []string{ `anyof(\n\r\t)` }, SepChar1R1 },
	SepE0R0:&Slot{"SepE", 0, 0, []string{ `Sep` }, SepE0R0 },
	SepE0R1:&Slot{"SepE", 0, 1, []string{ `Sep` }, SepE0R1 },
	SepE1R0:&Slot{"SepE", 1, 0, []string{  }, SepE1R0 },
	String0R0:&Slot{"String", 0, 0, []string{ `\"`,`StringChars`,`\"` }, String0R0 },
	String0R1:&Slot{"String", 0, 1, []string{ `\"`,`StringChars`,`\"` }, String0R1 },
	String0R2:&Slot{"String", 0, 2, []string{ `\"`,`StringChars`,`\"` }, String0R2 },
	String0R3:&Slot{"String", 0, 3, []string{ `\"`,`StringChars`,`\"` }, String0R3 },
	StringChars0R0:&Slot{"StringChars", 0, 0, []string{ `not(\"\\)`,`StringChars` }, StringChars0R0 },
	StringChars0R1:&Slot{"StringChars", 0, 1, []string{ `not(\"\\)`,`StringChars` }, StringChars0R1 },
	StringChars0R2:&Slot{"StringChars", 0, 2, []string{ `not(\"\\)`,`StringChars` }, StringChars0R2 },
	StringChars1R0:&Slot{"StringChars", 1, 0, []string{ `\\`,`EscapedChar`,`StringChars` }, StringChars1R0 },
	StringChars1R1:&Slot{"StringChars", 1, 1, []string{ `\\`,`EscapedChar`,`StringChars` }, StringChars1R1 },
	StringChars1R2:&Slot{"StringChars", 1, 2, []string{ `\\`,`EscapedChar`,`StringChars` }, StringChars1R2 },
	StringChars1R3:&Slot{"StringChars", 1, 3, []string{ `\\`,`EscapedChar`,`StringChars` }, StringChars1R3 },
	StringChars2R0:&Slot{"StringChars", 2, 0, []string{  }, StringChars2R0 },
	Symbol0R0:&Slot{"Symbol", 0, 0, []string{ `NonTerminal` }, Symbol0R0 },
	Symbol0R1:&Slot{"Symbol", 0, 1, []string{ `NonTerminal` }, Symbol0R1 },
	Symbol1R0:&Slot{"Symbol", 1, 0, []string{ `Terminal` }, Symbol1R0 },
	Symbol1R1:&Slot{"Symbol", 1, 1, []string{ `Terminal` }, Symbol1R1 },
	Symbols0R0:&Slot{"Symbols", 0, 0, []string{ `Symbol`,`Sep`,`Symbols` }, Symbols0R0 },
	Symbols0R1:&Slot{"Symbols", 0, 1, []string{ `Symbol`,`Sep`,`Symbols` }, Symbols0R1 },
	Symbols0R2:&Slot{"Symbols", 0, 2, []string{ `Symbol`,`Sep`,`Symbols` }, Symbols0R2 },
	Symbols0R3:&Slot{"Symbols", 0, 3, []string{ `Symbol`,`Sep`,`Symbols` }, Symbols0R3 },
	Symbols1R0:&Slot{"Symbols", 1, 0, []string{ `Symbol` }, Symbols1R0 },
	Symbols1R1:&Slot{"Symbols", 1, 1, []string{ `Symbol` }, Symbols1R1 },
	Terminal0R0:&Slot{"Terminal", 0, 0, []string{ `a`,`n`,`y` }, Terminal0R0 },
	Terminal0R1:&Slot{"Terminal", 0, 1, []string{ `a`,`n`,`y` }, Terminal0R1 },
	Terminal0R2:&Slot{"Terminal", 0, 2, []string{ `a`,`n`,`y` }, Terminal0R2 },
	Terminal0R3:&Slot{"Terminal", 0, 3, []string{ `a`,`n`,`y` }, Terminal0R3 },
	Terminal1R0:&Slot{"Terminal", 1, 0, []string{ `a`,`n`,`y`,`o`,`f`,`Sep`,`String` }, Terminal1R0 },
	Terminal1R1:&Slot{"Terminal", 1, 1, []string{ `a`,`n`,`y`,`o`,`f`,`Sep`,`String` }, Terminal1R1 },
	Terminal1R2:&Slot{"Terminal", 1, 2, []string{ `a`,`n`,`y`,`o`,`f`,`Sep`,`String` }, Terminal1R2 },
	Terminal1R3:&Slot{"Terminal", 1, 3, []string{ `a`,`n`,`y`,`o`,`f`,`Sep`,`String` }, Terminal1R3 },
	Terminal1R4:&Slot{"Terminal", 1, 4, []string{ `a`,`n`,`y`,`o`,`f`,`Sep`,`String` }, Terminal1R4 },
	Terminal1R5:&Slot{"Terminal", 1, 5, []string{ `a`,`n`,`y`,`o`,`f`,`Sep`,`String` }, Terminal1R5 },
	Terminal1R6:&Slot{"Terminal", 1, 6, []string{ `a`,`n`,`y`,`o`,`f`,`Sep`,`String` }, Terminal1R6 },
	Terminal1R7:&Slot{"Terminal", 1, 7, []string{ `a`,`n`,`y`,`o`,`f`,`Sep`,`String` }, Terminal1R7 },
	Terminal2R0:&Slot{"Terminal", 2, 0, []string{ `l`,`e`,`t`,`t`,`e`,`r` }, Terminal2R0 },
	Terminal2R1:&Slot{"Terminal", 2, 1, []string{ `l`,`e`,`t`,`t`,`e`,`r` }, Terminal2R1 },
	Terminal2R2:&Slot{"Terminal", 2, 2, []string{ `l`,`e`,`t`,`t`,`e`,`r` }, Terminal2R2 },
	Terminal2R3:&Slot{"Terminal", 2, 3, []string{ `l`,`e`,`t`,`t`,`e`,`r` }, Terminal2R3 },
	Terminal2R4:&Slot{"Terminal", 2, 4, []string{ `l`,`e`,`t`,`t`,`e`,`r` }, Terminal2R4 },
	Terminal2R5:&Slot{"Terminal", 2, 5, []string{ `l`,`e`,`t`,`t`,`e`,`r` }, Terminal2R5 },
	Terminal2R6:&Slot{"Terminal", 2, 6, []string{ `l`,`e`,`t`,`t`,`e`,`r` }, Terminal2R6 },
	Terminal3R0:&Slot{"Terminal", 3, 0, []string{ `n`,`u`,`m`,`b`,`e`,`r` }, Terminal3R0 },
	Terminal3R1:&Slot{"Terminal", 3, 1, []string{ `n`,`u`,`m`,`b`,`e`,`r` }, Terminal3R1 },
	Terminal3R2:&Slot{"Terminal", 3, 2, []string{ `n`,`u`,`m`,`b`,`e`,`r` }, Terminal3R2 },
	Terminal3R3:&Slot{"Terminal", 3, 3, []string{ `n`,`u`,`m`,`b`,`e`,`r` }, Terminal3R3 },
	Terminal3R4:&Slot{"Terminal", 3, 4, []string{ `n`,`u`,`m`,`b`,`e`,`r` }, Terminal3R4 },
	Terminal3R5:&Slot{"Terminal", 3, 5, []string{ `n`,`u`,`m`,`b`,`e`,`r` }, Terminal3R5 },
	Terminal3R6:&Slot{"Terminal", 3, 6, []string{ `n`,`u`,`m`,`b`,`e`,`r` }, Terminal3R6 },
	Terminal4R0:&Slot{"Terminal", 4, 0, []string{ `s`,`p`,`a`,`c`,`e` }, Terminal4R0 },
	Terminal4R1:&Slot{"Terminal", 4, 1, []string{ `s`,`p`,`a`,`c`,`e` }, Terminal4R1 },
	Terminal4R2:&Slot{"Terminal", 4, 2, []string{ `s`,`p`,`a`,`c`,`e` }, Terminal4R2 },
	Terminal4R3:&Slot{"Terminal", 4, 3, []string{ `s`,`p`,`a`,`c`,`e` }, Terminal4R3 },
	Terminal4R4:&Slot{"Terminal", 4, 4, []string{ `s`,`p`,`a`,`c`,`e` }, Terminal4R4 },
	Terminal4R5:&Slot{"Terminal", 4, 5, []string{ `s`,`p`,`a`,`c`,`e` }, Terminal4R5 },
	Terminal5R0:&Slot{"Terminal", 5, 0, []string{ `u`,`p`,`c`,`a`,`s`,`e` }, Terminal5R0 },
	Terminal5R1:&Slot{"Terminal", 5, 1, []string{ `u`,`p`,`c`,`a`,`s`,`e` }, Terminal5R1 },
	Terminal5R2:&Slot{"Terminal", 5, 2, []string{ `u`,`p`,`c`,`a`,`s`,`e` }, Terminal5R2 },
	Terminal5R3:&Slot{"Terminal", 5, 3, []string{ `u`,`p`,`c`,`a`,`s`,`e` }, Terminal5R3 },
	Terminal5R4:&Slot{"Terminal", 5, 4, []string{ `u`,`p`,`c`,`a`,`s`,`e` }, Terminal5R4 },
	Terminal5R5:&Slot{"Terminal", 5, 5, []string{ `u`,`p`,`c`,`a`,`s`,`e` }, Terminal5R5 },
	Terminal5R6:&Slot{"Terminal", 5, 6, []string{ `u`,`p`,`c`,`a`,`s`,`e` }, Terminal5R6 },
	Terminal6R0:&Slot{"Terminal", 6, 0, []string{ `l`,`o`,`w`,`c`,`a`,`s`,`e` }, Terminal6R0 },
	Terminal6R1:&Slot{"Terminal", 6, 1, []string{ `l`,`o`,`w`,`c`,`a`,`s`,`e` }, Terminal6R1 },
	Terminal6R2:&Slot{"Terminal", 6, 2, []string{ `l`,`o`,`w`,`c`,`a`,`s`,`e` }, Terminal6R2 },
	Terminal6R3:&Slot{"Terminal", 6, 3, []string{ `l`,`o`,`w`,`c`,`a`,`s`,`e` }, Terminal6R3 },
	Terminal6R4:&Slot{"Terminal", 6, 4, []string{ `l`,`o`,`w`,`c`,`a`,`s`,`e` }, Terminal6R4 },
	Terminal6R5:&Slot{"Terminal", 6, 5, []string{ `l`,`o`,`w`,`c`,`a`,`s`,`e` }, Terminal6R5 },
	Terminal6R6:&Slot{"Terminal", 6, 6, []string{ `l`,`o`,`w`,`c`,`a`,`s`,`e` }, Terminal6R6 },
	Terminal6R7:&Slot{"Terminal", 6, 7, []string{ `l`,`o`,`w`,`c`,`a`,`s`,`e` }, Terminal6R7 },
	Terminal7R0:&Slot{"Terminal", 7, 0, []string{ `n`,`o`,`t`,`Sep`,`String` }, Terminal7R0 },
	Terminal7R1:&Slot{"Terminal", 7, 1, []string{ `n`,`o`,`t`,`Sep`,`String` }, Terminal7R1 },
	Terminal7R2:&Slot{"Terminal", 7, 2, []string{ `n`,`o`,`t`,`Sep`,`String` }, Terminal7R2 },
	Terminal7R3:&Slot{"Terminal", 7, 3, []string{ `n`,`o`,`t`,`Sep`,`String` }, Terminal7R3 },
	Terminal7R4:&Slot{"Terminal", 7, 4, []string{ `n`,`o`,`t`,`Sep`,`String` }, Terminal7R4 },
	Terminal7R5:&Slot{"Terminal", 7, 5, []string{ `n`,`o`,`t`,`Sep`,`String` }, Terminal7R5 },
	Terminal8R0:&Slot{"Terminal", 8, 0, []string{ `CharLiteral` }, Terminal8R0 },
	Terminal8R1:&Slot{"Terminal", 8, 1, []string{ `CharLiteral` }, Terminal8R1 },
	Terminal9R0:&Slot{"Terminal", 9, 0, []string{ `String` }, Terminal9R0 },
	Terminal9R1:&Slot{"Terminal", 9, 1, []string{ `String` }, Terminal9R1 },
}

var slotIndex = map[Index]Label { 
	Index{ "Alternate",0,0 }: Alternate0R0,
	Index{ "Alternate",0,1 }: Alternate0R1,
	Index{ "Alternate",1,0 }: Alternate1R0,
	Index{ "Alternate",1,1 }: Alternate1R1,
	Index{ "Alternate",1,2 }: Alternate1R2,
	Index{ "Alternate",1,3 }: Alternate1R3,
	Index{ "Alternate",1,4 }: Alternate1R4,
	Index{ "Alternate",1,5 }: Alternate1R5,
	Index{ "Alternates",0,0 }: Alternates0R0,
	Index{ "Alternates",0,1 }: Alternates0R1,
	Index{ "Alternates",1,0 }: Alternates1R0,
	Index{ "Alternates",1,1 }: Alternates1R1,
	Index{ "Alternates",1,2 }: Alternates1R2,
	Index{ "Alternates",1,3 }: Alternates1R3,
	Index{ "Alternates",1,4 }: Alternates1R4,
	Index{ "Alternates",1,5 }: Alternates1R5,
	Index{ "CharLiteral",0,0 }: CharLiteral0R0,
	Index{ "CharLiteral",0,1 }: CharLiteral0R1,
	Index{ "CharLiteral",0,2 }: CharLiteral0R2,
	Index{ "CharLiteral",0,3 }: CharLiteral0R3,
	Index{ "CharLiteral",0,4 }: CharLiteral0R4,
	Index{ "CharLiteral",1,0 }: CharLiteral1R0,
	Index{ "CharLiteral",1,1 }: CharLiteral1R1,
	Index{ "CharLiteral",1,2 }: CharLiteral1R2,
	Index{ "CharLiteral",1,3 }: CharLiteral1R3,
	Index{ "EscapedChar",0,0 }: EscapedChar0R0,
	Index{ "EscapedChar",0,1 }: EscapedChar0R1,
	Index{ "EscapedChar",1,0 }: EscapedChar1R0,
	Index{ "EscapedChar",1,1 }: EscapedChar1R1,
	Index{ "EscapedChar",2,0 }: EscapedChar2R0,
	Index{ "EscapedChar",2,1 }: EscapedChar2R1,
	Index{ "EscapedChar",3,0 }: EscapedChar3R0,
	Index{ "EscapedChar",3,1 }: EscapedChar3R1,
	Index{ "EscapedChar",4,0 }: EscapedChar4R0,
	Index{ "EscapedChar",4,1 }: EscapedChar4R1,
	Index{ "EscapedChar",5,0 }: EscapedChar5R0,
	Index{ "EscapedChar",5,1 }: EscapedChar5R1,
	Index{ "GoGLL",0,0 }: GoGLL0R0,
	Index{ "GoGLL",0,1 }: GoGLL0R1,
	Index{ "GoGLL",0,2 }: GoGLL0R2,
	Index{ "GoGLL",0,3 }: GoGLL0R3,
	Index{ "GoGLL",0,4 }: GoGLL0R4,
	Index{ "GoGLL",0,5 }: GoGLL0R5,
	Index{ "Head",0,0 }: Head0R0,
	Index{ "Head",0,1 }: Head0R1,
	Index{ "Head",0,2 }: Head0R2,
	Index{ "Head",1,0 }: Head1R0,
	Index{ "Head",1,1 }: Head1R1,
	Index{ "NTChar",0,0 }: NTChar0R0,
	Index{ "NTChar",0,1 }: NTChar0R1,
	Index{ "NTChar",1,0 }: NTChar1R0,
	Index{ "NTChar",1,1 }: NTChar1R1,
	Index{ "NTChar",2,0 }: NTChar2R0,
	Index{ "NTChar",2,1 }: NTChar2R1,
	Index{ "NTChars",0,0 }: NTChars0R0,
	Index{ "NTChars",0,1 }: NTChars0R1,
	Index{ "NTChars",1,0 }: NTChars1R0,
	Index{ "NTChars",1,1 }: NTChars1R1,
	Index{ "NTChars",1,2 }: NTChars1R2,
	Index{ "NTID",0,0 }: NTID0R0,
	Index{ "NTID",0,1 }: NTID0R1,
	Index{ "NTID",1,0 }: NTID1R0,
	Index{ "NTID",1,1 }: NTID1R1,
	Index{ "NTID",1,2 }: NTID1R2,
	Index{ "NonTerminal",0,0 }: NonTerminal0R0,
	Index{ "NonTerminal",0,1 }: NonTerminal0R1,
	Index{ "Package",0,0 }: Package0R0,
	Index{ "Package",0,1 }: Package0R1,
	Index{ "Package",0,2 }: Package0R2,
	Index{ "Package",0,3 }: Package0R3,
	Index{ "Package",0,4 }: Package0R4,
	Index{ "Package",0,5 }: Package0R5,
	Index{ "Package",0,6 }: Package0R6,
	Index{ "Package",0,7 }: Package0R7,
	Index{ "Package",0,8 }: Package0R8,
	Index{ "Package",0,9 }: Package0R9,
	Index{ "Rule",0,0 }: Rule0R0,
	Index{ "Rule",0,1 }: Rule0R1,
	Index{ "Rule",0,2 }: Rule0R2,
	Index{ "Rule",0,3 }: Rule0R3,
	Index{ "Rule",0,4 }: Rule0R4,
	Index{ "Rule",0,5 }: Rule0R5,
	Index{ "Rule",0,6 }: Rule0R6,
	Index{ "Rule",0,7 }: Rule0R7,
	Index{ "Rules",0,0 }: Rules0R0,
	Index{ "Rules",0,1 }: Rules0R1,
	Index{ "Rules",1,0 }: Rules1R0,
	Index{ "Rules",1,1 }: Rules1R1,
	Index{ "Rules",1,2 }: Rules1R2,
	Index{ "Rules",1,3 }: Rules1R3,
	Index{ "Sep",0,0 }: Sep0R0,
	Index{ "Sep",0,1 }: Sep0R1,
	Index{ "Sep",1,0 }: Sep1R0,
	Index{ "Sep",1,1 }: Sep1R1,
	Index{ "Sep",1,2 }: Sep1R2,
	Index{ "SepChar",0,0 }: SepChar0R0,
	Index{ "SepChar",0,1 }: SepChar0R1,
	Index{ "SepChar",1,0 }: SepChar1R0,
	Index{ "SepChar",1,1 }: SepChar1R1,
	Index{ "SepE",0,0 }: SepE0R0,
	Index{ "SepE",0,1 }: SepE0R1,
	Index{ "SepE",1,0 }: SepE1R0,
	Index{ "String",0,0 }: String0R0,
	Index{ "String",0,1 }: String0R1,
	Index{ "String",0,2 }: String0R2,
	Index{ "String",0,3 }: String0R3,
	Index{ "StringChars",0,0 }: StringChars0R0,
	Index{ "StringChars",0,1 }: StringChars0R1,
	Index{ "StringChars",0,2 }: StringChars0R2,
	Index{ "StringChars",1,0 }: StringChars1R0,
	Index{ "StringChars",1,1 }: StringChars1R1,
	Index{ "StringChars",1,2 }: StringChars1R2,
	Index{ "StringChars",1,3 }: StringChars1R3,
	Index{ "StringChars",2,0 }: StringChars2R0,
	Index{ "Symbol",0,0 }: Symbol0R0,
	Index{ "Symbol",0,1 }: Symbol0R1,
	Index{ "Symbol",1,0 }: Symbol1R0,
	Index{ "Symbol",1,1 }: Symbol1R1,
	Index{ "Symbols",0,0 }: Symbols0R0,
	Index{ "Symbols",0,1 }: Symbols0R1,
	Index{ "Symbols",0,2 }: Symbols0R2,
	Index{ "Symbols",0,3 }: Symbols0R3,
	Index{ "Symbols",1,0 }: Symbols1R0,
	Index{ "Symbols",1,1 }: Symbols1R1,
	Index{ "Terminal",0,0 }: Terminal0R0,
	Index{ "Terminal",0,1 }: Terminal0R1,
	Index{ "Terminal",0,2 }: Terminal0R2,
	Index{ "Terminal",0,3 }: Terminal0R3,
	Index{ "Terminal",1,0 }: Terminal1R0,
	Index{ "Terminal",1,1 }: Terminal1R1,
	Index{ "Terminal",1,2 }: Terminal1R2,
	Index{ "Terminal",1,3 }: Terminal1R3,
	Index{ "Terminal",1,4 }: Terminal1R4,
	Index{ "Terminal",1,5 }: Terminal1R5,
	Index{ "Terminal",1,6 }: Terminal1R6,
	Index{ "Terminal",1,7 }: Terminal1R7,
	Index{ "Terminal",2,0 }: Terminal2R0,
	Index{ "Terminal",2,1 }: Terminal2R1,
	Index{ "Terminal",2,2 }: Terminal2R2,
	Index{ "Terminal",2,3 }: Terminal2R3,
	Index{ "Terminal",2,4 }: Terminal2R4,
	Index{ "Terminal",2,5 }: Terminal2R5,
	Index{ "Terminal",2,6 }: Terminal2R6,
	Index{ "Terminal",3,0 }: Terminal3R0,
	Index{ "Terminal",3,1 }: Terminal3R1,
	Index{ "Terminal",3,2 }: Terminal3R2,
	Index{ "Terminal",3,3 }: Terminal3R3,
	Index{ "Terminal",3,4 }: Terminal3R4,
	Index{ "Terminal",3,5 }: Terminal3R5,
	Index{ "Terminal",3,6 }: Terminal3R6,
	Index{ "Terminal",4,0 }: Terminal4R0,
	Index{ "Terminal",4,1 }: Terminal4R1,
	Index{ "Terminal",4,2 }: Terminal4R2,
	Index{ "Terminal",4,3 }: Terminal4R3,
	Index{ "Terminal",4,4 }: Terminal4R4,
	Index{ "Terminal",4,5 }: Terminal4R5,
	Index{ "Terminal",5,0 }: Terminal5R0,
	Index{ "Terminal",5,1 }: Terminal5R1,
	Index{ "Terminal",5,2 }: Terminal5R2,
	Index{ "Terminal",5,3 }: Terminal5R3,
	Index{ "Terminal",5,4 }: Terminal5R4,
	Index{ "Terminal",5,5 }: Terminal5R5,
	Index{ "Terminal",5,6 }: Terminal5R6,
	Index{ "Terminal",6,0 }: Terminal6R0,
	Index{ "Terminal",6,1 }: Terminal6R1,
	Index{ "Terminal",6,2 }: Terminal6R2,
	Index{ "Terminal",6,3 }: Terminal6R3,
	Index{ "Terminal",6,4 }: Terminal6R4,
	Index{ "Terminal",6,5 }: Terminal6R5,
	Index{ "Terminal",6,6 }: Terminal6R6,
	Index{ "Terminal",6,7 }: Terminal6R7,
	Index{ "Terminal",7,0 }: Terminal7R0,
	Index{ "Terminal",7,1 }: Terminal7R1,
	Index{ "Terminal",7,2 }: Terminal7R2,
	Index{ "Terminal",7,3 }: Terminal7R3,
	Index{ "Terminal",7,4 }: Terminal7R4,
	Index{ "Terminal",7,5 }: Terminal7R5,
	Index{ "Terminal",8,0 }: Terminal8R0,
	Index{ "Terminal",8,1 }: Terminal8R1,
	Index{ "Terminal",9,0 }: Terminal9R0,
	Index{ "Terminal",9,1 }: Terminal9R1,
}

var alternates = map[string][]Label{ 
	"GoGLL":[]Label{ GoGLL0R0 },
	"Package":[]Label{ Package0R0 },
	"Rules":[]Label{ Rules0R0,Rules1R0 },
	"Rule":[]Label{ Rule0R0 },
	"Head":[]Label{ Head0R0,Head1R0 },
	"Alternates":[]Label{ Alternates0R0,Alternates1R0 },
	"Alternate":[]Label{ Alternate0R0,Alternate1R0 },
	"Symbols":[]Label{ Symbols0R0,Symbols1R0 },
	"Symbol":[]Label{ Symbol0R0,Symbol1R0 },
	"NonTerminal":[]Label{ NonTerminal0R0 },
	"NTID":[]Label{ NTID0R0,NTID1R0 },
	"NTChars":[]Label{ NTChars0R0,NTChars1R0 },
	"NTChar":[]Label{ NTChar0R0,NTChar1R0,NTChar2R0 },
	"Terminal":[]Label{ Terminal0R0,Terminal1R0,Terminal2R0,Terminal3R0,Terminal4R0,Terminal5R0,Terminal6R0,Terminal7R0,Terminal8R0,Terminal9R0 },
	"CharLiteral":[]Label{ CharLiteral0R0,CharLiteral1R0 },
	"EscapedChar":[]Label{ EscapedChar0R0,EscapedChar1R0,EscapedChar2R0,EscapedChar3R0,EscapedChar4R0,EscapedChar5R0 },
	"String":[]Label{ String0R0 },
	"StringChars":[]Label{ StringChars0R0,StringChars1R0,StringChars2R0 },
	"SepE":[]Label{ SepE0R0,SepE1R0 },
	"Sep":[]Label{ Sep0R0,Sep1R0 },
	"SepChar":[]Label{ SepChar0R0,SepChar1R0 },
}

