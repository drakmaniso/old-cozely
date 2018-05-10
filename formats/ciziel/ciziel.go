package ciziel

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////

// A Data is either a String or a Table
type Data interface {
	Position() (line, col int)
}

////////////////////////////////////////////////////////////////////////////////

// A String is a sequence of characters.
type String struct {
	string
	line, col uint32
}

// String returns the content of the String.
func (a String) String() string {
	return a.string
}

// Position returns the position of the string in the source file.
func (a String) Position() (line, col int) {
	return int(a.line), int(a.col)
}

////////////////////////////////////////////////////////////////////////////////

// A Table is composed of a header and a set of sections (both of wich are
// optional).
type Table struct {
	header    []Data
	sections  map[string][]Data
	line, col uint32
}

// Position returns the position of the table in the source file.
func (a Table) Position() (line, col int) {
	return int(a.line), int(a.col)
}

// Header returns the values not associated with any labels (i.e. the values
// occuring before any definition).
func (a Table) Header() []Data {
	return a.header
}

// Section returns the values associated with a label.
func (a Table) Section(label string) []Data {
	return a.sections[label]
}

////////////////////////////////////////////////////////////////////////////////

func Parse(source io.Reader) Table {
	s := state{
		src:  bufio.NewReader(source),
		line: 1, col: 1,
	}
	s.ancestors = append(s.ancestors, &s.top)
loop:
	for {
		l, c := s.line, s.col
		t := s.scan()
		switch t {
		case invalid:
			println(l, ".", c, ": ", "invalid")

		case newline:
			println(l, ".", c, ": ", "newline")

		case whitespace:
			println(l, ".", c, ": ", "whitespace")

		case basic:
			println(l, ".", c, ": ", "basic")

		case quote:
			println(l, ".", c, ": ", "quote")

		case open:
			println(l, ".", c, ": ", "open")

		case colon:
			println(l, ".", c, ": ", "colon")

		case coma:
			println(l, ".", c, ": ", "coma")

		case close:
			println(l, ".", c, ": ", "close")

		case shortcut:
			println(l, ".", c, ": ", "shortcut")

		case comment:
			println(l, ".", c, ": ", "comment")

		case eof:
			println(l, ".", c, ": ", "eof")
			break loop

		default:
			println("*** what? ***")
		}
	}
	//TODO
	return Table{}
}

////////////////////////////////////////////////////////////////////////////////

const (
	specials = "\",:[]()<>"
)

type state struct {
	src       *bufio.Reader
	top       Table
	ancestors []*Table
	table     *Table
	section   string
	builder   strings.Builder
	line, col uint32
	err       error
}

type token int

const (
	invalid token = iota
	newline
	whitespace
	basic
	quote
	open
	colon
	coma
	close
	shortcut
	comment
	eof
)

func (a *state) scan() token {
	r, _, err := a.src.ReadRune()
	if err != nil {
		a.err = err
		return eof
	}
	switch r {
	case '\n':
		a.line++
		a.col = 1
		return newline

	case ' ', '\t':
	spaceloop:
		for {
			switch r {
			case ' ', '\t':
				a.col++
			default:
				a.src.UnreadRune()
				break spaceloop
			}
			r, _, err = a.src.ReadRune()
			if err != nil {
				break spaceloop
			}
		}
		return whitespace

	case '"':
		//TODO
		return quote

	case '(':
		a.col++
		return open

	case ')':
		a.col++
		return close

	case ',':
		a.col++
		return coma

	case ':':
		a.col++
		return colon

	case '<':
		//TODO
		return shortcut

	case '>':
		a.errmsg(a.line, a.col, "unexpected '>' outside shortcut")
		a.col++
		return invalid

	case '[':
		//TODO
		return comment

	case ']':
		a.errmsg(a.line, a.col, "unexpected ']' outside comment")
		a.col++
		return invalid

	default:
	basicloop:
		for {
			switch r {
			case '\n', '(', ')', ',', ':', '<', '>', '[', ']':
				a.src.UnreadRune()
				break basicloop
			default:
				a.col++
				a.builder.WriteRune(r)
			}
			r, _, err = a.src.ReadRune()
			if err != nil {
				break basicloop
			}
		}
		return basic
	}
}

func (a *state) errmsg(l, c uint32, msg string) {
	if a.err == nil {
		a.err = errors.New(
			strconv.Itoa(int(l)) + "." + strconv.Itoa(int(c)) + ": " + msg,
		)
	}
}
