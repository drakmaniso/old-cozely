package ciziel

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////

type Definition struct {
	Line     int
	indent   int
	Label    string
	Elements []string
	Children []Definition
}

func (a Definition) String() string {
	return a.string(0)
}

func (a Definition) string(level int) string {
	var s string // = strconv.Itoa(a.Line) + ": "
	for i := 0; i < level; i++ {
		s += "    "
	}
	if a.Label != "" {
		s += a.Label + ":"
	}
	for i, e := range a.Elements {
		if i == 0 {
			if a.Label != "" {
				s += " "
			}
		} else {
			s += ", "
		}
		s += e
	}
	s += "\n"
	for _, c := range a.Children {
		s += c.string(level + 1)
	}
	return s
}

////////////////////////////////////////////////////////////////////////////////

type state struct {
	src *bufio.Reader

	previous  token
	builder   strings.Builder
	line, col int

	top     Definition
	current *Definition
	indent  int

	err error
}

type token int

const (
	invalid token = iota
	newline
	whitespace
	basic
	quote
	colon
	coma
	comment
	eof
)

////////////////////////////////////////////////////////////////////////////////

func Parse(source io.Reader) []Definition {
	s := state{
		src:      bufio.NewReader(source),
		line:     1,
		previous: newline,
		top:      Definition{indent: 0},
	}

loop:
	for {
		t := s.scan()

		switch t {
		case invalid:

		case newline:
			s.addElement()
			s.current = nil
			s.indent = 0

		case whitespace:
			if s.previous == newline {
				s.indent = 0
				if t == whitespace {
					s.indent = countsp(s.builder.String())
				}
			}
			s.builder.Reset()

		case basic:

		case quote:
			s.builder.Reset()

		case colon:
			s.addLabel()

		case coma:
			s.addElement()

		case comment:
			s.builder.Reset()

		case eof:
			break loop

		default:
			println("*** what? ***")
		}
		s.previous = t
	}

	return s.top.Children
}

////////////////////////////////////////////////////////////////////////////////

func countsp(s string) int {
	r := 0
	for _, c := range s {
		switch c {
		case ' ':
			r++
		case '\t':
			r += 0x1000
		default:
			//TODO: error
		}
	}
	return r
}

func (s *state) startDef() {
	c := &s.top
	p := c
	for s.indent > c.indent && len(c.Children) > 0 {
		p = c
		c = &c.Children[len(c.Children)-1]
	}

	switch {
	case s.indent > c.indent:
		c.Children =
			append(c.Children, Definition{
				Line:   s.line,
				indent: s.indent,
			})
		s.current = &c.Children[len(c.Children)-1]

	case s.indent < c.indent:
		//TODO: err
		s.errmsg(s.line, s.col, "misaligned indentation")
		fallthrough
	default: // s.indent == c.indent:
		p.Children =
			append(p.Children, Definition{
				Line:   s.line,
				indent: s.indent,
			})
		s.current = &p.Children[len(p.Children)-1]
	}
}

func (s *state) addLabel() {
	if s.previous == basic {
		if s.current == nil {
			s.startDef()
		}
		var l string
		if s.previous == basic {
			l = strings.ToLower(s.builder.String())
		} else {
			l = s.builder.String()
		}
		if s.current.Label != "" {
			//TODO: error
		}
		s.current.Label = l
	}
	s.builder.Reset()
}

func (s *state) addElement() {
	if s.previous == basic {
		if s.current == nil {
			s.startDef()
		}
		var e string
		if s.previous == basic {
			e = strings.ToLower(s.builder.String())
		} else {
			e = s.builder.String()
		}
		s.current.Elements = append(s.current.Elements, e)
	}
	s.builder.Reset()
}

////////////////////////////////////////////////////////////////////////////////

func (s *state) scan() token {
	r, _, err := s.src.ReadRune()
	if err != nil {
		s.err = err
		return eof
	}
	switch r {
	case '\n':
		s.line++
		s.col = 1
		return newline

	case ' ', '\t':
	spaceloop:
		for {
			switch r {
			case ' ', '\t': //TODO: unicode whitespace?
				s.col++
				s.builder.WriteRune(r)
			default:
				s.src.UnreadRune()
				break spaceloop
			}
			r, _, err = s.src.ReadRune()
			if err != nil {
				break spaceloop
			}
		}
		return whitespace

	case '"':
		//TODO
		return quote

	case ',':
		s.col++
		return coma

	case ':':
		s.col++
		return colon

	case '(':
		//TODO
		return comment

	case ')':
		s.errmsg(s.line, s.col, "unexpected ')' outside comment")
		s.col++
		return invalid

	case '{':
		//TODO
		return comment

	case '}':
		s.errmsg(s.line, s.col, "unexpected '}' outside comment")
		s.col++
		return invalid

	default:
		sp := false
	basicloop:
		for {
			switch r {
			case '\n', '(', ')', '{', '}', ',', ':':
				s.src.UnreadRune()
				break basicloop
			case ' ', '\t': //TODO: unicode whitespace?
				s.col++
				sp = true
			default:
				s.col++
				if sp {
					s.builder.WriteRune(' ')
					sp = false
				}
				s.builder.WriteRune(r)
			}
			r, _, err = s.src.ReadRune()
			if err != nil {
				break basicloop
			}
		}
		return basic
	}
}

////////////////////////////////////////////////////////////////////////////////

func (s *state) errmsg(l, c int, msg string) {
	if s.err == nil {
		s.err = errors.New(
			strconv.Itoa(int(l)) + "." + strconv.Itoa(int(c)) + ": " + msg,
		)
	}
}
