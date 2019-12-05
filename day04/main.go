package main

import (
	"fmt"
	"strconv"
)

const (
	startVal = 372304
	endVal   = 847060
)

func nextNearest(fromVal int) int {
	valStr := strconv.Itoa(fromVal)
	digits := make([]int, len(valStr))
	for i, c := range valStr {
		curr, _ := strconv.Atoi(string(c))
		if i == 0 {
			digits[i] = curr
			continue
		}

		last := digits[i-1]
		if curr >= last {
			digits[i] = curr
			continue
		}

		break
	}
	var digitStr []byte
	for i, d := range digits {
		val := int64(d)
		if val == 0 && i > 0 {
			digitStr = append(digitStr, digitStr[i-1])
		} else {
			digitStr = strconv.AppendInt(digitStr, int64(d), 10)
		}
	}
	res, _ := strconv.Atoi(string(digitStr))
	return res
}

func prevNearest(fromVal int) int {
	valStr := strconv.Itoa(fromVal)
	digits := make([]int, len(valStr))
	for i, c := range valStr {
		curr, _ := strconv.Atoi(string(c))
		if i == 0 {
			digits[i] = curr
			continue
		}

		last := digits[i-1]
		if curr >= last {
			digits[i] = curr
			continue
		}

		digits[i-1] = digits[i-1] - 1
		break
	}
	var digitStr []byte
	for i, d := range digits {
		val := int64(d)
		if val == 0 && i > 0 {
			digitStr = append(digitStr, []byte("9")...)
		} else {
			digitStr = strconv.AppendInt(digitStr, int64(d), 10)
		}
	}
	res, _ := strconv.Atoi(string(digitStr))
	return res
}

func matches(val int) bool {
	valStr := strconv.Itoa(val)
	last := -1
	streaking := false
	for _, c := range valStr {
		curr, _ := strconv.Atoi(string(c))
		if last == -1 {
			last = curr
			continue
		}

		if curr == last {
			streaking = true
			continue
		}

		if curr > last {
			last = curr
			continue
		}

		return false
	}

	return streaking
}

func matches2(val int) bool {
	valStr := strconv.Itoa(val)
	last := -1
	streakLen := 1
	streaking := false
	for _, c := range valStr {
		curr, _ := strconv.Atoi(string(c))
		if last == -1 {
			last = curr
			continue
		}

		if curr == last {
			streakLen++
			continue
		}

		if streakLen == 2 {
			streaking = true
		}

		if curr > last {
			last = curr
			streakLen = 1
			continue
		}

		return false
	}

	return streaking || streakLen == 2
}

func main() {
	startAt := nextNearest(startVal)
	endAt := prevNearest(endVal)
	count := 0

	for i := startAt; i <= endAt; i++ {
		if matches(i) {
			count++
		}
	}

	fmt.Printf("%d matches\n", count)

	count = 0
	
	for i := startAt; i <= endAt; i++ {
		if matches2(i) {
			count++
		}
	}

	fmt.Printf("%d matches\n", count)
}
