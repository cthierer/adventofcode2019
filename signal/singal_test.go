package signal_test

import (
	"testing"

	"github.com/cthierer/advent-of-code/signal"
)

func assertStartOfPacket(t *testing.T, input string, expected int) {
	if packet := signal.FindStartOfPacket(input); packet != expected {
		t.Logf("found substring: %v", input[packet-4:packet])
		t.Logf("expected to find start of packet after %v characters, but instead found %v", expected, packet)
		t.Fail()
	}
}

func assertStartOfMessage(t *testing.T, input string, expected int) {
	if packet := signal.FindStartOfMessage(input); packet != expected {
		t.Logf("found substring: %v", input[packet-14:packet])
		t.Logf("expected to find start of message after %v characters, but instead found %v", expected, packet)
		t.Fail()
	}
}

func TestFindStartOfPacket(t *testing.T) {
	assertStartOfPacket(t, "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7)
	assertStartOfPacket(t, "bvwbjplbgvbhsrlpgdmjqwftvncz", 5)
	assertStartOfPacket(t, "nppdvjthqldpwncqszvftbrmjlhg", 6)
	assertStartOfPacket(t, "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10)
	assertStartOfPacket(t, "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11)
}

func TestFindStartOfMessage(t *testing.T) {
	assertStartOfMessage(t, "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19)
	assertStartOfMessage(t, "bvwbjplbgvbhsrlpgdmjqwftvncz", 23)
	assertStartOfMessage(t, "nppdvjthqldpwncqszvftbrmjlhg", 23)
	assertStartOfMessage(t, "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29)
	assertStartOfMessage(t, "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26)
}
