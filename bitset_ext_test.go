package bitset

import (
	"testing"
)

func TestBitSet_FirstSet(t *testing.T) {
	s1 := New(8).Set(4).Set(7).Set(0)
	t.Log(s1.DumpAsBits())
	t.Log(s1.FirstSet())
}

func TestBitSet_11(t *testing.T) {
	s1 := New(64).InsertAt(1)
	t.Log(s1.DumpAsBits())
	t.Log(s1.FirstSet())
}

func Test_copy(t *testing.T) {
	a := New(32).Set(1)
	t.Log(a.DumpAsBits())
	b := a.Clone()
	t.Log(b.DumpAsBits())

	a.Set(2)
	b.Set(3)
	t.Log(a.DumpAsBits())

	t.Log(b.DumpAsBits())

}

func Test_test(t *testing.T) {
	a := New(32).Set(1)
	t.Log(a.DumpAsBits())
	t.Log(a.Test(1))
	t.Log(a.Test(2))
	t.Log(a.Test(0))

}

func Test_fill(t *testing.T) {
	a := New(128)
	b := a.FlipRange(0, 128)
	t.Log(a.DumpAsBits())
	t.Log(b.DumpAsBits())

}
