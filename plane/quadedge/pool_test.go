package quadedge_test

import (
	"testing"

	"github.com/drakmaniso/glam/plane/quadedge"
)

//------------------------------------------------------------------------------

func TestPool_NewEdge(t *testing.T) {
	p := quadedge.NewPool(2)
	e0 := p.NewEdge()
	s0 := e0.String()
	if s0 != "0o=[Onext:0o Rnext:0l Dnext:0d Lnext:0r]" {
		t.Error("new edge incorrect, got: ", s0)
	}
	e1 := p.NewEdge()
	s1 := e1.String()
	if s1 != "1o=[Onext:1o Rnext:1l Dnext:1d Lnext:1r]" {
		t.Error("new edge incorrect, got: ", s1)
	}
}

//------------------------------------------------------------------------------

func TestEdge_Rot(t *testing.T) {
	p := quadedge.NewPool(1)
	e0 := p.NewEdge()
	e0.SetOrig(1)
	e0.SetRight(100)
	e0.SetDest(2)
	e0.SetLeft(200)
	e1 := e0.Rot()
	if e1.Orig() != 100 ||
		e1.Right() != 2 ||
		e1.Dest() != 200 ||
		e1.Left() != 1 {
		t.Error("first Rot() incorrect, got: ", e1)
	}
	e2 := e0.Rot().Rot().Rot().Rot()
	if e2 != e0 {
		t.Error("quadruple Rot() incorrect, got:", e2)
	}
}

func TestEdge_Sym(t *testing.T) {
	p := quadedge.NewPool(1)
	e0 := p.NewEdge()
	e0.SetOrig(1)
	e0.SetRight(100)
	e0.SetDest(2)
	e0.SetLeft(200)
	e1 := e0.Sym()
	if e1.Orig() != 2 ||
		e1.Right() != 200 ||
		e1.Dest() != 1 ||
		e1.Left() != 100 {
		t.Error("first Sym() incorrect, got: ", e1)
	}
	e2 := e0.Sym().Sym()
	if e2 != e0 {
		t.Error("double Sym() incorrect, got:", e2)
	}
}

func TestEdge_Tor(t *testing.T) {
	p := quadedge.NewPool(1)
	e0 := p.NewEdge()
	e0.SetOrig(1)
	e0.SetRight(100)
	e0.SetDest(2)
	e0.SetLeft(200)
	e1 := e0.Tor()
	if e1.Orig() != 200 ||
		e1.Right() != 1 ||
		e1.Dest() != 100 ||
		e1.Left() != 2 {
		t.Error("first Tor() incorrect, got: ", e1)
	}
	e2 := e0.Tor().Tor().Tor().Tor()
	if e2 != e0 {
		t.Error("quadruple Tor() incorrect, got:", e2)
	}
}

//------------------------------------------------------------------------------
