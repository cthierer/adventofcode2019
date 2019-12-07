package amplifier

// Signal represents an amplifier signal.
type Signal struct {
	Value int
}

// WriteInt sets the value of the signal.
func (s *Signal) WriteInt(val int) {
	s.Value = val
}
