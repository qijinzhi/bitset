package bitset

// FirstSet b最后为1的所在位置（非下标），从右到左
func (b *BitSet) FirstSet() uint {
	panicIfNull(b)
	for i := uint(0); i < b.length; i++ {
		if b.set[i>>log2WordSize]&(1<<wordsIndex(i)) != 0 {
			return i + 1
		}
	}
	return 0
}
