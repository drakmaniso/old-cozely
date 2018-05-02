package data

import (
	"io"
)

////////////////////////////////////////////////////////////////////////////////

type Data interface {
}

////////////////////////////////////////////////////////////////////////////////

type String []byte

////////////////////////////////////////////////////////////////////////////////

type List struct {
	items []struct {
		Data
		line, col uint32
	}
	pairs map[string]struct {
		Data
		line, col uint32
	}
}

func Parse(source io.Reader) *List {
	//TODO
	return nil
}

func (a *List) At(index int) Data {
	return a.items[index].Data
}

func (a *List) With(key string) Data {
	return a.pairs[key].Data
}

func (a *List) LocationAt(index int) (line, col int) {
	return int(a.items[index].line), int(a.items[index].col)
}

func (a *List) LocationWith(key string) (line, col int) {
	return int(a.pairs[key].line), int(a.pairs[key].col)
}
