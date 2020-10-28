package binary

import (
	bin "encoding/binary"
	"errors"
	"io"
	"math"
)

const (
	ByteNULL = 0x00
	ByteSOH  = 0x01
	ByteSTX  = 0x02
	ByteETX  = 0x03
	ByteES   = 0x19
	ByteFS   = 0x1C
	ByteGS   = 0x1D
	ByteRS   = 0x1E
	ByteUS   = 0x1F
	ByteDEL  = 0x7F
)

var (
	ErrInvalidBytes = errors.New(`binary: invalid bytes`)
)

func WriteBytes(w io.Writer, b []byte) (e error) {
	if len(b) == 0 {
		return nil
	}

	n, e := w.Write(b)
	if n != len(b) {
		return ErrInvalidBytes
	}
	return
}

func WriteByte(w io.Writer, b byte) (e error) {
	return WriteBytes(w, []byte{b})
}

func WriteString(w io.Writer, s string) (e error) {
	return WriteBytes(w, []byte(s))
}

func WriteRune(w io.Writer, v rune) (e error) {
	return WriteUint32(w, uint32(v))
}

func WriteUint8(w io.Writer, v uint8) (e error) {
	return WriteByte(w, v)
}

func WriteUint16(w io.Writer, v uint16) (e error) {
	b := make([]byte, 2)
	bin.LittleEndian.PutUint16(b, v)
	return WriteBytes(w, b)
}

func WriteUint32(w io.Writer, v uint32) (e error) {
	b := make([]byte, 4)
	bin.LittleEndian.PutUint32(b, v)
	return WriteBytes(w, b)
}

func WriteUint64(w io.Writer, v uint64) (e error) {
	b := make([]byte, 8)
	bin.LittleEndian.PutUint64(b, v)
	return WriteBytes(w, b)
}

func WriteInt8(w io.Writer, v int8) (e error) {
	return WriteByte(w, byte(v))
}

func WriteInt16(w io.Writer, v int16) (e error) {
	return WriteUint16(w, uint16(v))
}

func WriteInt32(w io.Writer, v int32) (e error) {
	return WriteUint32(w, uint32(v))
}

func WriteInt64(w io.Writer, v int64) (e error) {
	return WriteUint64(w, uint64(v))
}

func WriteFloat32(w io.Writer, v float32) (e error) {
	return WriteUint32(w, math.Float32bits(v))
}

func WriteFloat64(w io.Writer, v float64) (e error) {
	return WriteUint64(w, math.Float64bits(v))
}

func WriteUint8s(w io.Writer, vs []uint8) (e error) {
	return WriteBytes(w, vs)
}

func WriteUint16s(w io.Writer, vs []uint16) (e error) {
	b := make([]byte, len(vs)*2)
	for i, v := range vs {
		bin.LittleEndian.PutUint16(b[i*2:], v)
	}
	return WriteBytes(w, b)
}

func WriteUint32s(w io.Writer, vs []uint32) (e error) {
	b := make([]byte, len(vs)*4)
	for i, v := range vs {
		bin.LittleEndian.PutUint32(b[i*4:], v)
	}
	return WriteBytes(w, b)
}

func WriteUint64s(w io.Writer, vs []uint64) (e error) {
	b := make([]byte, len(vs)*8)
	for i, v := range vs {
		bin.LittleEndian.PutUint64(b[i*8:], v)
	}
	return WriteBytes(w, b)
}

func WriteInt8s(w io.Writer, vs []int8) (e error) {
	b := make([]byte, 0, len(vs))
	for _, v := range vs {
		b = append(b, uint8(v))
	}
	return WriteBytes(w, b)
}

func WriteInt16s(w io.Writer, vs []int16) (e error) {
	b := make([]byte, len(vs)*2)
	for i, v := range vs {
		bin.LittleEndian.PutUint16(b[i*2:], uint16(v))
	}
	return WriteBytes(w, b)
}

func WriteInt32s(w io.Writer, vs []int32) (e error) {
	b := make([]byte, len(vs)*4)
	for i, v := range vs {
		bin.LittleEndian.PutUint32(b[i*4:], uint32(v))
	}
	return WriteBytes(w, b)
}

func WriteInt64s(w io.Writer, vs []int64) (e error) {
	b := make([]byte, len(vs)*8)
	for i, v := range vs {
		bin.LittleEndian.PutUint64(b[i*8:], uint64(v))
	}
	return WriteBytes(w, b)
}

func WriteFloat32s(w io.Writer, vs []float32) (e error) {
	b := make([]byte, len(vs)*4)
	for i, v := range vs {
		bin.LittleEndian.PutUint32(b[i*4:], math.Float32bits(v))
	}
	return WriteBytes(w, b)
}

func WriteFloat64s(w io.Writer, vs []float64) (e error) {
	b := make([]byte, len(vs)*8)
	for i, v := range vs {
		bin.LittleEndian.PutUint64(b[i*8:], math.Float64bits(v))
	}
	return WriteBytes(w, b)
}

func WriteRunes(w io.Writer, vs []rune) (e error) {
	b := make([]byte, len(vs)*4)
	for i, v := range vs {
		bin.LittleEndian.PutUint32(b[i*4:], uint32(v))
	}
	return WriteBytes(w, b)
}
