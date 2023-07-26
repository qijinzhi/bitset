package bitset

// Incr128 增加1
func (b *BitSet) Incr128() *BitSet {
	panicIfNull(b)
	if (b.set[0] & allBits) == allBits {
		// 低位已经为最大数，则设置为0
		b.set[0] = 0
	} else {
		// 低位未为最大数，则设置+1；并直接返回b
		b.set[0] += 1
		return b
	}
	if (b.set[1] & allBits) == allBits {
		// 低位已经为最大数，则设置为0
		return nil
	} else {
		// 低位未为最大数，则设置+1；并直接返回b
		b.set[1] += 1
		return b
	}
}

// Cap128 128位对比b和c的大小。b>c返回1；b==c 返回0；b<c返回-1。
func (b *BitSet) Cap128(c *BitSet) int {
	panicIfNull(b)
	panicIfNull(c)
	if b.set[1] > c.set[1] {
		return 1
	} else if b.set[1] < c.set[1] {
		return -1
	} else {
		if b.set[0] > c.set[0] {
			return 1
		} else if b.set[0] < c.set[0] {
			return -1
		} else {
			return 0
		}
	}
}
