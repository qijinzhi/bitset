# bitset

*Go*在非负整数和布尔值之间进行映射的库。clone自*bits-and-blooms/bitset*仓库。

## 说明

### 示例:

```go
package main

import (
	"fmt"
	"math/rand"

	"github.com/qijinzhi/bitset"
)

func main() {
	fmt.Printf("Hello from BitSet!\n")
	var b bitset.BitSet
	// play some Go Fish
	for i := 0; i < 100; i++ {
		card1 := uint(rand.Intn(52))
		card2 := uint(rand.Intn(52))
		b.Set(card1)
		if b.Test(card2) {
			fmt.Println("Go Fish!")
		}
		b.Clear(card1)
	}

	// Chaining
	b.Set(10).Set(11)

	for i, e := b.NextSet(0); e; i, e = b.NextSet(i + 1) {
		fmt.Println("The following bit is set:", i)
	}
	if b.Intersection(bitset.New(100).Set(10)).Count() == 1 {
		fmt.Println("Intersection works.")
	} else {
		fmt.Println("Intersection doesn't work???")
	}
}
```


## 序列化

```Go
    const length = 9585
	const oneEvery = 97
	bs := bitset.New(length)
	// Add some bits
	for i := uint(0); i < length; i += oneEvery {
		bs = bs.Set(i)
	}

	var buf bytes.Buffer
	n, err := bs.WriteTo(&buf)
	if err != nil {
		// failure
	}
	// Here n == buf.Len()
```
