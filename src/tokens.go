package main

import "encoding/hex"

func OutputToken() {
	if ParsingSpecial == 3 {
		t, Uerr = hex.DecodeString(string(ParsingSpecialBuffer))
		checkUErr()

		TokenList = append(TokenList, Token{LoopLayer, 2 | uint32(t[0])<<8, LoopID, ""})

		ParsingSpecial = 0
		ParsingSpecialBuffer = make([]uint8, 0, 2)
	} else if ParsingSpecial == 9 {
		TokenList = append(TokenList, Token{LoopLayer, 0xFFFFFF02, LoopID, string(ParsingSpecialBuffer)})
		ParsingSpecialBuffer = make([]uint8, 0, 2)
		ParsingSpecial = 0
	} else {
		TokenList = append(TokenList, Token{LoopLayer, 258, LoopID, ""})
	}
}

func InputToken() {
	if ParsingSpecial == 3 {
		t, Uerr = hex.DecodeString(string(ParsingSpecialBuffer))
		checkUErr()

		TokenList = append(TokenList, Token{LoopLayer, 3 | uint32(t[0])<<8, LoopID, ""})

		ParsingSpecial = 0
		ParsingSpecialBuffer = make([]uint8, 0, 2)
	} else if ParsingSpecial == 9 {
		TokenList = append(TokenList, Token{LoopLayer, 0xFFFFFF03, LoopID, string(ParsingSpecialBuffer)})
		ParsingSpecialBuffer = make([]uint8, 0, 2)
		ParsingSpecial = 0
	} else {
		TokenList = append(TokenList, Token{LoopLayer, 259, LoopID, ""})
	}
}

func OpenBracket() {
	LoopID++
	LoopLayer++

	TokenList = append(TokenList, Token{LoopLayer, 4, LoopID, ""})
	t, Uerr = hex.DecodeString(string(ParsingSpecialBuffer))
	checkUErr()

	if ParsingSpecial == 3 {
		LoopLoopsTimes = append(LoopLoopsTimes, int16(t[0]))
	} else if ParsingSpecial == 10 {
		LoopLoopsTimes = append(LoopLoopsTimes, -2)
	} else {
		LoopLoopsTimes = append(LoopLoopsTimes, -1)
	}
	ParsingSpecial = 0
	ParsingSpecialBuffer = make([]uint8, 0, 2)

	LoopStack.Push(LoopID)
}
