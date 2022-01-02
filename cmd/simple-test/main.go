package main

import (
	"encoding/binary"
	"fmt"
	"math/bits"

	hfuzz "github.com/thebabush/go-hfuzz/hfuzz"
)

func main() {
	fmt.Println("Here we go!")

	max := 0xFFFFFF

	hf := hfuzz.New()
	hf.Persistent(func(data []byte) {
		if len(data) < 8 {
			return
		}
		bigNumber := binary.BigEndian.Uint64(data)

		const TARGET = uint64(0xC0CAF16ADEADBEEF)

		// Debug output
		ones := bits.OnesCount64(TARGET ^ bigNumber)
		if ones < max {
			fmt.Println("NEW BEST: ", ones, " ", fmt.Sprintf("%08X", bigNumber))
			max = ones
		}

		// Update coverage (manually)
		hf.TraceCmp8(0x666, bigNumber, TARGET)
		if bigNumber == TARGET {
			// panic("WIN")
			hfuzz.WinWithCode(3, ":D")
		}
	})
}
