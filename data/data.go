package data

import "io"

////////////////////////////////////////////////////////////////////////////////

type Data interface {
	Interpret(p ...InterpretFunc) interface{}
}

type InterpretFunc func(a Data, f ...InterpretFunc) interface{}

////////////////////////////////////////////////////////////////////////////////

type String []byte

func (a String) Interpret(f ...InterpretFunc) interface{} {
	for _, ff := range f {
		r := ff(a, f...)
		if r != nil {
			return r
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type List struct {
	Items    []Data
	Pairs    map[string]Data
	itemsPos []struct{ line, col int }
	pairsPos map[string]struct{ line, col int }
}

func Parse(source io.Reader) *List {
	//TODO
	return nil
}

func (a *List) Interpret(f ...InterpretFunc) interface{} {
	for _, ff := range f {
		r := ff(a, f...)
		if r != nil {
			return r
		}
	}
	return nil
}
