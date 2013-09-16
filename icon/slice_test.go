package icon

import (
	"fmt"
	"github.com/mg/i/itesting"
	"testing"
)

func TestBool(t *testing.T) {
	itr := Bools([]bool{true, false, true, true, false})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, true, false, true, true, false)
}

func ExampleBool() {
	slice := []bool{true, false, true, true, false}
	itr := Bools(slice)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Bool())
	}
}

func ExampleBoolList() {
	itr := BoolList(true, false, true, true, false)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Bool())
	}
}

func TestByte(t *testing.T) {
	itr := Bytes([]byte{'1', 'a', 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, byte('1'), byte('a'), byte(100))
}

func ExampleBytes() {
	slice := []byte{'1', 'a', 100}
	itr := Bytes(slice)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Byte())
	}
}

func ExampleBytelList() {
	itr := ByteList('1', 'a', 100)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Byte())
	}
}

func TestComplex128(t *testing.T) {
	itr := Complex128s([]complex128{complex(1, 1), 0 + 3i, 4 - 2i, 1 + 0i})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, complex(1, 1), 0+3i, 4-2i, 1+0i)
}

func ExampleComplex128s() {
	slice := []complex128{complex(1, 1), 0 + 3i, 4 - 2i, 1 + 0i}
	itr := Complex128s(slice)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Complex128())
	}
}

func ExampleComplex128List() {
	itr := Complex128List(complex(1, 1), 0+3i, 4-2i, 1+0i)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Complex128())
	}
}

func TestComplex64(t *testing.T) {
	itr := Complex64s([]complex64{complex(1, 1), 0 + 3i, 4 - 2i, 1 + 0i})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, complex64(complex(1, 1)), complex64(0+3i), complex64(4-2i), complex64(1+0i))
}

func ExampleComplex64s() {
	slice := []complex64{complex(1, 1), 0 + 3i, 4 - 2i, 1 + 0i}
	itr := Complex64s(slice)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Complex64())
	}
}

func ExampleComplex64List() {
	itr := Complex64List(complex(1, 1), 0+3i, 4-2i, 1+0i)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Complex64())
	}
}

func TestFloat32(t *testing.T) {
	itr := Float32s([]float32{1.0, 34.0, 0.123123, 1})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, float32(1.0), float32(34.0), float32(0.123123), float32(1))
}

func ExampleFloat32s() {
	slice := []float32{1.0, 34.0, 0.123123, 1}
	itr := Float32s(slice)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Float32())
	}
}

func ExampleFloat32List() {
	itr := Float32List(1.0, 34.0, 0.123123, 1)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Float32())
	}
}

func TestFloat64(t *testing.T) {
	itr := Float64s([]float64{1.0, 34.0, 0.123123, 1})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, 1.0, 34.0, 0.123123, float64(1))
}

func ExampleFloat64s() {
	slice := []float64{1.0, 34.0, 0.123123, 1}
	itr := Float64s(slice)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Float64())
	}
}

func ExampleFloat64List() {
	itr := Float64List(1.0, 34.0, 0.123123, 1)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Float64())
	}
}

func TestInt(t *testing.T) {
	itr := Ints([]int{1, 5, 6, -4, 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, 1, 5, 6, -4, 100)
}

func ExampleInts() {
	slice := []int{1, 5, 6, -4, 100}
	itr := Ints(slice)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Int())
	}
}

func ExampleIntList() {
	itr := IntList(1, 5, 6, -4, 100)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Int())
	}
}

func TestInt16(t *testing.T) {
	itr := Int16s([]int16{1, 5, 6, -4, 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, int16(1), int16(5), int16(6), int16(-4), int16(100))
}

func ExampleInt16s() {
	slice := []int16{1, 5, 6, -4, 100}
	itr := Int16s(slice)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Int16())
	}
}

func ExampleInt16List() {
	itr := Int16List(1, 5, 6, -4, 100)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Int16())
	}
}

