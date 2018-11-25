package drow

import (
	"fmt"
	"math/rand"
	"testing"
)

const TestRuns = 5
const MaxDice = 4
const ModRange = 6 // Add 1 to this to allow for zeros! <3

func TestParserXdY(t *testing.T) {
	for _, die := range dice {
		for i := 0; i < TestRuns; i++ {
			rollTimes := rand.Intn(MaxDice) + 1
			result := ParseString(fmt.Sprintf("%dd%d", rollTimes, die.Sides))
			fmt.Println(result)
		}
	}
}

func TestParserdY(t *testing.T) {
	for _, die := range dice {
		for i := 0; i < TestRuns; i++ {
			result := ParseString(fmt.Sprintf("d%d", die.Sides))
			if result.EndResult == 0 {
				t.Errorf("Did not parse correctly")
			}
		}
	}
}

func TestParserXdYPlusZ(t *testing.T) {
	for _, die := range dice {
		for i := 0; i < TestRuns; i++ {
			rollTimes := rand.Intn(MaxDice) + 1
			modifier := rand.Intn(ModRange)
			result := ParseString(fmt.Sprintf("%dd%d+%d", rollTimes, die.Sides, modifier))
			fmt.Println(result)
			}
	}
}

func TestParserXdYMinusZ(t *testing.T) {
	for _, die := range dice {
		for i := 0; i < TestRuns; i++ {
			rollTimes := rand.Intn(MaxDice) + 1
			modifier := rand.Intn(ModRange)
			result := ParseString(fmt.Sprintf("%dd%d-%d", rollTimes, die.Sides, modifier))
			fmt.Println(result)
		}
	}
}

func TestParserXdYhZ(t *testing.T) {
	for _, die := range dice {
		for i := 0; i < TestRuns; i++ {
			rollTimes := rand.Intn(MaxDice)+1
			highest := rand.Intn(rollTimes)+1
			result := ParseString(fmt.Sprintf("%dd%dh%d", rollTimes, die.Sides, highest))
			fmt.Println(result)
			}
	}
}

func TestParseXdYh0(t *testing.T) {
	for _, die := range dice {
		rollTimes := 7
		highest := 0
		result := ParseString(fmt.Sprintf("%dd%dh%d", rollTimes, die.Sides, highest))
		if result.Error == false {
			t.Errorf("Failed on highest=0")
		}
		}
}

func TestParserXdYlZ(t *testing.T) {
	for _, die := range dice {
		for i := 0; i < TestRuns; i++ {
			rollTimes := rand.Intn(MaxDice)+1
			highest := rand.Intn(rollTimes)+1
			result := ParseString(fmt.Sprintf("%dd%dl%d", rollTimes, die.Sides, highest))
			fmt.Println(result)
		}
	}
}

func TestParseXdYl0(t *testing.T) {
	for _, die := range dice {
		rollTimes := 7
		highest := 0
		result := ParseString(fmt.Sprintf("%dd%dl%d", rollTimes, die.Sides, highest))
		if result.Error == false {
			t.Errorf("Failed on highest=0")
		}
	}
}

