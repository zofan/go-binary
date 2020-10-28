package binary

import (
	"bytes"
	"fmt"
	"math"
	"testing"
)

func Test(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	fmt.Println(WriteByte(buf, 1))
	fmt.Println(WriteBytes(buf, []byte{2, 3, 4, 5}))
	fmt.Println(WriteString(buf, `abc`))
	fmt.Println(WriteRune(buf, '🙃'))

	fmt.Println(WriteUint8(buf, math.MaxUint8))
	fmt.Println(WriteUint16(buf, math.MaxUint16))
	fmt.Println(WriteUint32(buf, math.MaxUint32))
	fmt.Println(WriteUint64(buf, math.MaxUint64))

	fmt.Println(WriteInt8(buf, math.MaxInt8))
	fmt.Println(WriteInt16(buf, math.MaxInt16))
	fmt.Println(WriteInt32(buf, math.MaxInt32))
	fmt.Println(WriteInt64(buf, math.MaxInt64))

	fmt.Println(WriteFloat32(buf, math.MaxFloat32))
	fmt.Println(WriteFloat64(buf, math.MaxFloat64))

	fmt.Println(`----------`)

	fmt.Println(WriteRunes(buf, []rune{'🙃', '🙃', '🙃'}))

	fmt.Println(WriteUint8s(buf, []uint8{math.MaxUint8, math.MaxUint8 / 2}))
	fmt.Println(WriteUint16s(buf, []uint16{math.MaxUint16, math.MaxUint16 / 2}))
	fmt.Println(WriteUint32s(buf, []uint32{math.MaxUint32, math.MaxUint32 / 2}))
	fmt.Println(WriteUint64s(buf, []uint64{math.MaxUint64, math.MaxUint64 / 2}))

	fmt.Println(WriteInt8s(buf, []int8{math.MaxInt8, math.MaxInt8 / 2}))
	fmt.Println(WriteInt16s(buf, []int16{math.MaxInt16, math.MaxInt16 / 2}))
	fmt.Println(WriteInt32s(buf, []int32{math.MaxInt32, math.MaxInt32 / 2}))
	fmt.Println(WriteInt64s(buf, []int64{math.MaxInt64, math.MaxInt64 / 2}))

	fmt.Println(WriteFloat32s(buf, []float32{math.MaxFloat32, math.MaxFloat32 / 2}))
	fmt.Println(WriteFloat64s(buf, []float64{math.MaxFloat64, math.MaxFloat64 / 2}))

	var x []uint64
	for i := 0; i < 200; i++ {
		x = append(x, math.MaxUint64)
	}
	fmt.Println(WriteUint64s(buf, x))

	fmt.Println(`----------`)

	fmt.Println(ReadByte(buf))
	fmt.Println(ReadBytes(buf, 4))
	fmt.Println(ReadString(buf, 3))
	fmt.Println(ReadRune(buf))

	fmt.Println(ReadUint8(buf))
	fmt.Println(ReadUint16(buf))
	fmt.Println(ReadUint32(buf))
	fmt.Println(ReadUint64(buf))

	fmt.Println(ReadInt8(buf))
	fmt.Println(ReadInt16(buf))
	fmt.Println(ReadInt32(buf))
	fmt.Println(ReadInt64(buf))

	fmt.Println(ReadFloat32(buf))
	fmt.Println(ReadFloat64(buf))

	fmt.Println(`----------`)

	fmt.Println(ReadRunes(buf, 3))

	fmt.Println(ReadUint8s(buf, 2))
	fmt.Println(ReadUint16s(buf, 2))
	fmt.Println(ReadUint32s(buf, 2))
	fmt.Println(ReadUint64s(buf, 2))

	fmt.Println(ReadInt8s(buf, 2))
	fmt.Println(ReadInt16s(buf, 2))
	fmt.Println(ReadInt32s(buf, 2))
	fmt.Println(ReadInt64s(buf, 2))

	fmt.Println(ReadFloat32s(buf, 2))
	fmt.Println(ReadFloat64s(buf, 2))
}
