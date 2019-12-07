package main

import (
	"fmt"
	"math"

	"github.com/cthierer/adventofcode2019/day07/amplifier"
	"github.com/cthierer/adventofcode2019/day07/intcode"
)

func runSequence(phaseSettings []int, program *intcode.Program) (int, error) {
	amps := &amplifier.Collection{}

	for _, phaseSetting := range phaseSettings {
		amps.Add(&amplifier.Amplifier{PhaseSetting: phaseSetting})
	}

	val, err := amps.Run(program)
	if err != nil {
		return 0, err
	}

	return val, err
}

func getPhaseSettings(num int, numSettings, lowestSetting int) (bool, []int) {
	signals := make([]int, numSettings)
	for idx := range signals {
		val := int(math.Trunc(math.Mod(float64(num)/math.Pow(10, float64(idx)), 10))) - 1 + lowestSetting
		if val < lowestSetting || val >= numSettings+lowestSetting {
			return false, nil
		}

		for i := 0; i < idx; i++ {
			if signals[i] == val {
				return false, nil
			}
		}

		signals[idx] = val
	}
	return true, signals
}

func main() {
	program := intcode.Program{Instructions: []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 42, 67, 88, 105, 114, 195, 276, 357, 438, 99999, 3, 9, 101, 4, 9, 9, 102, 3, 9, 9, 1001, 9, 2, 9, 102, 4, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 4, 9, 102, 4, 9, 9, 101, 2, 9, 9, 1002, 9, 5, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 1001, 9, 4, 9, 1002, 9, 4, 9, 101, 2, 9, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 101, 4, 9, 9, 102, 3, 9, 9, 1001, 9, 5, 9, 4, 9, 99, 3, 9, 102, 5, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99}}
	maxThrusterSignal := 0

	for i := 11111; i <= 55555; i++ {
		valid, phaseSettings := getPhaseSettings(i, 5, 5)
		if !valid {
			continue
		}

		thrusterSignal, err := runSequence(phaseSettings, &program)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			return
		}

		if thrusterSignal > maxThrusterSignal {
			maxThrusterSignal = thrusterSignal
		}
	}

	fmt.Printf("maxThrusterSignal = %d\n", maxThrusterSignal)
}
