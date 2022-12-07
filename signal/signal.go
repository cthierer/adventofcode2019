package signal

type buffer struct {
	contents []rune
	pos      int
}

func (b *buffer) find(c rune) int {
	for i := 0; i <= b.pos; i += 1 {
		if b.contents[i] == c {
			return i
		}
	}

	return -1
}

func (b *buffer) add(c rune) {
	b.contents[b.pos] = c
	b.pos += 1
}

func (b *buffer) shift(n int) {
	shifted := make([]rune, len(b.contents))
	for i := n; i < len(b.contents); i += 1 {
		shifted[i-n] = b.contents[i]
	}
	b.contents = shifted
	b.pos -= n
}

func (b *buffer) full() bool {
	return b.pos >= len(b.contents)
}

func FindStartOfPacket(stream string) int {
	b := buffer{contents: make([]rune, 4)}

	for i := 0; i < len(stream); i += 1 {
		c := rune(stream[i])
		if match := b.find(c); match >= 0 {
			b.shift(match + 1)
		}

		b.add(c)
		if b.full() {
			return i + 1
		}
	}

	return -1
}
