package core

type BufferView interface {
	GetLineCount() int
	//GetLineLength(index int) //TODO
	GetLine(index int) []rune
	//GetLinePart(index, from, to int) []rune //TODO
	Update(from, to int)
	GetCursorIndex(line, char int) (int, error)
}

type defaultBufferView struct {
	buffer      Buffer
	lineIndices []int
}

func GetBufferView(buffer Buffer) BufferView {
	bufferView := defaultBufferView{buffer: buffer}
	bufferView.Update(0, buffer.Size())
	return BufferView(&bufferView)
}

func (bw *defaultBufferView) GetLineCount() int {
	return len(bw.lineIndices) + 1
}

func (bw *defaultBufferView) Update(from, to int) {
	if from != 0 || to != bw.buffer.Size() {
		panic("TODO - can only be called with full range for now")
	}
	bw.lineIndices = nil
	for i, r := range bw.buffer.Read(from, to) {
		if r == '\n' {
			bw.lineIndices = append(bw.lineIndices, i)
		}
	}
}

func (bw *defaultBufferView) GetLine(index int) []rune {
	lastLineIndex := len(bw.lineIndices)

	var from, to int

	switch {
	case index == 0:
		from = 0
		to = bw.lineIndices[0]

	case index < lastLineIndex:
		from = bw.lineIndices[index-1] + 1
		to = bw.lineIndices[index]

	case index == lastLineIndex:
		from = bw.lineIndices[index-1] + 1
		to = bw.buffer.Size()

	default:
		from, to = 0, 0 //TODO
	}

	return bw.buffer.Read(from, to)
}

func (bw *defaultBufferView) GetCursorIndex(line, char int) (int, error) {
	if bw.buffer.Size() == 0 {
		return 0, EmptyBufferError{}
	} else if line == 0 {
		return char, nil
	} else {
		cursorIndex := bw.lineIndices[line-1] + char + 1
		if cursorIndex < bw.buffer.Size() {
			return cursorIndex, nil
		} else {
			return bw.buffer.Size() - 1, IndexError{}
		}
	}
}
