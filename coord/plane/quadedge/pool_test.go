package quadedge_test

import (
	"testing"

	"github.com/cozely/cozely/coord/plane/quadedge"
)

////////////////////////////////////////////////////////////////////////////////

func TestPool_New(t *testing.T) {
	p := quadedge.NewPool(2)
	e0 := quadedge.New(p)
	if e0.SameRing(e0.Sym()) {
		t.Error("new edge incorrect, eOrg==eDest: ", e0)
	}
	if !e0.Rot().SameRing(e0.Tor()) {
		t.Error("new edge incorrect, eRight!=eLeft: ", e0)
	}
	if e0.Sym().SameRing(e0) {
		t.Error("new edge incorrect, eDest==eOrg: ", e0)
	}
	if !e0.Tor().SameRing(e0.Rot()) {
		t.Error("new edge incorrect, eLeft!=eRight: ", e0)
	}
	if e0.OrigNext() != e0 ||
		e0.RightNext() != e0.Sym() ||
		e0.DestNext() != e0 ||
		e0.LeftNext() != e0.Sym() {
		t.Error("new edge incorrect, got: ", e0)
	}
	s0 := e0.String()
	if s0 != "0o=[orig:->0o right:->0l dest:->0d left:->0r]" {
		t.Error("new edge incorrect, got: ", s0)
	}
	e1 := quadedge.New(p)
	e1.SetOrig(1)
	e1.SetRight(100)
	e1.SetDest(2)
	e1.SetLeft(200)
	s1 := e1.String()
	if s1 != "1o=[orig:1->1o right:100->1l dest:2->1d left:200->1r]" {
		t.Error("new edge incorrect, got: ", s1)
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestEdge_Rot(t *testing.T) {
	p := quadedge.NewPool(1)
	e0 := quadedge.New(p)
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
	e0 := quadedge.New(p)
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
	e0 := quadedge.New(p)
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
	e3 := e0.Tor().Rot()
	if e3 != e0 {
		t.Error("Tor().Rot() incorrect, got:", e2)
	}
}

////////////////////////////////////////////////////////////////////////////////
