package ciziel

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////

type Data interface {
	Position() (line, col int)
	String() string
	add(st *state, d Data)
}

////////////////////////////////////////////////////////////////////////////////

type String struct {
	line, col int
	string
}

func (s String) Position() (line, col int) {
	return s.line, s.col
}

func (s String) String() string {
	return s.string
}

func (s String) add(st *state, d Data) {
	panic(nil)
}

////////////////////////////////////////////////////////////////////////////////

type Array []Data

func (a Array) Position() (line, col int) {
	if len(a) == 0 {
		return 0, 0
	}
	return a[0].Position()
}

func (a Array) String() string {
	s := "["
	for i, d := range a {
		if i > 0 {
			s += ", "
		}
		s += d.String()
	}
	return s + "]"
}

func (a *Array) add(st *state, d Data) {
	*a = append(*a, d)
}

////////////////////////////////////////////////////////////////////////////////

type Map map[string]Data

func (m Map) Position() (line, col int) {
	l, c := 0, 0
	for _, d := range m {
		ll, cc := d.Position()
		if l > ll {
			l = ll
			c = cc
		} else if l == ll {
			if c > cc {
				c = cc
			}
		}
	}
	return l, c
}

func (m Map) String() string {
	s := "{"
	f := true // first iteration
	for k, d := range m {
		if !f {
			s += ", "
		}
		s += k + ": " + d.String()
		f = false
	}
	return s + "}"
}

func (m *Map) add(st *state, d Data) {
	k := st.live.key()
	if k == nil {
		//TODO: error
		panic(nil)
	}
	v, ok := (*m)[*k]
	if ok {
		(*m)[*k] = &Array{v, d}
		return
	}
	(*m)[*k] = d
}

////////////////////////////////////////////////////////////////////////////////

type state struct {
	src *bufio.Reader

	previous    token
	line, col   int
	builder     strings.Builder
	bline, bcol int

	live stack

	err error
}

////////////////////////////////////////////////////////////////////////////////

func Parse(source io.Reader) Data {
	st := state{
		src:  bufio.NewReader(source),
		line: 1,
	}
	st.live.push(nil)

loop:
	for {
		t := st.scan()

		switch t {
		case invalid:
			//TODO
			panic(nil)

		case open:
			st.live.push(nil)

		case close:
			if st.live.peek() == nil {
				st.live.replace(&Array{})
			}
			if st.previous == basic {
				s := String{
					line: st.bline, col: st.bcol,
					string: st.builder.String(),
				}
				st.builder.Reset()
				st.live.peek().add(&st, s)
			}
			d := st.live.pop()
			st.live.peek().add(&st, d)

		case separator:
			if st.previous == basic {
				if st.live.peek() == nil {
					st.live.replace(&Array{})
				}
				s := String{
					line: st.bline, col: st.bcol,
					string: st.builder.String(),
				}
				st.builder.Reset()
				st.live.peek().add(&st, s)
			}

		case whitespace:

		case basic:

		case quote:
			st.builder.Reset()

		case colon:
			if st.live.peek() == nil {
				st.live.replace(&Map{})
			}
			s := st.builder.String()
			st.builder.Reset()
			st.live.setkey(&s)
			println("LABEL: ", s)

		case comment:
			st.builder.Reset()

		case eof:
			break loop

		default:
			println("*** what? ***")
		}
		st.previous = t
	}

	return st.live.stack[0]
}

////////////////////////////////////////////////////////////////////////////////

type token int

const (
	invalid token = iota
	whitespace
	separator
	basic
	quote
	open
	close
	colon
	comment
	eof
)

func (s *state) scan() token {
	s.bline, s.bcol = s.line, s.col
	r, _, err := s.src.ReadRune()
	if err != nil {
		s.err = err
		return eof
	}
	switch r {
	case '\n':
		s.line++
		s.col = 1
		return separator

	case ' ', '\t':
	spaceloop:
		for {
			switch r {
			case ' ', '\t': //TODO: unicode whitespace?
				s.col++
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
		return separator

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
		s.col++
		return open

	case '}':
		s.col++
		return close

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
