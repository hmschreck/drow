package drow

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
	"errors"
)

// Representation of a die.  Mostly for doing methods on.
type Die struct {
	Randomizer rand.Rand
	Sides int
	ExpectedValue float64
	Name string
}

// Generate a new die with specified number of sides
func NewDie(sides int) *Die {
	new_die := new(Die)
	new_die.Sides = sides
	new_die.Name = fmt.Sprintf("D%d", sides)
	new_die.Randomizer = *rand.New(rand.NewSource(time.Now().UnixNano()))
	new_die.ExpectedValue = float64((float64(new_die.Sides)+float64(1))/float64(2))
	return new_die
}

// Roll a die
func(die *Die) Roll() int {
	return die.Randomizer.Intn(die.Sides) + 1
}

func (die *Die) RollMultiple(times int) []int {
	rolls := make([]int, times)
	for i := 0; i < times; i++ {
		rolls[i] = die.Roll()
	}
	return rolls
}

// Takes number of times to roll, as well an integer for what to keep - negative numbers will keep lowest,
// positive numbers will keep highest.  For example, -2 will keep the lowest 2, 2 will keep the highest two.
func (die *Die) HighLow(times int, take int) (results []int, allRolls []int, err error) {
	// Cannot take more than you roll, obviously.
	if math.Abs(float64(take)) > float64(times) {
		return make([]int, 0), make([]int, 0), errors.New("Cannot take more than rolled.")
	}
	if take == 0 {
		return make([]int, 0), make([]int, 0), errors.New("Cannot take zero dice")
	}
	// Roll our dice "times" times and sort them in ascending order
	rolls := die.RollMultiple(times)
	sort.Ints(rolls)
	// If we are taking
	if take > 0 {
		return rolls[len(rolls)-take:], rolls, nil
	} else {
		// Absolute value of take
		take = 0 - take
		return rolls[:take], rolls, nil
	}
}

func (die *Die) Highest(times int, take int) (results []int, allRolls []int, err error) {
	return die.HighLow(times, take)
}

func (die *Die) Lowest(times int, take int) (results []int, allRolls []int, err error) {
	return die.HighLow(times, 0-take)
}

func (die *Die) RollAtAdvantage() (result int, rolls []int, err error) {
	_result, _rolls, err := die.Highest(2, 1)
	return _result[0], _rolls, err
}

func (die *Die) RollAtDisadvantage() (result int, rolls []int, err error) {
	_result, _rolls, err := die.Lowest(2, 1)
	return _result[0], _rolls, err
}

var D20 = NewDie(20)
var D4  = NewDie(4)
var D6 =  NewDie(6)
var D8 =  NewDie(8)
var D10 = NewDie(10)
var D12 = NewDie(12)
var D100 =NewDie(100)





