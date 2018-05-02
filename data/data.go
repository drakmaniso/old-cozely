package data

import (
	"errors"
	"io"
)

////////////////////////////////////////////////////////////////////////////////

type Data interface {
	Label() []byte
	Source() (line, col int)
}

////////////////////////////////////////////////////////////////////////////////

type String struct {
	label     []byte
	value     []byte
	line, col uint32
}

func (a String) Label() []byte {
	return a.label
}

func (a String) Source() (line, col int) {
	return int(a.line), int(a.col)
}

////////////////////////////////////////////////////////////////////////////////

type List struct {
	label      []byte
	items      []Data
	dictionary map[string][]Data
	line, col  uint32
}

func Parse(source io.Reader) List {
	//TODO
	return List{}
}

// Items returns a slice of all the items of the list (with or without
// labels).
func (a List) Items() []Data {
	return a.items
}

// WithLabel returns a slice of all items with a specific label in the list.
func (a List) WithLabel(label string) []Data {
	return a.dictionary[label]
}

// WithName tries to find the only item in the list with a specific label. If
// not present, nil is returned. If multiple items have the same label, the last
// one is returned, as well as an error.
func (a List) WithName(label string) (Data, error) {
	d := a.dictionary[label]
	switch len(d) {
	case 0:
		return nil, nil
	case 1:
		return d[0], nil
	default:
		return d[len(d)-1], errors.New("ambiguous data: label \"" + label + "\" is used multiple times inside the same list")
	}
}

// WithoutLabel return the n-th unlabeled item of the list (or nil if there
// isn't enough items).
func (a List) WithoutLabel(n int) Data {
	j := 0
	for range a.items {
		if j == n {
			return a.items[j]
		}
		if len(a.items[j].Label()) > 0 {
			j++
		}
	}
	return nil
}

func (a List) Label() []byte {
	return a.label
}

func (a List) Source() (line, col int) {
	return int(a.line), int(a.col)
}
