package icon

import (
	"github.com/mg/i/itesting"
	"testing"
)

func TestBool(t *testing.T) {
	itr := Bools([]bool{true, false, true, true, false})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, true, false, true, true, false)
}

func TestByte(t *testing.T) {
	itr := Bytes([]byte{'1', 'a', 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, byte('1'), byte('a'), byte(100))
}

func TestComplex128(t *testing.T) {
	itr := Complex128s([]complex128{complex(1, 1), 0 + 3i, 4 - 2i, 1 + 0i})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, complex(1, 1), 0+3i, 4-2i, 1+0i)
}

func TestComplex64(t *testing.T) {
	itr := Complex64s([]complex64{complex(1, 1), 0 + 3i, 4 - 2i, 1 + 0i})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, complex64(complex(1, 1)), complex64(0+3i), complex64(4-2i), complex64(1+0i))
}

func TestFloat32(t *testing.T) {
	itr := Float32s([]float32{1.0, 34.0, 0.123123, 1})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, float32(1.0), float32(34.0), float32(0.123123), float32(1))
}

func TestFloat64(t *testing.T) {
	itr := Float64s([]float64{1.0, 34.0, 0.123123, 1})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, 1.0, 34.0, 0.123123, float64(1))
}

func TestInt(t *testing.T) {
	itr := Ints([]int{1, 5, 6, -4, 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, 1, 5, 6, -4, 100)
}

func TestInt16(t *testing.T) {
	itr := Int16s([]int16{1, 5, 6, -4, 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, int16(1), int16(5), int16(6), int16(-4), int16(100))
}

func TestInt32(t *testing.T) {
	itr := Int32s([]int32{1, 5, 6, -4, 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, int32(1), int32(5), int32(6), int32(-4), int32(100))
}

func TestInt64(t *testing.T) {
	itr := Int64s([]int64{1, 5, 6, -4, 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, int64(1), int64(5), int64(6), int64(-4), int64(100))
}

func TestInt8(t *testing.T) {
	itr := Int8s([]int8{1, 5, 6, -4, 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, int8(1), int8(5), int8(6), int8(-4), int8(100))
}

func TestRune(t *testing.T) {
	itr := Runes([]rune{'a', 'þ', '世', '界', 123})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, 'a', 'þ', '世', '界', rune(123))
}

func TestString(t *testing.T) {
	itr := Strings([]string{"hello", "goodbye", "世界"})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, "hello", "goodbye", "世界")
}

func TestUint16(t *testing.T) {
	itr := Uint16s([]uint16{1, 5, 6, 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, uint16(1), uint16(5), uint16(6), uint16(100))
}

func TestUint32(t *testing.T) {
	itr := Uint32s([]uint32{1, 5, 6, 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, uint32(1), uint32(5), uint32(6), uint32(100))
}

func TestUint64(t *testing.T) {
	itr := Uint64s([]uint64{1, 5, 6, 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, uint64(1), uint64(5), uint64(6), uint64(100))
}

func TestUint8(t *testing.T) {
	itr := Uint8s([]uint8{1, 5, 6, 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, uint8(1), uint8(5), uint8(6), uint8(100))
}
