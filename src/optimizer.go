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
