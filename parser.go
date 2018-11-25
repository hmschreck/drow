package drow

import (
	"regexp"
	"strconv"
)

var diceNoRegex = regexp.MustCompile("^([0-9]*)")
var diceTypeRegex = regexp.MustCompile("d([0-9]*)")
var highLowRegex = regexp.MustCompile("([hl])([0-9]*)")
var modifierRegex = regexp.MustCompile("([+-])([0-9]*)")

type ParseResult struct {
	EndResult int
	UsedRolls []int
	Rolls []int
	TotalBeforeModifier int
	Modifier int
	HighLow int
	NumberOfRolls int
	Rolled Die
	ParseString string
	Parsed bool
	Error bool
}



func (parsed *ParseResult) Process() {
	if parsed.Parsed || parsed.Error{
		return
	}
	result, err := strconv.Atoi(parsed.ParseString)
	if err == nil {
		parsed.EndResult = result
		parsed.Parsed = true
		return
	}
	parsed.NumberOfRolls = 1
	diceToRoll := diceNoRegex.FindStringSubmatch(parsed.ParseString)
	if diceToRoll[0] != ""{
		parsed.NumberOfRolls, _ = strconv.Atoi(diceToRoll[1])
	}
	sides := diceTypeRegex.FindStringSubmatch(parsed.ParseString)
	sidesInt, _ := strconv.Atoi(sides[1])
	parsed.Rolled = *NewDie(sidesInt)
	highLow := highLowRegex.FindStringSubmatch(parsed.ParseString)
	if highLow != nil {
		if highLow[2] == "0" {
			parsed.Error = true
			return
		}
		take, _ := strconv.Atoi(highLow[2])
		if highLow[1] == "h" {
			parsed.HighLow = take
		} else {

			parsed.HighLow = 0 - take
		}
	} else {
		parsed.HighLow = 0
	}
	modifier := modifierRegex.FindStringSubmatch(parsed.ParseString)
	if modifier != nil {
		modifierInt, _ := strconv.Atoi(modifier[2])
		if modifier[1] == "+" {
			parsed.Modifier = modifierInt
		} else {
			parsed.Modifier = 0 - modifierInt
		}
	} else {
		parsed.Modifier = 0
	}
	if parsed.HighLow == 0 {
		parsed.Rolls = parsed.Rolled.RollMultiple(parsed.NumberOfRolls)
		parsed.UsedRolls = parsed.Rolls
	} else {
		parsed.UsedRolls, parsed.Rolls, _ = parsed.Rolled.HighLow(parsed.NumberOfRolls, parsed.HighLow)
	}
	parsed.TotalBeforeModifier = Sum(parsed.UsedRolls)
	parsed.EndResult = parsed.TotalBeforeModifier + parsed.Modifier
	parsed.Parsed = true
	}

func ParseDieCode(parseString string) *ParseResult {
	parsed := new(ParseResult)
	parsed.ParseString = parseString
	parsed.Process()
	return parsed
}