func TestInt32(t *testing.T) {
	itr := Int32s([]int32{1, 5, 6, -4, 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, int32(1), int32(5), int32(6), int32(-4), int32(100))
}

func ExampleInt32s() {
	slice := []int32{1, 5, 6, -4, 100}
	itr := Int32s(slice)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Int32())
	}
}

func ExampleInt32List() {
	itr := Int32List(1, 5, 6, -4, 100)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Int32())
	}
}

func TestInt64(t *testing.T) {
	itr := Int64s([]int64{1, 5, 6, -4, 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, int64(1), int64(5), int64(6), int64(-4), int64(100))
}

func ExampleInt64s() {
	slice := []int64{1, 5, 6, -4, 100}
	itr := Int64s(slice)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Int64())
	}
}

func ExampleInt64List() {
	itr := Int64List(1, 5, 6, -4, 100)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Int64())
	}
}

func TestInt8(t *testing.T) {
	itr := Int8s([]int8{1, 5, 6, -4, 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, int8(1), int8(5), int8(6), int8(-4), int8(100))
}

func ExampleInt8s() {
	slice := []int8{1, 5, 6, -4, 100}
	itr := Int8s(slice)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Int8())
	}
}

func ExampleInt8List() {
	itr := Int8List(1, 5, 6, -4, 100)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Int8())
	}
}

func TestRune(t *testing.T) {
	itr := Runes([]rune{'a', 'þ', '世', '界', 123})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, 'a', 'þ', '世', '界', rune(123))
}

func ExampleRunes() {
	slice := []rune{'a', 'þ', '世', '界', 123}
	itr := Runes(slice)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Rune())
	}
}

func ExampleRuneList() {
	itr := RuneList('a', 'þ', '世', '界', 123)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Rune())
	}
}

func TestString(t *testing.T) {
	itr := Strings([]string{"hello", "goodbye", "世界"})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, "hello", "goodbye", "世界")
}

func ExampleStrings() {
	slice := []string{"hello", "goodbye", "世界"}
	itr := Strings(slice)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.String())
	}
}

func ExampleStringList() {
	itr := StringList("hello", "goodbye", "世界")
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.String())
	}
}

func TestUint16(t *testing.T) {
	itr := Uint16s([]uint16{1, 5, 6, 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, uint16(1), uint16(5), uint16(6), uint16(100))
}

func ExampleUint16s() {
	slice := []uint16{1, 5, 6, 100}
	itr := Uint16s(slice)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Uint16())
	}
}

func ExampleUint16List() {
	itr := Uint16List(1, 5, 6, 100)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Uint16())
	}
}

func TestUint32(t *testing.T) {
	itr := Uint32s([]uint32{1, 5, 6, 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, uint32(1), uint32(5), uint32(6), uint32(100))
}

func ExampleUint32s() {
	slice := []uint32{1, 5, 6, 100}
	itr := Uint32s(slice)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Uint32())
	}
}

func ExampleUint32List() {
	itr := Uint32List(1, 5, 6, 100)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Uint32())
	}
}

func TestUint64(t *testing.T) {
	itr := Uint64s([]uint64{1, 5, 6, 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, uint64(1), uint64(5), uint64(6), uint64(100))
}

func ExampleUint64s() {
	slice := []uint64{1, 5, 6, 100}
	itr := Uint64s(slice)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Uint64())
	}
}

func ExampleUint64List() {
	itr := Uint64List(1, 5, 6, 100)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Uint64())
	}
}

func TestUint8(t *testing.T) {
	itr := Uint8s([]uint8{1, 5, 6, 100})
	itesting.AssertRandomAccess(t, itr, itesting.Strict)
	itr.First()
	itesting.AssertIteration(t, itr, uint8(1), uint8(5), uint8(6), uint8(100))
}

func ExampleUint8s() {
	slice := []uint8{1, 5, 6, 100}
	itr := Uint8s(slice)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Uint8())
	}
}

func ExampleUint8List() {
	itr := Uint8List(1, 5, 6, 100)
	for ; !itr.AtEnd(); itr.Next() {
		fmt.Println(itr.Uint8())
	}
}
