package main

type OToken struct {
	Token    Token
	Repeated int
}

func Optimize(TokenList []Token) []OToken {
	var StackedAdds int
	var StackedMinus int
	var TempToken Token
	var TempOToken []OToken
	for _, element := range TokenList {
		if element.ID == 0 {
			TempToken = element
			StackedAdds++
			continue
		} else if element.ID == 1 {
			TempToken = element
			StackedMinus++
			continue
		}
		if StackedAdds != 0 && element.ID != 0 {
			TempOToken = append(TempOToken, OToken{TempToken, StackedAdds})
			StackedAdds = 0
		} else if StackedMinus != 0 && element.ID != 1 {
			TempOToken = append(TempOToken, OToken{TempToken, StackedMinus})
			StackedMinus = 0
		}
		TempOToken = append(TempOToken, OToken{element, 1})
	}
	return TempOToken
}

func UnrollSimpleLoops(TokenList []Token) []Token {
	var TempTokenList []Token

	var i int
	for i < len(TokenList) {
		if TokenList[i].ID == 4 && TokenList[i+1].ID == 0 && TokenList[i+2].ID == 5 {
			TempTokenList = append(TempTokenList, Token{0, 9, 0, ""})
			for j := 0; j < int(LoopLoopsTimes[TokenList[i].LoopID]); j++ {
				TempTokenList = append(TempTokenList, TokenList[i+1])
			}
			i += 3
		} else if TokenList[i].ID == 4 && TokenList[i+1].ID == 1 && TokenList[i+2].ID == 5 {
			for j := 0; j < int(LoopLoopsTimes[TokenList[i].LoopID]); j++ {
				TempTokenList = append(TempTokenList, TokenList[i+1])
			}
			i += 3
		} else {
			TempTokenList = append(TempTokenList, TokenList[i])
			i++
		}
	}

	return TempTokenList
}
