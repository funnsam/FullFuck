package main

/*
BSD 2-Clause License

Copyright (c) 2022, funnsam
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

import (
	"fmt"
	"os"
)

type Token struct {
	Lable       uint64 // Lable is for the register for the loop count
	ID          uint32 // ID is used for identifying tokens
	LoopID      uint64 // LoopID is uniqe for each loop
	URCLSpecial string // URCLSpecial is for special things
}

var Uerr error
var InputFile []byte
var LoopID uint64
var LoopLayer uint64
var TokenList []Token
var ParsingSpecial uint8
var ParsingSpecialBuffer []uint8
var LoopLoopsTimes = []int16{0}
var t []uint8
var OutputFile []byte
var NoHLT bool

var FullFuckToURCLTable = []string{
	"ADD R1 R1 %d\n",
	"SUB R1 R1 %d\n",
	"OUT %v R1\n",
	"IN R1 %v\n",
	"BRZ .loop%d_e R1\nMOV R%d %d\n.loop%d\n",
	"DEC R%d R%d\nBNZ .loop%d R%d\n.loop%d_e\n",
	"POP R1\n",
	"PSH R1\n",
	"MOV R1 R%d\n",
}

// ParseSpecial used list: 1, 2, 3, 8, 9, 10

func SpecialParsing(element byte) {
	if ParsingSpecial == 1 || ParsingSpecial == 2 {
		ParsingSpecialBuffer = append(ParsingSpecialBuffer, element)
		ParsingSpecial++
	} else if ParsingSpecial == 8 {
		if element == '%' {
			ParsingSpecial = 9
		} else {
			ParsingSpecialBuffer = append(ParsingSpecialBuffer, element)
		}
	} else if ParsingSpecial == 10 {
		switch element {
		case '<', 'i', 'I':
			TokenList = append(TokenList, Token{LoopLayer, 6, LoopID, ""})
		case '>', 'o', 'O':
			TokenList = append(TokenList, Token{LoopLayer, 7, LoopID, ""})
		case '[':
			OpenBracket()
		}
		ParsingSpecial = 0
	}
}

func Parse(element byte) {
	if ParsingSpecial == 1 || ParsingSpecial == 2 || ParsingSpecial == 8 || ParsingSpecial == 10 {
		SpecialParsing(element)
		return
	}
	switch element {
	case '+': // 0
		TokenList = append(TokenList, Token{LoopLayer, 0, LoopID, ""})
	case '-': // 1
		TokenList = append(TokenList, Token{LoopLayer, 1, LoopID, ""})
	case '>': // 2
		OutputToken()
	case '<': // 3
		InputToken()
	case '[': // 4
		OpenBracket()
	case ']': // 5
		TokenList = append(TokenList, Token{LoopLayer, 5, LoopID, ""})
		LoopLayer--
	case '?': // 6, 7
		ParsingSpecial = 10
	case '$': // 8
		TokenList = append(TokenList, Token{LoopLayer, 8, LoopID, ""})
	case '0':
		ParsingSpecial = 1
	case '%':
		ParsingSpecial = 8
	}
}

func CompileToURCL(OTokenList []OToken) []byte {
	var ReturnValue []byte
	for _, element := range OTokenList {
		var resultAppend string
		switch element.Token.ID & 0xFF {
		case 0:
			resultAppend = fmt.Sprintf(FullFuckToURCLTable[0], element.Repeated)
		case 1:
			resultAppend = fmt.Sprintf(FullFuckToURCLTable[1], element.Repeated)
		case 2:
			if element.Token.ID == 0xFFFFFF02 {
				resultAppend = fmt.Sprintf(FullFuckToURCLTable[2], "%"+element.Token.URCLSpecial)
			} else {
				resultAppend = fmt.Sprintf(FullFuckToURCLTable[2], element.Token.ID>>8)
			}
		case 3:
			if element.Token.ID == 0xFFFFFF03 {
				resultAppend = fmt.Sprintf(FullFuckToURCLTable[3], "%"+element.Token.URCLSpecial)
			} else {
				resultAppend = fmt.Sprintf(FullFuckToURCLTable[3], element.Token.ID>>8)
			}
		case 4:
			if LoopLoopsTimes[element.Token.LoopID] == -1 {
				resultAppend = fmt.Sprintf(".loop%d\n", element.Token.LoopID)
			} else if LoopLoopsTimes[element.Token.LoopID] == -2 {
				resultAppend = fmt.Sprintf("BRZ .loop%d_e R1\nMOV R%d R1\n.loop%d\n", element.Token.LoopID, element.Token.Lable+1, element.Token.LoopID)
			} else {
				resultAppend = fmt.Sprintf(FullFuckToURCLTable[4], element.Token.LoopID, element.Token.Lable+1, LoopLoopsTimes[element.Token.LoopID], element.Token.LoopID)
			}
		case 5:
			if LoopLoopsTimes[element.Token.LoopID] == -1 {
				resultAppend = fmt.Sprintf("JMP .loop%d\n", element.Token.LoopID)
			} else {
				resultAppend = fmt.Sprintf(FullFuckToURCLTable[5], element.Token.Lable+1, element.Token.Lable+1, element.Token.LoopID, element.Token.Lable+1, element.Token.LoopID)
			}
		case 8:
			resultAppend = fmt.Sprintf(FullFuckToURCLTable[8], element.Token.Lable+1)
		default:
			resultAppend = FullFuckToURCLTable[element.Token.ID]
		}
		ReturnValue = append(ReturnValue, []byte(resultAppend)...)
	}
	return ReturnValue
}

func main() {
	if len(os.Args) <= 2 {
		fmt.Print("\x1b[1;31m>:(\t\tSee the docs! -> https://github.com/funnsam/FullFuck/wiki/Basic-syntax\n\x1b[1;0m")
		os.Exit(-1)
	}

	for _, element := range os.Args {
		if element == "--no-hlt" {
			NoHLT = true
		}
	}

	InputFile, Uerr = os.ReadFile(os.Args[1])
	checkUErr()

	for _, element := range InputFile {
		Parse(element)
	}

	OutputFile = append(OutputFile, []byte("MINSTACK 0xEF\n")...)
	OutputFile = append(OutputFile, CompileToURCL(Optimize(TokenList))...)
	if !NoHLT {
		OutputFile = append(OutputFile, []byte("HLT")...)
	}
	os.WriteFile(os.Args[2], []byte(OutputFile), 0664)
}

func checkUErr() {
	if Uerr != nil {
		panic(Uerr)
	}
}
