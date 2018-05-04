package data

import (
	"errors"
	"io"
)

////////////////////////////////////////////////////////////////////////////////

// A Data is either a String or a List
type Data interface {
	Label() []byte
	Source() (line, col int)
}

////////////////////////////////////////////////////////////////////////////////

// A String is a sequence of characters.
type String struct {
	label     []byte
	value     []byte
	line, col uint32
}

// Source returns the position of the string in the source file.
func (a String) Source() (line, col int) {
	return int(a.line), int(a.col)
}

// Label returns the label of the string.
func (a String) Label() []byte {
	return a.label
}

// Value returns the content of the string.
func (a String) Value() []byte {
	return a.value
}

////////////////////////////////////////////////////////////////////////////////

// A List is a sequence of Data.
type List struct {
	label      []byte
	items      []Data
	dictionary map[string][]Data
	line, col  uint32
}

// Source returns the position of the list in the source file.
func (a List) Source() (line, col int) {
	return int(a.line), int(a.col)
}

// Label returns the label of the list.
func (a List) Label() []byte {
	return a.label
}

// Items returns a slice of all the items of the list (with or without
// labels).
func (a List) Items() []Data {
	return a.items
}

// FirstWithLabel returns the index of the first occurence of label in the list.
func (a List) FirstWithLabel(label string) int {
	for i, v := range a.items {
		if string(v.Label()) == label {
			return i
		}
	}
	return -1
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

////////////////////////////////////////////////////////////////////////////////

func Parse(source io.ByteScanner) List {
	d := decoder{
		src: source,
	}
	d.scan = (*decoder).scanStart
	d.doScan()
	//TODO
	return List{}
}

////////////////////////////////////////////////////////////////////////////////

const (
	specials = "\",:[]{}()"
)

type decoder struct {
	src       io.ByteScanner
	top       List
	ancestors *List
	item      *Data // The item in construction
	bracket   byte  // Expected closing bracket for lists
	scan      scanner
	err       error
}

type scanner func(*decoder) scanner

func (a *decoder) doScan() {
	for a.scan != nil {
		a.scan = a.scan(a)
	}
}

func (a *decoder) scanStart() scanner {
	b, err := a.src.ReadByte()
	if err != nil {
		//TODO
		panic(err)
	}
	switch b {
	case '[':
		return (*decoder).scanComment
	case ']':
		if a.err == nil {
			//TODO: add position
			a.err = errors.New("unexpected ']' outside comment")
		}
		return (*decoder).scanStart
	case '"':
		//TODO: start a quote
		return (*decoder).scanQuote
	case ',':
		//TODO: add empty item
		return (*decoder).scanStart
	case '(':
		//TODO: start a new list
		return (*decoder).scanStart
	case ')':
		//TODO: close current list
		return (*decoder).scanAfterList
	case '{':
		//TODO: start a shortcut
		return (*decoder).scanShortcut
	case '}':
		if a.err == nil {
			//TODO: add position
			a.err = errors.New("unexpected '}' outside shortcut")
		}
		return (*decoder).scanStart
	case ':':
		//TODO: add a label to current item (check if already one)
		return (*decoder).scanStart
	default:
		//TODO: start basic string
		return (*decoder).scanBasic
	}
}

func (a *decoder) scanComment() scanner   { return nil }
func (a *decoder) scanQuote() scanner     { return nil }
func (a *decoder) scanShortcut() scanner  { return nil }
func (a *decoder) scanAfterList() scanner { return nil }
func (a *decoder) scanBasic() scanner     { return nil }
