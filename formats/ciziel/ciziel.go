package ciziel

import (
	"bufio"
	"errors"
	"io"
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
	d := decoder{
		src: bufio.NewReader(source),
	}
	d.ancestors = append(d.ancestors, &d.top)
	d.scan = (*decoder).scanStart
	d.doScan()
	//TODO
	return Table{}
}

////////////////////////////////////////////////////////////////////////////////

const (
	specials = "\",:[]()<>"
)

type decoder struct {
	src       *bufio.Reader
	top       Table
	ancestors []*Table
	table     *Table
	section   string
	builder   strings.Builder
	scan      scanner
	line, col uint32
	err       error
}

type scanner func(*decoder) scanner

func (a *decoder) doScan() {
	for a.scan != nil {
		a.scan = a.scan(a)
	}
}

func (a *decoder) scanStart() scanner {
	r, _, err := a.src.ReadRune()
	if err != nil {
		a.err = err
		return nil
	}
	a.col++
	switch r {
	case '\n':
		a.line++
		a.col = 1
		return (*decoder).scanStart
	case ' ', '\t':
		return (*decoder).scanStart
	case '"':
		//TODO: start a quote
		return (*decoder).scanQuote
	case '(':
		//TODO: start a new table
		return (*decoder).scanStart
	case ')':
		l := len(a.ancestors)
		if l > 0 {
			a.table = a.ancestors[l-1]
			a.ancestors = a.ancestors[:l-1]
			return (*decoder).scanAfterTable
		}
		if a.err == nil {
			//TODO: add position
			a.err = errors.New("unexpected ')' in top-level table")
		}
		return nil
	case ',':
		a.table.header = append(a.table.header, String{
			line: a.line,
			col:  a.col,
		})
		return (*decoder).scanStart
	case ':':
		if a.err == nil {
			//TODO: add position
			a.err = errors.New("unexpected ':' with no label")
		}
		return (*decoder).scanStart
	case '<':
		//TODO: start a shortcut
		return (*decoder).scanShortcut
	case '>':
		if a.err == nil {
			//TODO: add position
			a.err = errors.New("unexpected '>' outside shortcut")
		}
		return (*decoder).scanStart
	case '[':
		return (*decoder).scanComment
	case ']':
		if a.err == nil {
			//TODO: add position
			a.err = errors.New("unexpected ']' outside comment")
		}
		return (*decoder).scanStart
	default:
		a.builder.WriteRune(r)
		return (*decoder).scanBasic
	}
}

func (a *decoder) scanBasic() scanner {
	r, _, err := a.src.ReadRune()
	if err != nil {
		a.err = err
		return nil
	}
	a.col++
	switch r {
	case '\n':
		a.line++
		a.col = 1
		fallthrough
	case ',':
		// It's a basic string
		s := String{
			string: a.builder.String(),
		}
		if a.section == "" {
			a.table.header = append(a.table.header, s)
		} else {
			a.table.sections[a.section] = append(a.table.sections[a.section], s)
		}
		a.builder.Reset()
		return (*decoder).scanStart
	case '(':
	case ')':
		// It's a basic string
		s := String{
			string: a.builder.String(),
		}
		if a.section == "" {
			a.table.header = append(a.table.header, s)
		} else {
			a.table.sections[a.section] = append(a.table.sections[a.section], s)
		}
		a.builder.Reset()
	case ':':
	case '<':
	case '>':
	case '[':
	case ']':
	case '"':
	case ' ', '\t':
		fallthrough
	default:
	}
	return nil
}

func (a *decoder) scanComment() scanner {
	r, _, err := a.src.ReadRune()
	if err != nil {
		a.err = err
		return nil
	}
	a.col++
	switch r {
	case '\n':
	case ' ', '\t':
	case '"':
	case '(':
	case ')':
	case ',':
	case ':':
	case '<':
	case '>':
	case '[':
	case ']':
	default:
	}
	return nil
}

func (a *decoder) scanQuote() scanner      { return nil }
func (a *decoder) scanShortcut() scanner   { return nil }
func (a *decoder) scanAfterTable() scanner { return nil }
