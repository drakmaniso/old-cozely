package ciziel

type stack struct {
	stack []Data
	keys  []*string
}

func (ds *stack) push(d Data) {
	ds.stack = append(ds.stack, d)
	ds.keys = append(ds.keys, nil)
}

func (ds *stack) pop() Data {
	d := ds.stack[len(ds.stack)-1]
	ds.stack = ds.stack[:len(ds.stack)-1]
	ds.keys = ds.keys[:len(ds.keys)-1]
	return d
}

func (ds *stack) replace(d Data) {
	ds.stack[len(ds.stack)-1] = d
}

func (ds *stack) setkey(k *string) {
	ds.keys[len(ds.keys)-1] = k
}

func (ds *stack) key() *string {
	return ds.keys[len(ds.keys)-1]
}

func (ds *stack) peek() Data {
	return ds.stack[len(ds.stack)-1]
}

func (ds *stack) empty() bool {
	return len(ds.stack) == 0
}
