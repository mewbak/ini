// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package ini

import (
	"io"
	"unicode"
	"unicode/utf8"
)

type TokenType uint8

const (
	TokSection TokenType = iota
	TokKey
	TokKeyList
	TokValue
)

type TokenHandler func(TokenType, []byte)

// Reader is a lexer for ini file data.
type Reader struct {
	out   TokenHandler
	data  []byte
	start int
	pos   int
	size  int
}

// Read reads ini elements until EOF.
// Generated tokens are passed into a call to the supplied token handler.
func (ir *Reader) Read(data []byte, f TokenHandler) {
	if len(data) == 0 || data[len(data)-1] != '\n' {
		data = append(data, '\n')
	}

	ir.out = f
	ir.data = data
	ir.start = 0
	ir.pos = 0

	for {
		if ir.comment() {
			continue
		}

		if ir.section() {
			continue
		}

		if ir.key() {
			continue
		}

		return
	}
}

func (ir *Reader) comment() bool {
	if ir.accept(';') {
		for !ir.acceptUntil('\n') {
			ir.skip()
		}
		ir.skip()
		return true
	}
	return false
}

func (ir *Reader) section() bool {
	ir.acceptSpace()
	ir.ignore()

	if !ir.accept('[') {
		return false
	}

	ir.ignore()

	if !ir.acceptUntil(']') {
		return false
	}

	ir.out(TokSection, ir.data[ir.start:ir.pos])
	ir.ignore()
	ir.skip()
	return true
}

func (ir *Reader) key() bool {
	ir.acceptSpace()
	ir.ignore()

	r, ok := ir.acceptUntilAny('<', '=')

	if !ok {
		return false
	}

	if r == '<' {
		ir.out(TokKeyList, ir.data[ir.start:ir.pos])
	} else {
		ir.out(TokKey, ir.data[ir.start:ir.pos])
	}

	ir.ignore()
	ir.skip()

	return ir.value()
}

func (ir *Reader) value() bool {
	if !ir.acceptUntil('\n') {
		return false
	}

	ir.out(TokValue, ir.data[ir.start:ir.pos])
	ir.ignore()
	return true
}

// accept consumes the next rune, only if it is the same as the supplied rune.
func (ir *Reader) accept(v rune) bool {
	next, err := ir.next()

	if err != nil {
		return false
	}

	if v != next {
		ir.rewind()
		return false
	}

	return true
}

// acceptUntil consumes runes for as long as they
// are /not/ equal to the supplied rune.
func (ir *Reader) acceptUntil(v rune) bool {
	pos := ir.pos

	for {
		next, err := ir.next()

		if err != nil {
			return false
		}

		if v == next {
			ir.rewind()
			break
		}
	}

	return ir.pos > pos
}

// acceptUntilAny consumes runes for as long as they
// are /not/ in the list of runes
func (ir *Reader) acceptUntilAny(argv ...rune) (r rune, b bool) {
	var err error
	pos := ir.pos

	for {
		if r, err = ir.next(); err != nil {
			return
		}

		if hasRune(argv, r) {
			ir.rewind()
			break
		}
	}

	b = ir.pos > pos
	return
}

// acceptSpace consumes runes for as long as they
// qualify as unicode whitespace.
func (ir *Reader) acceptSpace() bool {
	pos := ir.pos

	for {
		next, err := ir.next()

		if err != nil {
			return false
		}

		if !unicode.IsSpace(next) {
			ir.rewind()
			break
		}
	}

	return ir.pos > pos
}

// next retuns the next unicode codepoint in the input.
func (ir *Reader) next() (rune, error) {
	if ir.pos >= len(ir.data) {
		return 0, io.EOF
	}

	var r rune
	r, ir.size = utf8.DecodeRune(ir.data[ir.pos:])
	ir.pos += ir.size
	return r, nil
}

// ignore the input so far.
func (ir *Reader) ignore() { ir.start = ir.pos }

// rewind rewinds to the last rune.
// Can be called only once per nextRune() call.
func (ir *Reader) rewind() { ir.pos -= ir.size }

// skip Skips the nextRune character.
func (ir *Reader) skip() {
	ir.next()
	ir.ignore()
}

// Emit emits a new token if the given type.
func (ir *Reader) emit(tt TokenType) {
	ir.out(tt, ir.data[ir.start:ir.pos])
}

func hasRune(v []rune, r rune) bool {
	for _, vr := range v {
		if vr == r {
			return true
		}
	}

	return false
}
