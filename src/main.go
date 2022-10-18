package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

type Token struct {
	Lable  uint64 // Lable is for the register for the loop count
	ID     uint8  // ID is used for identifying tokens
	LoopID uint64 // LoopID is uniqe for each loop
}

var Uerr error
var InputFile []byte
var LoopID uint64
var LoopLayer uint64
var TokenList []Token
var ParsingHex uint8
var ParsingHexChars []uint8
var LoopLoopsTimes = []uint8{0}
var t []uint8
var OutputFile []byte

var FullFuckToURCLTable = []string{
	"INC R1 R1\n",
	"DEC R1 R1\n",
	"OUT %TEXT R1\n",
	"BRZ .loop%d_e R1\nMOV R%d %d\n.loop%d\n",
	"DEC R%d R%d\nBNZ .loop%d R%d\n.loop%d_e\n",
}

func main() {
	if len(os.Args) <= 2 {
		fmt.Print("\x1b[1;31m>:(\n\x1b[1;0m")
		os.Exit(-1)
	}

	InputFile, Uerr = os.ReadFile(os.Args[1])
	checkUErr()

	for _, element := range InputFile {
		if ParsingHex == 1 || ParsingHex == 2 {
			ParsingHexChars = append(ParsingHexChars, element)
			ParsingHex++
			continue
		}
		switch element {
		case '+':
			TokenList = append(TokenList, Token{LoopLayer, 0, LoopID})
		case '-':
			TokenList = append(TokenList, Token{LoopLayer, 1, LoopID})
		case '>':
			TokenList = append(TokenList, Token{LoopLayer, 2, LoopID})
		case '0':
			ParsingHex = 1
		case '[':
			if ParsingHex == 3 {
				LoopID++
				LoopLayer++

				TokenList = append(TokenList, Token{LoopLayer, 3, LoopID})
				t, Uerr = hex.DecodeString(string(ParsingHexChars))
				checkUErr()
				LoopLoopsTimes = append(LoopLoopsTimes, t[0])

				ParsingHex = 0
				ParsingHexChars = make([]uint8, 0, 2)
			}
		case ']':
			TokenList = append(TokenList, Token{LoopLayer, 4, LoopID})
			LoopLayer--
		}
	}

	// Compile to URCL now
	for _, element := range TokenList {
		var resultAppend string
		switch element.ID {
		case 3:
			resultAppend = fmt.Sprintf(FullFuckToURCLTable[3], element.LoopID, element.Lable+1, LoopLoopsTimes[element.LoopID], element.LoopID)
		case 4:
			resultAppend = fmt.Sprintf(FullFuckToURCLTable[4], element.Lable+1, element.Lable+1, element.LoopID, element.Lable+1, element.LoopID)
		default:
			resultAppend = FullFuckToURCLTable[element.ID]
		}
		OutputFile = append(OutputFile, []byte(resultAppend)...)
	}
	OutputFile = append(OutputFile, []byte("HLT")...)
	os.WriteFile(os.Args[2], []byte(OutputFile), 0664)
}

func checkUErr() {
	if Uerr != nil {
		panic(Uerr)
	}
}
