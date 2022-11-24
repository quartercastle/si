package si

import (
	"fmt"
	"testing"
)

func TestDataPretty(t *testing.T) {
	value, unit := Readable(24*Byte, ByteBase)
	fmt.Println(value, unit)
}
