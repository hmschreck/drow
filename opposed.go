package drow

import "strings"

type Opposed struct {
	ParseString string
	LeftSide string
	RightSide string
	LeftResult ParseResult
	RightResult ParseResult
}

func ParseOpposed(parseString string) *Opposed {
	opposed := new(Opposed)
	opposed.ParseString = parseString
	split := strings.Split(opposed.ParseString, "opp")
	opposed.LeftSide = split[0]
	opposed.RightSide = split[1]
	opposed.LeftResult = *ParseDieCode(opposed.LeftSide)
	opposed.RightResult = *ParseDieCode(opposed.RightSide)
	return opposed
}

func (opposed *Opposed) Evaluate() int8 {
	if opposed.LeftResult.EndResult > opposed.RightResult.EndResult {
		return -1
	} else if opposed.LeftResult.EndResult == opposed.RightResult.EndResult {
		return 0
	} else {
		return 1
	}
}
