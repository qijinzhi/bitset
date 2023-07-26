package bitset

import "testing"

func TestBitSet_Incr128(t *testing.T) {
	s1 := New(128)
	s1.FlipRange(0, 65)
	t.Log(s1.DumpAsBits())
	// t.Log(s1.set, s1.Count())
	s2 := s1.Incr128()
	t.Log(s2.DumpAsBits())
}

func TestBitSet_Incr1282(t *testing.T) {
	s1 := New(128)
	s1.FlipRange(0, 64)
	t.Log(s1.DumpAsBits())
	// t.Log(s1.set, s1.Count())
	s2 := s1.Incr128()
	t.Log(s2.DumpAsBits())
}
