package bowling

import (
	"strconv"
	"strings"
)

func Score(game string) int {
	score := 0
	frames := strings.Split(game, " ")

	for _, frame := range frames {
		balls := strings.Split(frame, "")
		if balls[0] != "-" {
			pins, _ := strconv.Atoi(balls[0])
			score += pins
		}

		if balls[1] != "-" {
			pins, _ := strconv.Atoi(balls[1])
			score += pins
		}
	}

	return score
}
