package drow

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestParserXdY(t *testing.T) {
	for _, die := range dice {
		for i := 0; i < 100; i++ {
			rollTimes := rand.Intn(50)
			result := ParseString(fmt.Sprintf("%dd%d", rollTimes, die.Sides))
		}
	}

}
