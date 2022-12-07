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

func findUniqueSubstring(b *buffer, in string) int {
	for i := 0; i < len(in); i += 1 {
		c := rune(in[i])
		if match := b.find(c); match >= 0 {
			b.shift(match + 1)
		}

		b.add(c)
		if b.full() {
			return i
		}
	}
	return -1
}

func findStartOfSegment(stream string, segmentLen int) int {
	b := buffer{contents: make([]rune, segmentLen)}
	if last := findUniqueSubstring(&b, stream); last >= 0 {
		return last + 1
	}
	return -1
}

const (
	lenStartOfPacket  = 4
	lenStartOfMessage = 14
)

func FindStartOfPacket(stream string) int {
	return findStartOfSegment(stream, lenStartOfPacket)
}

func FindStartOfMessage(stream string) int {
	return findStartOfSegment(stream, lenStartOfMessage)
}
