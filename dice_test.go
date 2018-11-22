package drow

import (
	"fmt"
	"math"
	"testing"
)

var dice = []*Die{D4, D6, D8, D10, D12, D20, D100}
const TestRolls = 100000
const TestRollMaxVariance = 0.01

// A basic check for appropriate randomness, based on a large distribution of rolls.  1% variance is... still pretty reasonable
func TestDie_Roll(t *testing.T) {
	for _, die := range dice {
		expectedSum := die.ExpectedValue * float64(TestRolls)
		sum := 0
		for i := 0; i < TestRolls; i++ {
			sum += die.Roll()
		}
		var variance  float64 = float64(math.Abs(expectedSum - float64(sum))/expectedSum)
		if variance > TestRollMaxVariance {
			t.Errorf(fmt.Sprintf("Die %s failed its test with a variance of %v", die.Name, variance))
		}
	}
}

const TestRollTimes = 10
func TestDie_RollMultiple(t *testing.T) {
	for _, die := range dice {
		rolls := die.RollMultiple(TestRollTimes)
		if len(rolls) != TestRollTimes {
			t.Errorf(fmt.Sprintf("Die %s return %d results, expected %d", die.Name, len(rolls), TestRollTimes))
		}
	}
}

func TestDie_Highest(t *testing.T) {
	for _, die := range dice {
		timesToRoll := 4
		take := 1
		results, rolls, err := die.Highest(timesToRoll, take)
		if err != nil {
			t.Errorf("Encountered an error.")
		}
		if results[0] != rolls[len(rolls)-1] {
			t.Errorf("Did not receive highest value.  Received %v, all results %v", results, rolls)
		}
		// test for correct multiple returns
		take = 2
		results, rolls, err = die.Highest(timesToRoll,  take)
		if err != nil {
			t.Errorf("Encountered an error.")
		}
		rollsEnd := rolls[len(rolls)-take:]
		for i := 0; i < take; i++ {
			if results[i] != rollsEnd[i] {
				t.Errorf("Returned Results do not match highest.  %v vs %v", results, rolls)
			}
		}
		// test for failure on invalid take
		take = 0
		_, _, err = die.Highest(timesToRoll, take)
		if err == nil {
			t.Errorf("Did not receive error on invalid input.")
		}
		take = timesToRoll+2
		_, _, err = die.Highest(timesToRoll, take)
		if err == nil {
			t.Errorf("Did not receive error when take > times to roll")
		}
	}
}

func TestDie_Lowest(t *testing.T) {
	for _, die := range dice {
		timesToRoll := 4
		take := 1
		results, rolls, err := die.Lowest(timesToRoll, take)
		if err != nil {
			t.Errorf("Encountered an error.")
		}
		if results[0] != rolls[0] {
			t.Errorf("Did not receive lowest value.  Received %v, all results %v", results, rolls)
		}
		// test for correct multiple returns
		take = 2
		results, rolls, err = die.Lowest(timesToRoll,  take)
		if err != nil {
			t.Errorf("Encountered an error.")
		}
		rollsStart := rolls[:take]
		for i := 0; i < take; i++ {
			if results[i] != rollsStart[i] {
				t.Errorf("Returned Results do not match highest.  %v vs %v", results, rolls)
			}
		}
		// test for failure on invalid take
		take = 0
		_, _, err = die.Lowest(timesToRoll, take)
		if err == nil {
			t.Errorf("Did not receive error on invalid input of 0.")
		}
		take = timesToRoll+2
		_, _, err = die.Lowest(timesToRoll, take)
		if err == nil {
			t.Errorf("Did not receive error when take > times to roll")
		}
	}
}